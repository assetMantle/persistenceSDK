/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"strings"
)

var _ types.Property = (*property)(nil)

type property struct {
	ID   types.ID   `json:"id"`
	Fact types.Fact `json:"fact"`
}

func (property property) GetID() types.ID     { return property.ID }
func (property property) GetFact() types.Fact { return property.Fact }
func NewProperty(id types.ID, fact types.Fact) types.Property {
	return property{
		ID:   id,
		Fact: fact,
	}
}
func ReadProperty(PropertyIDAndStringData string) types.Property {
	propertyIDAndFactList := strings.Split(PropertyIDAndStringData, constants.PropertyIDAndDataSeparator)
	if len(propertyIDAndFactList) == 2 && propertyIDAndFactList[0] != "" {
		return NewProperty(NewID(propertyIDAndFactList[0]), NewFact(NewStringData(propertyIDAndFactList[1])))
	}
	return nil
}
