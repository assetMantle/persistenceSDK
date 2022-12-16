package base

import (
	"github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/base"
	"github.com/AssetMantle/modules/schema/traits"
)

type propertyList struct {
	lists.List
}

var _ lists.PropertyList = (*propertyList)(nil)

func (propertyList propertyList) GetProperty(propertyID ids.PropertyID) properties.Property {
	if sanitizedPropertyList, err := propertyList.Sanitize(); err == nil {
		propertyList.List = NewList(propertiesToListables(sanitizedPropertyList.(lists.PropertyList).GetList()...)...)
	}
	switch propertyID.(type) {
	case ids.PropertyID:
		propertyID = propertyID.(ids.PropertyID)
	default:
		panic(constants.MetaDataError)
	}
	if i, found := propertyList.Search(base.NewEmptyMesaPropertyFromID(propertyID)); found {
		return propertyList.GetList()[i]
	}

	return nil
}
func (propertyList propertyList) GetList() []properties.Property {
	if propertyList.List == nil {
		propertyList.List = NewList()
	}
	Properties := make([]properties.Property, propertyList.List.Size())
	for i, listable := range propertyList.List.Get() {
		Properties[i] = listable.(properties.Property)
	}
	return Properties
}
func (propertyList propertyList) GetPropertyIDList() lists.IDList {
	if sanitizedPropertyList, err := propertyList.Sanitize(); err == nil {
		propertyList.List = NewList(propertiesToListables(sanitizedPropertyList.(lists.PropertyList).GetList()...)...)
	}
	propertyIDList := NewIDList()
	for _, property := range propertyList.GetList() {
		propertyIDList = propertyIDList.Add(property.GetID())
	}
	return propertyIDList
}
func (propertyList propertyList) Add(properties ...properties.Property) lists.PropertyList {
	if sanitizedPropertyList, err := propertyList.Sanitize(); err == nil {
		propertyList.List = NewList(propertiesToListables(sanitizedPropertyList.(lists.PropertyList).GetList()...)...)
	}
	propertyList.List = propertyList.List.Add(propertiesToListables(properties...)...)
	return propertyList
}
func (propertyList propertyList) Remove(properties ...properties.Property) lists.PropertyList {
	if sanitizedPropertyList, err := propertyList.Sanitize(); err == nil {
		propertyList.List = NewList(propertiesToListables(sanitizedPropertyList.(lists.PropertyList).GetList()...)...)
	}
	propertyList.List = propertyList.List.Remove(propertiesToListables(properties...)...)
	return propertyList
}
func (propertyList propertyList) Mutate(properties ...properties.Property) lists.PropertyList {
	if sanitizedPropertyList, err := propertyList.Sanitize(); err == nil {
		propertyList.List = NewList(propertiesToListables(sanitizedPropertyList.(lists.PropertyList).GetList()...)...)
	}
	propertyList.List = propertyList.List.Mutate(propertiesToListables(properties...)...)
	return propertyList
}
func (propertyList propertyList) ScrubData() lists.PropertyList {
	if sanitizedPropertyList, err := propertyList.Sanitize(); err == nil {
		propertyList.List = NewList(propertiesToListables(sanitizedPropertyList.(lists.PropertyList).GetList()...)...)
	}
	newPropertyList := NewPropertyList()
	for _, listable := range propertyList.List.Get() {
		if property := listable.(properties.Property); property.IsMeta() {
			newPropertyList = newPropertyList.Add(property.(properties.MetaProperty).ScrubData())
		} else {
			newPropertyList = newPropertyList.Add(property)
		}
	}
	return newPropertyList
}

func (propertyList propertyList) Sanitize() (lists.PropertyList, error) {
	if propertyList.List == nil {
		propertyList.List = NewList()
	}
	return propertyList, nil
}

func propertiesToListables(properties ...properties.Property) []traits.Listable {
	listables := make([]traits.Listable, len(properties))
	for i, property := range properties {
		listables[i] = property
	}
	return listables
}

func NewPropertyList(properties ...properties.Property) lists.PropertyList {
	return propertyList{List: NewList(propertiesToListables(properties...)...)}
}
