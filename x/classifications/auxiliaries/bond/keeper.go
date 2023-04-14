// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package bond

import (
	"context"

	errorConstants "github.com/AssetMantle/schema/x/errors/constants"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/internal/key"
	"github.com/AssetMantle/modules/x/classifications/internal/mappable"
	"github.com/AssetMantle/modules/x/classifications/internal/module"
)

type auxiliaryKeeper struct {
	mapper        helpers.Mapper
	bankKeeper    bankKeeper.Keeper
	stakingKeeper stakingKeeper.Keeper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)

	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.classificationID))

	Mappable := classifications.Get(key.NewKey(auxiliaryRequest.classificationID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("classification with ID %s not found", auxiliaryRequest.classificationID.AsString())
	}
	classification := mappable.GetClassification(Mappable)

	if err := auxiliaryKeeper.bankKeeper.SendCoinsFromAccountToModule(sdkTypes.UnwrapSDKContext(context), auxiliaryRequest.accAddress, module.Name, sdkTypes.NewCoins(sdkTypes.NewCoin(auxiliaryKeeper.stakingKeeper.BondDenom(sdkTypes.UnwrapSDKContext(context)), sdkTypes.NewInt(classification.GetBondAmount())))); err != nil {
		return nil, err
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	auxiliaryKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case bankKeeper.Keeper:
			auxiliaryKeeper.bankKeeper = value
		case stakingKeeper.Keeper:
			auxiliaryKeeper.stakingKeeper = value
		}
	}
	return auxiliaryKeeper
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
