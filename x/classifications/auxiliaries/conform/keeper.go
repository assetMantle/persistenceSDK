// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package conform

import (
	"context"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/classifications/internal/key"
	"github.com/AssetMantle/modules/x/classifications/internal/mappable"
)

type auxiliaryKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	auxiliaryRequest := auxiliaryRequestFromInterface(request)
	classifications := auxiliaryKeeper.mapper.NewCollection(context).Fetch(key.NewKey(auxiliaryRequest.ClassificationID))

	Mappable := classifications.Get(key.NewKey(auxiliaryRequest.ClassificationID))
	if Mappable == nil {
		return nil, errorConstants.EntityNotFound.Wrapf("classification with ID %s not found", auxiliaryRequest.ClassificationID.AsString())
	}
	classification := mappable.GetClassification(Mappable)

	if auxiliaryRequest.Immutables != nil {
		if len(auxiliaryRequest.Immutables.GetImmutablePropertyList().GetList()) != len(classification.GetImmutables().GetImmutablePropertyList().GetList()) {
			return nil, errorConstants.IncorrectFormat
		}

		for _, immutableProperty := range classification.GetImmutables().GetImmutablePropertyList().GetList() {
			if property := auxiliaryRequest.Immutables.GetImmutablePropertyList().GetProperty(immutableProperty.GetID()); property == nil || immutableProperty.GetDataID().GetHashID().Compare(baseIDs.GenerateHashID()) != 0 && property.GetDataID().GetHashID().Compare(immutableProperty.GetDataID().GetHashID()) != 0 {
				return nil, errorConstants.IncorrectFormat
			}
		}
	}

	if auxiliaryRequest.Mutables != nil {
		if len(auxiliaryRequest.Mutables.GetMutablePropertyList().GetList()) != len(classification.GetMutables().GetMutablePropertyList().GetList()) {
			return nil, errorConstants.IncorrectFormat
		}

		for _, mutableProperty := range classification.GetMutables().GetMutablePropertyList().GetList() {
			if property := auxiliaryRequest.Mutables.GetMutablePropertyList().GetProperty(mutableProperty.GetID()); property == nil {
				return nil, errorConstants.IncorrectFormat
			}
		}
	}

	return newAuxiliaryResponse(), nil
}

func (auxiliaryKeeper) Initialize(mapper helpers.Mapper, _ helpers.ParameterManager, _ []interface{}) helpers.Keeper {
	return auxiliaryKeeper{mapper: mapper}
}

func keeperPrototype() helpers.AuxiliaryKeeper {
	return auxiliaryKeeper{}
}
