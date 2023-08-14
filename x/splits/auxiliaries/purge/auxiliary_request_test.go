// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package purge

import (
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
)

func createTestInput1() (ids.IdentityID, ids.AssetID, types.Int) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	testOwnerID := baseIDs.NewIdentityID(classificationID, immutables)
	testAssetID := base.GenerateCoinAssetID("OwnerID")
	testValue := types.OneInt()
	return testOwnerID, testAssetID, testValue
}

func TestNewAuxiliaryRequest(t *testing.T) {
	testOwnerID, testAssetID, testValue := createTestInput1()
	type args struct {
		ownerID ids.IdentityID
		assetID ids.AssetID
		value   types.Int
	}
	tests := []struct {
		name string
		args args
		want helpers.AuxiliaryRequest
	}{
		{"+ve", args{testOwnerID, testAssetID, testValue}, NewAuxiliaryRequest(testOwnerID, testAssetID, testValue)},

		{"+ve with nil", args{baseIDs.PrototypeIdentityID(), baseIDs.PrototypeAssetID(), testValue}, NewAuxiliaryRequest(baseIDs.PrototypeIdentityID(), baseIDs.PrototypeAssetID(), testValue)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuxiliaryRequest(tt.args.ownerID, tt.args.assetID, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuxiliaryRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaryRequestFromInterface(t *testing.T) {
	testOwnerID, testAssetID, testValue := createTestInput1()
	type args struct {
		request helpers.AuxiliaryRequest
	}
	tests := []struct {
		name string
		args args
		want auxiliaryRequest
	}{
		{"+ve", args{NewAuxiliaryRequest(testOwnerID, testAssetID, testValue)}, auxiliaryRequest{testOwnerID, testAssetID, testValue}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := auxiliaryRequestFromInterface(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("auxiliaryRequestFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_auxiliaryRequest_Validate(t *testing.T) {
	testOwnerID, testAssetID, testValue := createTestInput1()
	type fields struct {
		OwnerID ids.IdentityID
		AssetID ids.AssetID
		Value   types.Int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testOwnerID, testAssetID, testValue}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auxiliaryRequest := auxiliaryRequest{
				OwnerID: tt.fields.OwnerID,
				AssetID: tt.fields.AssetID,
				Value:   tt.fields.Value,
			}
			if err := auxiliaryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
