package invariants

import (
	"github.com/AssetMantle/schema/x/helpers"
	"github.com/AssetMantle/schema/x/helpers/base"

	"github.com/AssetMantle/modules/modules/maintainers/internal/invariants/dummy"
	"github.com/AssetMantle/modules/modules/maintainers/internal/module"
)

func Prototype() helpers.Invariants {
	return base.NewInvariants(module.Name, "", dummy.Invariant)
}
