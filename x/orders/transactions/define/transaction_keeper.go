// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/define"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/super"
	"github.com/AssetMantle/modules/x/orders/utilities"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties/constants"
	"github.com/AssetMantle/schema/qualified/base"
)

type transactionKeeper struct {
	mapper                helpers.Mapper
	parameterManager      helpers.ParameterManager
	defineAuxiliary       helpers.Auxiliary
	superAuxiliary        helpers.Auxiliary
	authenticateAuxiliary helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context context.Context, message helpers.Message) (helpers.TransactionResponse, error) {
	return transactionKeeper.Handle(context, message.(*Message))
}

func (transactionKeeper transactionKeeper) Handle(context context.Context, message *Message) (*TransactionResponse, error) {
	if _, err := transactionKeeper.authenticateAuxiliary.GetKeeper().Help(context, authenticate.NewAuxiliaryRequest(message)); err != nil {
		return nil, err
	}

	immutables := base.NewImmutables(
		message.ImmutableMetaProperties.Add(
			baseLists.AnyPropertiesToProperties(
				message.ImmutableProperties.Add(
					constants.ExchangeRateProperty.ToAnyProperty(),
					constants.CreationHeightProperty.ToAnyProperty(),
					constants.MakerAssetIDProperty.ToAnyProperty(),
					constants.TakerAssetIDProperty.ToAnyProperty(),
					constants.MakerIDProperty.ToAnyProperty(),
					constants.TakerIDProperty.ToAnyProperty(),
				).Get()...,
			)...,
		),
	)

	mutables := base.NewMutables(
		message.MutableMetaProperties.Add(
			baseLists.AnyPropertiesToProperties(
				message.MutableProperties.Add(
					constants.ExpiryHeightProperty.ToAnyProperty(),
					constants.MakerSplitProperty.ToAnyProperty(),
				).Get()...,
			)...,
		),
	)

	auxiliaryResponse, err := transactionKeeper.defineAuxiliary.GetKeeper().Help(context, define.NewAuxiliaryRequest(message.GetFromAddress(), immutables, mutables))
	if err != nil {
		return nil, err
	}
	classificationID := define.GetClassificationIDFromResponse(auxiliaryResponse)

	if _, err := transactionKeeper.superAuxiliary.GetKeeper().Help(context, super.NewAuxiliaryRequest(classificationID, message.GetFromIdentityID(), mutables, utilities.SetModulePermissions(true, true)...)); err != nil {
		return nil, err
	}

	return newTransactionResponse(classificationID), nil
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, parameterManager helpers.ParameterManager, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper, transactionKeeper.parameterManager = mapper, parameterManager

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case define.Auxiliary.GetName():
				transactionKeeper.defineAuxiliary = value
			case super.Auxiliary.GetName():
				transactionKeeper.superAuxiliary = value
			case authenticate.Auxiliary.GetName():
				transactionKeeper.authenticateAuxiliary = value
			}
		}
	}

	helpers.PanicOnUninitializedKeeperFields(transactionKeeper)
	return transactionKeeper
}
func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
