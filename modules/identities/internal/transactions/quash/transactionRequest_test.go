// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package quash

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/AssetMantle/schema/schema"
	baseData "github.com/AssetMantle/schema/x/data/base"
	"github.com/AssetMantle/schema/x/helpers"
	baseHelpers "github.com/AssetMantle/schema/x/helpers/base"
	"github.com/AssetMantle/schema/x/helpers/constants"
	baseIDs "github.com/AssetMantle/schema/x/ids/base"
	baseLists "github.com/AssetMantle/schema/x/lists/base"
	baseProperties "github.com/AssetMantle/schema/x/properties/base"
	baseQualified "github.com/AssetMantle/schema/x/qualified/base"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/utilities/transaction"
)

func createTestInput(t *testing.T) (rest.BaseReq, string, *baseIDs.IdentityID) {
	immutables := baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	testBaseReq := rest.BaseReq{From: fromAddress, ChainID: "test", Fees: types.NewCoins()}
	testToAddress := "cosmos1vx8knpllrj7n963p9ttd80w47kpacrhuts497x"
	return testBaseReq, testToAddress, testFromID.(*baseIDs.IdentityID)
}

func Test_newTransactionRequest(t *testing.T) {
	testBaseReq, testToAddress, testFromID := createTestInput(t)
	type args struct {
		baseReq    rest.BaseReq
		fromID     string
		identityID string
	}
	tests := []struct {
		name string
		args args
		want helpers.TransactionRequest
	}{
		{"+ve with nil", args{}, transactionRequest{}},
		{"+ve", args{testBaseReq, testToAddress, testFromID.AsString()}, transactionRequest{testBaseReq, testToAddress, testFromID.AsString()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTransactionRequest(tt.args.baseReq, tt.args.fromID, tt.args.identityID); !reflect.DeepEqual(got, tt.want) {
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
	var legacyAmino = codec.NewLegacyAmino()
	schema.RegisterLegacyAminoCodec(legacyAmino)
	std.RegisterLegacyAminoCodec(legacyAmino)
	legacyAmino.Seal()

	cliCommand := baseHelpers.NewCLICommand("", "", "", []helpers.CLIFlag{constants.FromID, constants.IdentityID})

	testBaseReq, _, testFromID := createTestInput(t)
	type fields struct {
		BaseReq    rest.BaseReq
		FromID     string
		IdentityID string
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
		{"+ve", fields{testBaseReq, testFromID.AsString(), testFromID.AsString()}, args{cliCommand: cliCommand, context: baseHelpers.TestClientContext}, transactionRequest{cliCommand.ReadBaseReq(baseHelpers.TestClientContext), cliCommand.ReadString(constants.FromID), cliCommand.ReadString(constants.IdentityID)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			got, err := transactionRequest.FromCLI(tt.args.cliCommand, tt.args.context)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromCLI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromCLI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_FromJSON(t *testing.T) {
	testBaseReq, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		BaseReq    rest.BaseReq
		FromID     string
		IdentityID string
	}
	type args struct {
		rawMessage json.RawMessage
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionRequest
		wantErr bool
	}{
		{"+ve", fields{testBaseReq, testToAddress, testFromID.AsString()}, args{types.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(&Message{types.AccAddress("cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c").String(), testFromID, testFromID}))}, transactionRequest{testBaseReq, testToAddress, testFromID.AsString()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			got, err := transactionRequest.FromJSON(tt.args.rawMessage)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_GetBaseReq(t *testing.T) {
	testBaseReq, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		BaseReq    rest.BaseReq
		FromID     string
		IdentityID string
	}
	tests := []struct {
		name   string
		fields fields
		want   rest.BaseReq
	}{
		{"+ve", fields{testBaseReq, testToAddress, testFromID.AsString()}, testBaseReq},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			if got := transactionRequest.GetBaseReq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBaseReq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionRequest_MakeMsg(t *testing.T) {
	testBaseReq, _, testFromID := createTestInput(t)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := types.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	type fields struct {
		BaseReq    rest.BaseReq
		FromID     string
		IdentityID string
	}
	tests := []struct {
		name    string
		fields  fields
		want    helpers.Message
		wantErr bool
	}{
		{"+ve", fields{testBaseReq, testFromID.AsString(), testFromID.AsString()}, &Message{fromAccAddress.String(), testFromID, testFromID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
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

func Test_transactionRequest_RegisterCodec(t *testing.T) {
	testBaseReq, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		BaseReq    rest.BaseReq
		FromID     string
		IdentityID string
	}
	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve with nil", fields{}, args{codec.NewLegacyAmino()}},
		{"+ve", fields{testBaseReq, testToAddress, testFromID.AsString()}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			tr.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_transactionRequest_Validate(t *testing.T) {
	testBaseReq, testToAddress, testFromID := createTestInput(t)
	type fields struct {
		BaseReq    rest.BaseReq
		FromID     string
		IdentityID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve with nil", fields{}, true},
		{"+ve", fields{testBaseReq, testToAddress, testFromID.AsString()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionRequest := transactionRequest{
				BaseReq:    tt.fields.BaseReq,
				FromID:     tt.fields.FromID,
				IdentityID: tt.fields.IdentityID,
			}
			if err := transactionRequest.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
