// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package parameters

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/schema/helpers"

	"github.com/AssetMantle/modules/modules/assets/internal/parameters/dummy"
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name      string
		want      helpers.ParameterList
		wantError error
	}{
		{"+ve", baseHelpers.NewParameterList(dummy.ValidatableParameter), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
