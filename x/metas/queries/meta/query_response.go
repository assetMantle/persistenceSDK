// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package meta

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/metas/record"
)

var _ helpers.QueryResponse = (*QueryResponse)(nil)

func responsePrototype() helpers.QueryResponse {
	return &QueryResponse{}
}
func newQueryResponse(Record helpers.Record) *QueryResponse {
	return &QueryResponse{
		Record: Record.(*record.Record),
	}
}
