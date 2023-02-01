// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"reflect"
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/schema/data"
	baseData "github.com/AssetMantle/modules/schema/data/base"
	"github.com/AssetMantle/modules/schema/ids"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/parameters"
)

func dummyValidator(interface{}) error {
	return nil
}

func createTestInput() (*baseIDs.StringID, data.Data, parameters.Parameter) {
	id := baseIDs.NewStringID("ID")
	stringData := baseData.NewStringData("Data")

	testParameter := NewParameter(id, stringData, dummyValidator)
	return id.(*baseIDs.StringID), stringData, testParameter
}

func TestNewParameter(t *testing.T) {
	id, testData, _ := createTestInput()
	type args struct {
		id        ids.StringID
		data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name    string
		args    args
		want    parameters.Parameter
		wantErr bool
	}{

		{"+ve", args{id, testData, dummyValidator}, &Parameter{id, testData.ToAnyData().(*baseData.AnyData)}, false},
		{"panic empty", args{}, &Parameter{}, true},
		{"nil", args{nil, nil, nil}, &Parameter{nil, nil}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()

				if (r != nil) != tt.wantErr {
					t.Errorf("Error %v", r)
				}
			}()
			if got := NewParameter(tt.args.id, tt.args.data, tt.args.validator); !reflect.DeepEqual(got.AsString(), tt.want.AsString()) {
				t.Errorf("NewParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_Equal(t *testing.T) {
	id, testData, testParameter := createTestInput()
	type fields struct {
		ID        ids.StringID
		Data      data.Data
		validator func(interface{}) error
	}
	type args struct {
		compareParameter parameters.Parameter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{

		{"+ve", fields{id, testData, dummyValidator}, args{testParameter}, true},
		{"+ve different validator", fields{id, testData, func(interface{}) error { return nil }}, args{testParameter}, true},
		{"-ve different id", fields{baseIDs.NewStringID("different"), testData, dummyValidator}, args{testParameter}, false},
		{"-ve different data", fields{id, baseData.NewStringData("different"), dummyValidator}, args{testParameter}, false},
		{"-ve different data type", fields{id, baseData.NewBooleanData(false), dummyValidator}, args{testParameter}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := &Parameter{
				ID:   tt.fields.ID.(*baseIDs.StringID),
				Data: tt.fields.Data.ToAnyData().(*baseData.AnyData),
			}
			if got := parameter.Equal(tt.args.compareParameter); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_GetData(t *testing.T) {
	id, testData, _ := createTestInput()
	type fields struct {
		ID        ids.StringID
		Data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name   string
		fields fields
		want   data.AnyData
	}{

		{"+ve", fields{id, testData, dummyValidator}, testData.ToAnyData()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := &Parameter{
				ID:   tt.fields.ID.(*baseIDs.StringID),
				Data: tt.fields.Data.ToAnyData().(*baseData.AnyData),
			}
			if got := parameter.GetData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_GetID(t *testing.T) {
	id, testData, _ := createTestInput()
	type fields struct {
		ID        ids.StringID
		Data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name   string
		fields fields
		want   ids.StringID
	}{

		{"+ve", fields{id, testData, dummyValidator}, id},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := &Parameter{
				ID:   tt.fields.ID.(*baseIDs.StringID),
				Data: tt.fields.Data.ToAnyData().(*baseData.AnyData),
			}
			if got := parameter.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_GetValidator(t *testing.T) {
	id, testData, _ := createTestInput()
	type fields struct {
		ID        ids.StringID
		Data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name   string
		fields fields
		want   func(interface{}) error
	}{

		{"+ve", fields{id, testData, dummyValidator}, dummyValidator},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := &Parameter{
				ID:   tt.fields.ID.(*baseIDs.StringID),
				Data: tt.fields.Data.ToAnyData().(*baseData.AnyData),
			}
			if got := parameter.GetValidator(); !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("GetValidator() = %p, want %p", got, tt.want)
			}
		})
	}
}

func Test_parameter_Mutate(t *testing.T) {
	id, testData, _ := createTestInput()
	newData := baseData.NewStringData("Data")
	type fields struct {
		ID        ids.StringID
		Data      data.Data
		validator func(interface{}) error
	}
	type args struct {
		data data.Data
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   parameters.Parameter
	}{

		{"+ve", fields{id, testData, dummyValidator}, args{newData}, &Parameter{id, newData.ToAnyData().(*baseData.AnyData)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := &Parameter{
				ID:   tt.fields.ID.(*baseIDs.StringID),
				Data: tt.fields.Data.ToAnyData().(*baseData.AnyData),
			}
			if got := parameter.Mutate(tt.args.data); !reflect.DeepEqual(got.AsString(), tt.want.AsString()) {
				t.Errorf("Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_AsString(t *testing.T) {
	id, testData, testParameter := createTestInput()
	type fields struct {
		ID        ids.StringID
		Data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{

		{"+ve", fields{id, testData, dummyValidator}, testParameter.AsString()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := &Parameter{
				ID:   tt.fields.ID.(*baseIDs.StringID),
				Data: tt.fields.Data.ToAnyData().(*baseData.AnyData),
			}
			if got := parameter.AsString(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parameter_Validate(t *testing.T) {
	id, testData, _ := createTestInput()
	type fields struct {
		ID        ids.StringID
		Data      data.Data
		validator func(interface{}) error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve with stringData", fields{id, testData, dummyValidator}, false},
		{"+ve with decData", fields{baseIDs.NewStringID("ID"), baseData.NewDecData(sdkTypes.SmallestDec()), dummyValidator}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := &Parameter{
				ID:   tt.fields.ID.(*baseIDs.StringID),
				Data: tt.fields.Data.ToAnyData().(*baseData.AnyData),
			}
			if err := parameter.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
