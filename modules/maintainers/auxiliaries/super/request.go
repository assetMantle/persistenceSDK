// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package super

import (
	"github.com/asaskevich/govalidator"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/types"
)

type auxiliaryRequest struct {
	ClassificationID  types.ID           `json:"classificationID" valid:"required~required field classificationID missing"`
	IdentityID        types.ID           `json:"identityID" valid:"required~required field identityID missing"`
	MutableProperties lists.PropertyList `json:"mutableProperties" valid:"required~required field mutableProperties missing"`
}

var _ helpers.AuxiliaryRequest = (*auxiliaryRequest)(nil)

func (auxiliaryRequest auxiliaryRequest) Validate() error {
	_, err := govalidator.ValidateStruct(auxiliaryRequest)
	return err
}

func auxiliaryRequestFromInterface(request helpers.AuxiliaryRequest) auxiliaryRequest {
	switch value := request.(type) {
	case auxiliaryRequest:
		return value
	default:
		return auxiliaryRequest{}
	}
}

func NewAuxiliaryRequest(classificationID types.ID, identityID types.ID, mutableProperties lists.PropertyList) helpers.AuxiliaryRequest {
	return auxiliaryRequest{
		ClassificationID:  classificationID,
		IdentityID:        identityID,
		MutableProperties: mutableProperties,
	}
}
