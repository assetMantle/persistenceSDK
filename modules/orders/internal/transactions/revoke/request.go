// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/helpers/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

//type transactionRequest struct {
//	BaseReq          rest.BaseReq `json:"baseReq"`
//	FromID           string       `json:"fromID" valid:"required~required field fromID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field fromID"`
//	ToID             string       `json:"toID" valid:"required~required field toID missing, matches(^[A-Za-z0-9-_=.|]+$)~invalid field toID"`
//	ClassificationID string       `json:"classificationID" valid:"required~required field classificationID missing, matches(^[A-Za-z0-9-_=.]+$)~invalid field classificationID"`
//}

var _ helpers.TransactionRequest = (*TransactionRequest)(nil)

// Validate godoc
// @Summary Revoke order transaction
// @Description Revoke order transaction
// @Accept text/plain
// @Produce json
// @Tags Orders
// @Param body transactionRequest true "Request body to revoke order"
// @Success 200 {object} transactionResponse "Message for a successful response."
// @Failure default  {object}  transactionResponse "Message for an unexpected error response."
// @Router /orders/revoke [post]
func (transactionRequest *TransactionRequest) Validate() error {
	_, err := govalidator.ValidateStruct(transactionRequest)
	return err
}
func (transactionRequest *TransactionRequest) RegisterInterface(registry types.InterfaceRegistry) {
	//TODO implement me
	panic("implement me")
}
func (transactionRequest *TransactionRequest) FromCLI(cliCommand helpers.CLICommand, context client.Context) (helpers.TransactionRequest, error) {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(context).From,
		cliCommand.ReadString(constants.FromID),
		cliCommand.ReadString(constants.ToID),
		cliCommand.ReadString(constants.ClassificationID),
	), nil
}
func (transactionRequest *TransactionRequest) FromJSON(rawMessage json.RawMessage) (helpers.TransactionRequest, error) {
	if err := json.Unmarshal(rawMessage, &transactionRequest); err != nil {
		return nil, err
	}

	return transactionRequest, nil
}
func (transactionRequest *TransactionRequest) GetBaseReq() rest.BaseReq {
	panic("Implement me")
}
func (transactionRequest *TransactionRequest) MakeMsg() (sdkTypes.Msg, error) {
	from, err := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if err != nil {
		return nil, err
	}

	fromID, err := baseIDs.ReadIdentityID(transactionRequest.FromId)
	if err != nil {
		return nil, err
	}

	toID, err := baseIDs.ReadIdentityID(transactionRequest.ToId)
	if err != nil {
		return nil, err
	}

	classificationID, err := baseIDs.ReadClassificationID(transactionRequest.ClassificationId)
	if err != nil {
		return nil, err
	}

	return newMessage(
		from,
		fromID,
		toID,
		classificationID,
	), nil
}
func (*TransactionRequest) RegisterLegacyAminoCodec(legacyAmino *codec.LegacyAmino) {
	codecUtilities.RegisterModuleConcrete(legacyAmino, &TransactionRequest{})
}
func requestPrototype() helpers.TransactionRequest {
	return &TransactionRequest{}
}

func newTransactionRequest(from string, fromID string, toID string, classificationID string) helpers.TransactionRequest {
	return &TransactionRequest{
		From:             from,
		FromId:           fromID,
		ToId:             toID,
		ClassificationId: classificationID,
	}
}
