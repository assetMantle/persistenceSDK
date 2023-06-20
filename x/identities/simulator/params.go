// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math"
	"math/rand"

	"github.com/AssetMantle/schema/go/data"
	simulationTypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/identities/constants"
	"github.com/AssetMantle/modules/x/identities/parameters/max_provision_address_count"
)

func (simulator) ParamChangeList(_ *rand.Rand) []simulationTypes.ParamChange {
	currentMaxProvisionAddressCount := max_provision_address_count.Parameter.GetMetaProperty().GetData().Get().(data.NumberData).Get()
	return []simulationTypes.ParamChange{
		simulation.NewSimParamChange(constants.ModuleName,
			string(max_provision_address_count.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(rand.Intn(math.MaxInt))
				if err != nil {
					panic(err)
				}
				return string(bytes)
			},
		),
		simulation.NewSimParamChange(constants.ModuleName,
			string(max_provision_address_count.Parameter.GetMetaProperty().GetID().Bytes()),
			func(r *rand.Rand) string {
				bytes, err := baseHelpers.CodecPrototype().GetLegacyAmino().MarshalJSON(currentMaxProvisionAddressCount)
				if err != nil {
					panic(err)
				}
				return string(bytes)
			},
		),
	}
}
