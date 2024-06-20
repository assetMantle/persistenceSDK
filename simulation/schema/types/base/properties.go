// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"math"
	"math/rand"

	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/properties"
)

func GenerateRandomPropertyList(r *rand.Rand) lists.PropertyList {
	randomPositiveInt := int(math.Abs(float64(r.Int()))) % 11

	propertyList := make([]properties.Property, randomPositiveInt)
	for i := 0; i < randomPositiveInt; i++ {
		propertyList[i] = GenerateRandomProperty(r)
	}
	return baseLists.NewPropertyList(propertyList...)
}

func GenerateRandomMetaPropertyList(r *rand.Rand) lists.PropertyList {
	randomPositiveInt := int(math.Abs(float64(r.Int()))) % 11

	propertyList := make([]properties.Property, randomPositiveInt)

	for i := 0; i < randomPositiveInt; i++ {
		propertyList[i] = GenerateRandomMetaProperty(r)
	}

	return baseLists.NewPropertyList(propertyList...)
}
func GenerateRandomMetaPropertyListWithoutData(r *rand.Rand) lists.PropertyList {
	randomPositiveInt := int(math.Abs(float64(r.Int()))) % 11

	propertyList := make([]properties.Property, randomPositiveInt)

	for i := 0; i < randomPositiveInt; i++ {
		propertyList[i] = GenerateRandomMetaPropertyWithoutData(r)
	}

	return baseLists.NewPropertyList(propertyList...)
}
