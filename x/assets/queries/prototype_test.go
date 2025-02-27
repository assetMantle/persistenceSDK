// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queries

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/x/assets/queries/asset"
)

func TestPrototype(t *testing.T) {
	tests := []struct {
		name      string
		want      string
		getString string
	}{

		{"+ve", asset.Query.GetServicePath(), "assets"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Prototype(); !reflect.DeepEqual(got.GetQuery(tt.getString).GetServicePath(), tt.want) {
				t.Errorf("Prototype() = %v, want %v", got, tt.want)
			}
		})
	}
}
