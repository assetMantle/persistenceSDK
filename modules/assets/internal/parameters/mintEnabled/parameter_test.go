// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mintEnabled

import (
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
	"testing"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	baseTypes "github.com/AssetMantle/modules/schema/parameters/base"
)

func Test_validator(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantError bool
	}{
		{"-ve incorrectFormat", args{baseIDs.NewStringID("")}, true},
		{"+ve", args{Parameter}, false},
		{"-ve InvalidParameter", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID(""), baseData.NewStringData("")))}, true},
		{"+ve with booleanData", args{baseData.NewBooleanData(false)}, false},
		{"-ve with different type of Data", args{baseData.NewStringData("stringData")}, true},
		{"+ve with true booleanData", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("mintEnabled"), baseData.NewBooleanData(true)))}, false},
		{"+ve with false booleanData", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("mintEnabled"), baseData.NewBooleanData(false)))}, false},
		{"+ve with incorrect ID", args{baseTypes.NewParameter(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID"), baseData.NewBooleanData(false)))}, true},
		{"-ve nil", args{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validator(tt.args.i); (err != nil) != tt.wantError {
				t.Errorf("validator() error = %v, wantErr %v", err, tt.wantError)
			}
		})
	}
}
