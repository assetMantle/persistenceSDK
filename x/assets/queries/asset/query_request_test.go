// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package asset

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	baseLists "github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/viper"
	"reflect"
	"testing"
)

var (
	immutables       = baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData"))))
	mutables         = baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("MutableData"))))
	classificationID = baseIDs.NewClassificationID(immutables, mutables)
	testAssetID      = baseIDs.NewAssetID(classificationID, immutables).(*baseIDs.AssetID)
	testKey          = key.NewKey(testAssetID).(*key.Key)
)

func Test_newQueryRequest(t *testing.T) {
	type args struct {
		assetID ids.AssetID
	}
	tests := []struct {
		name string
		args args
		want helpers.QueryRequest
	}{
		{"+ve", args{testAssetID}, newQueryRequest(testAssetID)},
		{"+ve", args{baseIDs.PrototypeAssetID().(*baseIDs.AssetID)}, newQueryRequest(baseIDs.PrototypeAssetID().(*baseIDs.AssetID))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueryRequest(tt.args.assetID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQueryRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_FromCLI(t *testing.T) {
	cliCommand := base.NewCLICommand("", "", "", []helpers.CLIFlag{constants.AssetID})

	viper.Set(constants.AssetID.GetName(), testAssetID.AsString())
	type fields struct {
		Key *key.Key
	}
	type args struct {
		cliCommand helpers.CLICommand
		context    client.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.QueryRequest
		wantErr bool
	}{
		{"+ve", fields{testKey}, args{cliCommand, client.Context{}.WithCodec(base.CodecPrototype())}, newQueryRequest(testAssetID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qu := &QueryRequest{
				Key: tt.fields.Key,
			}
			got, err := qu.FromCLI(tt.args.cliCommand, tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromCLI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryRequest_Validate(t *testing.T) {
	type fields struct {
		Key *key.Key
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{testKey}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queryRequest := &QueryRequest{
				Key: tt.fields.Key,
			}
			if err := queryRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_requestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.QueryRequest
	}{
		{"+ve", &QueryRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := requestPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
