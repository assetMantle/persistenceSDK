// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package define

import (
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ helpers.TransactionResponse = (*TransactionResponse)(nil)

func (transactionResponse *TransactionResponse) GetResult() *sdkTypes.Result {
	return &sdkTypes.Result{
		Data: []byte(transactionResponse.IdentityID.AsString()),
	}
}
func newTransactionResponse(classificationID ids.ClassificationID) *TransactionResponse {
	return &TransactionResponse{classificationID.(*baseIDs.ClassificationID)}
}
