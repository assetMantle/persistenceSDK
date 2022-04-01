// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package super

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Auxiliary = base.NewAuxiliary(
	"super",
	keeperPrototype,
)

var AuxiliaryMock = base.NewAuxiliary(
	"super",
	keeperPrototypeMock,
)
