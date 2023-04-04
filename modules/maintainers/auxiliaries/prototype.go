// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package auxiliaries

import (
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/deputize"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/maintain"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/revoke"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/super"
	"github.com/AssetMantle/modules/modules/maintainers/auxiliaries/verify"
)

func Prototype() helpers.Auxiliaries {
	return baseHelpers.NewAuxiliaries(
		deputize.Auxiliary,
		maintain.Auxiliary,
		revoke.Auxiliary,
		super.Auxiliary,
		verify.Auxiliary,
	)
}
