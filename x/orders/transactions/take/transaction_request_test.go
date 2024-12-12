// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package take

import (
	"fmt"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/data/base"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/spf13/viper"
)

var (
	fromAddress       = "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, _ = types.AccAddressFromBech32(fromAddress)
	immutables        = baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables          = baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))

	testClassificationID     = baseIDs.NewClassificationID(immutables, mutables)
	testFromID               = baseIDs.NewIdentityID(testClassificationID, immutables).(*baseIDs.IdentityID)
	testOrderID              = baseIDs.NewOrderID(testClassificationID, immutables).(*baseIDs.OrderID)
	commonTransactionRequest = helpers.PrototypeCommonTransactionRequest()
	takerSplit               = types.NewInt(60)
)

func Test_newTransactionRequest(t *testing.T) {
	type args struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		fromID                   string
		takerSplit               string
		orderID                  string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve", args{commonTransactionRequest, testFromID.AsString(), takerSplit.String(), testOrderID.AsString()}, transactionRequest{commonTransactionRequest, testFromID.AsString(), takerSplit.String(), testOrderID.AsString()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.commonTransactionRequest, tt.args.fromID, tt.args.takerSplit, tt.args.orderID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTransactionRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_requestPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionRequest
	}{
		{"+ve", transactionRequest{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := requestPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("requestPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_FromCLI(t *testing.T) {
	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromIdentityID, constants.TakerSplit, constants.OrderID})

	viper.Set(constants.FromIdentityID.GetName(), testFromID.AsString())
	viper.Set(constants.TakerSplit.GetName(), takerSplit.String())
	viper.Set(constants.OrderID.GetName(), testOrderID.AsString())
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		TakerSplit               string
		OrderID                  string
	}
	type args struct {
		cliCommand helpers.CLICommand
		context    client.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionRequest
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), takerSplit.String(), testOrderID.AsString()}, args{cliCommand, client.Context{}.WithCodec(baseHelpers.CodecPrototype())}, transactionRequest{commonTransactionRequest, testFromID.AsString(), takerSplit.String(), testOrderID.AsString()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				TakerSplit:               tt.fields.TakerSplit,
				OrderID:                  tt.fields.OrderID,
			}
			got, err := transactionRequest.FromCLI(tt.args.cliCommand, tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromCLI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_GetBaseReq(t *testing.T) {
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		TakerSplit               string
		OrderID                  string
	}
	tests := []struct {
		name   string
		fields fields
		want   helpers.CommonTransactionRequest
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), takerSplit.String(), testOrderID.AsString()}, commonTransactionRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				TakerSplit:               tt.fields.TakerSplit,
				OrderID:                  tt.fields.OrderID,
			}
			if got := transactionRequest.GetCommonTransactionRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommonTransactionRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		TakerSplit               string
		OrderID                  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    types.Msg
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), takerSplit.String(), testOrderID.AsString()}, NewMessage(fromAccAddress, testFromID, takerSplit, testOrderID), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				TakerSplit:               tt.fields.TakerSplit,
				OrderID:                  tt.fields.OrderID,
			}
			got, err := transactionRequest.MakeMsg()
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeMsg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeMsg() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	type fields struct {
		commonTransactionRequest helpers.CommonTransactionRequest
		FromID                   string
		TakerSplit               string
		OrderID                  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{commonTransactionRequest, testFromID.AsString(), takerSplit.String(), testOrderID.AsString()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				CommonTransactionRequest: tt.fields.commonTransactionRequest,
				FromID:                   tt.fields.FromID,
				TakerSplit:               tt.fields.TakerSplit,
				OrderID:                  tt.fields.OrderID,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
