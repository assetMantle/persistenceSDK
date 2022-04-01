// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package identity

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Query = base.NewQuery(
	"identities",
	"",
	"",

	module.Name,

	requestPrototype,
	responsePrototype,
	keeperPrototype,

	flags.IdentityID,
)
