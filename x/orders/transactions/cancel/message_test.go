// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package cancel

import (
	"reflect"
	"testing"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/utilities/transaction"
	"github.com/AssetMantle/modules/x/orders/module"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/ids"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type fields struct {
	From    string
	FromID  *baseIDs.IdentityID
	OrderID *baseIDs.OrderID
}

func CreateTestInputForMessages(t *testing.T) (*baseIDs.OrderID, *baseIDs.IdentityID, sdkTypes.AccAddress, sdkTypes.Msg) {
	// testOrderID := baseIDs.NewStringID("orderID")
	immutables := baseQualified.NewImmutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID2"), baseData.NewStringData("Data2"))))
	mutables := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("Data1"))))
	// mutables2 := baseQualified.NewMutables(base.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID"), baseData.NewStringData(""))))

	testClassificationID := baseIDs.NewClassificationID(immutables, mutables)
	testFromID := baseIDs.NewIdentityID(testClassificationID, immutables)
	testOrderID := baseIDs.NewOrderID(testClassificationID, immutables)

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)

	testMessage := newMessage(fromAccAddress, testFromID, testOrderID)

	return testOrderID.(*baseIDs.OrderID), testFromID.(*baseIDs.IdentityID), fromAccAddress, testMessage
}

func Test_messageFromInterface(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, testMessage := CreateTestInputForMessages(t)
	type args struct {
		msg sdkTypes.Msg
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		{"+ve", args{testMessage}, &Message{fromAccAddress.String(), testFromID, testOrderID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := messageFromInterface(tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messageFromInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_messagePrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.Message
	}{
		{"+ve", &Message{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := messagePrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("messagePrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSignBytes(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, testMessage := CreateTestInputForMessages(t)

	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID}, sdkTypes.MustSortJSON(transaction.RegisterLegacyAminoCodec(messagePrototype).MustMarshalJSON(testMessage))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			if got := message.GetSignBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSignBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_GetSigners(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, _ := CreateTestInputForMessages(t)

	tests := []struct {
		name   string
		fields fields
		want   []sdkTypes.AccAddress
	}{
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID}, []sdkTypes.AccAddress{fromAccAddress}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			if got := message.GetSigners(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSigners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_RegisterCodec(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, _ := CreateTestInputForMessages(t)

	type args struct {
		legacyAmino *codec.LegacyAmino
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID}, args{codec.NewLegacyAmino()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			me.RegisterLegacyAminoCodec(tt.args.legacyAmino)
		})
	}
}

func Test_message_Route(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, _ := CreateTestInputForMessages(t)

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID}, module.Name},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			if got := message.Route(); got != tt.want {
				t.Errorf("Route() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_Type(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, _ := CreateTestInputForMessages(t)

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID}, Transaction.GetName()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			if got := message.Type(); got != tt.want {
				t.Errorf("Type() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_ValidateBasic(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, _ := CreateTestInputForMessages(t)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"+ve", fields{fromAccAddress.String(), testFromID, testOrderID}, false},
		{"-ve", fields{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := &Message{
				From:    tt.fields.From,
				FromID:  tt.fields.FromID,
				OrderID: tt.fields.OrderID,
			}
			if err := message.ValidateBasic(); (err != nil) != tt.wantErr {
				t.Errorf("ValidateBasic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newMessage(t *testing.T) {
	testOrderID, testFromID, fromAccAddress, _ := CreateTestInputForMessages(t)

	type args struct {
		from    sdkTypes.AccAddress
		fromID  ids.IdentityID
		orderID ids.OrderID
	}
	tests := []struct {
		name string
		args args
		want sdkTypes.Msg
	}{
		{"+ve", args{fromAccAddress, testFromID, testOrderID}, &Message{fromAccAddress.String(), testFromID, testOrderID}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMessage(tt.args.from, tt.args.fromID, tt.args.orderID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
