// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package simulator

import (
	"math/rand"

	"github.com/AssetMantle/schema/x/data"
	baseData "github.com/AssetMantle/schema/x/data/base"
	baseParameters "github.com/AssetMantle/schema/x/parameters/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/AssetMantle/modules/helpers"
	baseSimulation "github.com/AssetMantle/modules/simulation/schema/types/base"
	"github.com/AssetMantle/modules/x/metas/internal/common"
	"github.com/AssetMantle/modules/x/metas/internal/genesis"
	"github.com/AssetMantle/modules/x/metas/internal/mappable"
	metasModule "github.com/AssetMantle/modules/x/metas/internal/module"
	"github.com/AssetMantle/modules/x/metas/internal/parameters/revealEnabled"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {
	var Data data.Data

	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		revealEnabled.ID.AsString(),
		&Data,
		simulationState.Rand,
		func(rand *rand.Rand) { Data = baseData.NewDecData(sdkTypes.NewDecWithPrec(int64(rand.Intn(99)), 2)) },
	)

	mappableList := make([]helpers.Mappable, simulationState.Rand.Intn(99))

	for i := range mappableList {
		mappableList[i] = mappable.NewMappable(baseSimulation.GenerateRandomData(simulationState.Rand))
	}

	genesisState := genesis.Prototype().Initialize(mappableList, baseParameters.NewParameterList(revealEnabled.Parameter.Mutate(Data)))

	simulationState.GenState[metasModule.Name] = common.LegacyAmino.MustMarshalJSON(genesisState)
}
