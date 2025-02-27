// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/helpers"
	baseData "github.com/AssetMantle/schema/data/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
)

var (
	testDataID  = baseIDs.GenerateDataID(baseData.NewStringData("Data")).(*baseIDs.DataID)
	testDataID1 = baseIDs.PrototypeDataID().(*baseIDs.DataID)
)

func TestNewKey(t *testing.T) {
	type args struct {
		dataID ids.DataID
	}
	tests := []struct {
		name string
		args args
		want helpers.Key
	}{
		{"+ve", args{testDataID}, &Key{testDataID}},
		{"+ve", args{testDataID1}, &Key{testDataID1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKey(tt.args.dataID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Key
	}{
		{"+ve", &Key{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_Equals(t *testing.T) {
	type fields struct {
		DataID *baseIDs.DataID
	}
	type args struct {
		compareKey helpers.Key
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"+ve", fields{testDataID}, args{NewKey(testDataID)}, true},
		{"+ve", fields{testDataID}, args{NewKey(testDataID1)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				DataID: tt.fields.DataID,
			}
			if got := key.Equals(tt.args.compareKey); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_GenerateStoreKeyBytes(t *testing.T) {
	type fields struct {
		DataID *baseIDs.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{testDataID}, (&Key{testDataID}).GeneratePrefixedStoreKeyBytes()},
		{"+ve with nil", fields{testDataID1}, (&Key{testDataID1}).GeneratePrefixedStoreKeyBytes()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				DataID: tt.fields.DataID,
			}
			if got := key.GeneratePrefixedStoreKeyBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateStoreKeyGenerateStoreKeyBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_key_IsPartial(t *testing.T) {
	type fields struct {
		DataID *baseIDs.DataID
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"+ve", fields{testDataID}, false},
		{"-ve", fields{testDataID1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key := &Key{
				DataID: tt.fields.DataID,
			}
			if got := key.IsPartial(); got != tt.want {
				t.Errorf("IsPartial() = %v, want %v", got, tt.want)
			}
		})
	}
}
