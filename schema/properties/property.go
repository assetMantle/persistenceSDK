// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type Property interface {
	GetID() ids.ID
	GetDataID() ids.ID
	GetKey() ids.ID
	GetType() ids.ID

	IsMeta() bool

	traits.Listable
}
