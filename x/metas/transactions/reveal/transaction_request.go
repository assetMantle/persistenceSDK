// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package reveal

import (
	"encoding/json"
	codecUtilities "github.com/AssetMantle/schema/codec/utilities"
	"github.com/AssetMantle/schema/data/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
)

type transactionRequest struct {
	helpers.CommonTransactionRequest `json:"commonTransactionRequest"`
	Data                             string `json:"data"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

// Validate godoc
// @Summary Reveal metas transaction
// @Description Reveal metas transaction
// @Accept text/plain
// @Produce json
// @Tags Metas
// @Param body  transactionRequest true "Request body to reveal meta transaction"
// @Success 200 {object} transactionResponse   "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /metas/reveal [post]
func (transactionRequest transactionRequest) Validate() error {
	if msg, err := transactionRequest.MakeMsg(); err != nil {
		return err
	} else if err := msg.(helpers.Message).ValidateBasic(); err != nil {
		return err
	}

	return nil
}
func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadCommonTransactionRequest(context),
		cliCommand.ReadString(constants.Data),
	), nil
}
func (transactionRequest transactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if err := json.Unmarshal(rawMessage, &transactionRequest); err != nil {
		return nil, err
	}

	return transactionRequest, nil
}
func (transactionRequest transactionRequest) GetCommonTransactionRequest() helpers.CommonTransactionRequest {
	return transactionRequest.CommonTransactionRequest
}
func (transactionRequest transactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, err := sdkTypes.AccAddressFromBech32(transactionRequest.GetCommonTransactionRequest().GetFrom())
	if err != nil {
		return nil, err
	}

	data, err := base.PrototypeAnyData().FromString(transactionRequest.Data)
	if err != nil {
		return nil, err
	}

	return NewMessage(
		from,
		data,
	), nil
}
func (transactionRequest) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, transactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}
func newTransactionRequest(commonTransactionRequest helpers.CommonTransactionRequest, data string) helpers.TransactionRequest {
	return transactionRequest{
		CommonTransactionRequest: commonTransactionRequest,
		Data:                     data,
	}
}
