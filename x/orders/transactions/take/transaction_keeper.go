// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package take

import (
	"context"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/data"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/documents/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	propertyConstants "github.com/AssetMantle/schema/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/burn"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/x/orders/constants"
	"github.com/AssetMantle/modules/x/orders/key"
	"github.com/AssetMantle/modules/x/orders/mappable"
	"github.com/AssetMantle/modules/x/orders/record"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/transfer"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	burnAuxiliary         helpers.Auxiliary
	supplementAuxiliary   helpers.Auxiliary
	transferAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	//TODO check if dec or int is correct usage there
	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message)); err != nil {
		return nil, err
	}

	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(message.OrderID))

	Mappable := orders.GetMappable(key.NewKey(message.OrderID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("order with ID %s not found", message.OrderID.AsString())
	}
	order := mappable.GetOrder(Mappable)

	if order.GetTakerID().Compare(baseIDs.PrototypeIdentityID()) != 0 && order.GetTakerID().Compare(message.GetFromIdentityID()) != 0 {
		return nil, errorConstants.NotAuthorized.Wrapf("taker ID %s is not authorized to take private order with ID %s", message.GetFromIdentityID().AsString(), message.OrderID.AsString())
	}
	takerSplit, _ := sdkTypes.NewIntFromString(message.TakerSplit)
	makerReceiveTakerSplit := order.GetMakerSplit().ToLegacyDec().MulTruncate(order.GetExchangeRate()).MulTruncate(sdkTypes.SmallestDec())
	takerReceiveMakerSplit := takerSplit.ToLegacyDec().QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(order.GetExchangeRate())

	switch updatedMakerSplit := order.GetMakerSplit().Sub(takerReceiveMakerSplit.TruncateInt()); {
	case updatedMakerSplit.Equal(sdkTypes.ZeroInt()):
		if takerSplit.ToLegacyDec().LT(makerReceiveTakerSplit) {
			return nil, errorConstants.InsufficientBalance.Wrapf("taker split %s is less than the required amount %s for order execution", message.TakerSplit, makerReceiveTakerSplit.String())
		}

		orders.Remove(record.NewRecord(order))
	case updatedMakerSplit.LT(sdkTypes.ZeroInt()):
		if takerSplit.ToLegacyDec().LT(makerReceiveTakerSplit) {
			return nil, errorConstants.InsufficientBalance.Wrapf("taker split %s is less than the required amount %s for order execution", message.TakerSplit, makerReceiveTakerSplit.String())
		}

		takerReceiveMakerSplit = order.GetMakerSplit().ToLegacyDec()

		orders.Remove(record.NewRecord(order))
	default:
		makerReceiveTakerSplit = takerSplit.ToLegacyDec()
		mutableProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(propertyConstants.MakerSplitProperty.GetKey(), baseData.NewNumberData(updatedMakerSplit)))

		updatedOrder := base.NewOrder(order.GetClassificationID(), order.GetImmutables(), order.GetMutables().Mutate(baseLists.AnyPropertiesToProperties(mutableProperties.Get()...)...))

		if err := updatedOrder.ValidateBasic(); err != nil {
			return nil, err
		}

		orders.Mutate(record.NewRecord(updatedOrder))
	}

	if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.GetFromIdentityID(), order.GetMakerID(), order.GetTakerAssetID(), makerReceiveTakerSplit.TruncateInt())); err != nil {
		return nil, err
	}

	if _, err := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(constants.ModuleIdentity.GetModuleIdentityID(), message.GetFromIdentityID(), order.GetMakerAssetID(), takerReceiveMakerSplit.TruncateInt())); err != nil {
		return nil, err
	}

	bondAmount := sdkTypes.ZeroInt()
	if bondAmountProperty := order.GetProperty(propertyConstants.BondAmountProperty.GetID()); bondAmountProperty == nil || !bondAmountProperty.IsMeta() {
		return nil, errorConstants.MetaDataError.Wrapf("order with ID %s has no revealed bond amount", message.OrderID)
	} else {
		bondAmount = bondAmountProperty.Get().(properties.MetaProperty).GetData().Get().(data.NumberData).Get()
	}

	if _, err := transactionKeeper.burnAuxiliary.GetKeeper().Help(context, burn.NewAuxiliaryRequest(order.GetClassificationID(), bondAmount)); err != nil {
		return nil, err
	}

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				transactionKeeper.transferAuxiliary = value
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
				transactionKeeper.authenticateAuxiliary = value
			case burn.Auxiliary.GetName():
				transactionKeeper.burnAuxiliary = value
			}
		}
	}

	helpers.PanicOnUninitializedKeeperFields(transactionKeeper)
	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
