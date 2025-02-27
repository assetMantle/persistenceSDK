// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintainers

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/maintainers/record"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(collection helpers.Collection) *QueryResponse {
	return &QueryResponse{helpers.RecordsToImplementations(&record.Record{}, collection.Get())}
}
