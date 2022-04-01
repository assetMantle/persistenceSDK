// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type property struct {
	ID     propertyID `json:"id"`
	DataID dataID     `json:"dataID"`
}

var _ types.Property = (*property)(nil)

func (property property) GetID() types.ID {
	return property.ID
}
func (property property) GetDataID() types.ID {
	return property.DataID
}
func (property property) GetKeyID() types.ID {
	return property.ID.KeyID
}
func (property property) GetTypeID() types.ID {
	return property.ID.TypeID
}
func (property property) GetHashID() types.ID {
	return property.DataID.HashID
}

func NewProperty(keyID types.ID, data types.Data) types.Property {
	return property{
		ID: propertyID{
			KeyID:  keyID,
			TypeID: data.GetTypeID(),
		},
		DataID: dataIDFromInterface(data.GetID()),
	}
}

func ReadProperty(propertyString string) (types.Property, error) {
	property, err := ReadMetaProperty(propertyString)
	if err != nil {
		return nil, err
	}

	return property.RemoveData(), nil
}
