// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unwrap

import (
	"context"

	"github.com/AssetMantle/schema/go/data/base"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	constantProperties "github.com/AssetMantle/schema/go/properties/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/burn"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	bankKeeper            bankKeeper.Keeper
	burnAuxiliary         helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	fromAddress := message.GetFromAddress()

	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(fromAddress, message.FromID)); err != nil {
		return nil, err
	}

	if err := transactionKeeper.bankKeeper.SendCoinsFromModuleToAccount(sdkTypes.UnwrapSDKContext(context), constants.ModuleName, fromAddress, message.Coins); err != nil {
		return nil, err
	}

	for _, coin := range message.Coins {
		if err := sdkTypes.ValidateDenom(coin.Denom); err != nil {
			return nil, errorConstants.InvalidRequest.Wrapf("coin denom %s is invalid", coin.Denom)
		}

		if _, found := transactionKeeper.parameterManager.Fetch(context).GetParameter(constantProperties.UnwrapAllowedCoinsProperty.GetID()).GetMetaProperty().GetData().Get().(*base.ListData).Search(base.NewStringData(coin.Denom)); !found {
			return nil, errorConstants.NotAuthorized.Wrapf("coin %s is not allowed to be unwrapped", coin.Denom)
		}

		coinAssetID := baseDocuments.NewCoinAsset(coin.Denom).GetCoinAssetID()

		Mappable := transactionKeeper.mapper.NewCollection(context).Fetch(key.NewKey(coinAssetID)).GetMappable(key.NewKey(coinAssetID))
		if Mappable == nil {
			return nil, errorConstants.EntityNotFound.Wrapf("asset %s not found", coinAssetID)
		}

		value := coin.Amount
		if value.IsNegative() {
			return nil, errorConstants.InvalidRequest.Wrapf("coin amount %s is negative", value)
		}

		if _, err := transactionKeeper.burnAuxiliary.GetKeeper().Help(context, burn.NewAuxiliaryRequest(message.FromID, coinAssetID, value)); err != nil {
			return nil, err
		}
	}

	return newTransactionResponse(), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case bankKeeper.Keeper:
			transactionKeeper.bankKeeper = value
		case helpers.Auxiliary:
			switch value.GetName() {
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			case burn.Auxiliary.GetName():
				transactionKeeper.burnAuxiliary = value
			}
		}
	}

	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
