// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"context"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/modules/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/modules/orders/internal/key"
	"github.com/AssetMantle/modules/modules/orders/internal/mappable"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/modules/splits/auxiliaries/transfer"
	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/documents/base"
	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/helpers"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseLists "github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/properties/constants"
	"github.com/AssetMantle/modules/schema/properties/utilities"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	baseTypes "github.com/AssetMantle/modules/schema/types/base"
)

type transactionKeeper struct {
	mapper                     helpers.Mapper
	parameters                 helpers.ParameterManager
	conformAuxiliary           helpers.Auxiliary
	supplementAuxiliary        helpers.Auxiliary
	transferAuxiliary          helpers.Auxiliary
	authenticateAuxiliary      helpers.Auxiliary
	maintainersVerifyAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {

	if auxiliaryResponse := transactionKeeper.maintainersVerifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.ClassificationID, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	fromAddress, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil {
		panic("Could not get from address from Bech32 string")
	}

	if auxiliaryResponse := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	if auxiliaryResponse := transactionKeeper.transferAuxiliary.GetKeeper().Help(context, transfer.NewAuxiliaryRequest(message.FromID, module.ModuleIdentityID, message.MakerOwnableID, message.MakerOwnableSplit)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	immutableMetaProperties := message.ImmutableMetaProperties.
		Add(baseProperties.NewMetaProperty(constants.ExchangeRateProperty.GetKey(), baseData.NewDecData(message.TakerOwnableSplit.QuoTruncate(sdkTypes.SmallestDec()).QuoTruncate(message.MakerOwnableSplit)))).
		Add(baseProperties.NewMetaProperty(constants.CreationHeightProperty.GetKey(), baseData.NewHeightData(baseTypes.NewHeight(sdkTypes.UnwrapSDKContext(context).BlockHeight())))).
		Add(baseProperties.NewMetaProperty(constants.MakerOwnableIDProperty.GetKey(), baseData.NewIDData(message.MakerOwnableID))).
		Add(baseProperties.NewMetaProperty(constants.TakerOwnableIDProperty.GetKey(), baseData.NewIDData(message.TakerOwnableID))).
		Add(baseProperties.NewMetaProperty(constants.MakerIDProperty.GetKey(), baseData.NewIDData(message.FromID))).
		Add(baseProperties.NewMetaProperty(constants.TakerIDProperty.GetKey(), baseData.NewIDData(message.TakerID)))

	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(utilities.AnyPropertyListToPropertyList(append(immutableMetaProperties.GetList(), message.ImmutableProperties.GetList()...)...)...))

	orderID := baseIDs.NewOrderID(message.ClassificationID, immutables)
	orders := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(orderID))
	if orders.Get(key.NewKey(orderID)) != nil {
		return nil, errorConstants.EntityAlreadyExists
	}

	if orderLife := message.ExpiryHeight.Get() - sdkTypes.UnwrapSDKContext(context).BlockHeight(); orderLife < 0 || orderLife > transactionKeeper.parameters.GetParameter(constants.MaxOrderLifeProperty.GetID()).GetMetaProperty().GetData().Get().(data.HeightData).Get().Get() {
		return nil, errorConstants.InvalidRequest
	}

	mutableMetaProperties := message.MutableMetaProperties.
		Add(baseProperties.NewMetaProperty(constants.ExpiryHeightProperty.GetKey(), baseData.NewHeightData(message.ExpiryHeight))).
		Add(baseProperties.NewMetaProperty(constants.MakerOwnableSplitProperty.GetKey(), baseData.NewDecData(message.MakerOwnableSplit)))

	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(utilities.AnyPropertyListToPropertyList(append(mutableMetaProperties.GetList(), message.MutableProperties.GetList()...)...)...))

	if auxiliaryResponse := transactionKeeper.conformAuxiliary.GetKeeper().Help(context, conform.NewAuxiliaryRequest(message.ClassificationID, immutables, mutables)); !auxiliaryResponse.IsSuccessful() {
		return nil, auxiliaryResponse.GetError()
	}

	orders.Add(mappable.NewMappable(base.NewOrder(message.ClassificationID, immutables, mutables)))

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameters helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameters = mapper, parameters

	for _, externalKeeper := range auxiliaries {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case conform.Auxiliary.GetName():
				transactionKeeper.conformAuxiliary = value
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case transfer.Auxiliary.GetName():
				transactionKeeper.transferAuxiliary = value
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.maintainersVerifyAuxiliary = value
			}
		default:
			panic(errorConstants.UninitializedUsage)
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
