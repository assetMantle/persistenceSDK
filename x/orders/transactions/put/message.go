// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package put

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	typesSchema "github.com/AssetMantle/schema/types"
	baseTypes "github.com/AssetMantle/schema/types/base"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var _ helpers.Message = (*Message)(nil)

func (message *Message) GetFromAddress() sdkTypes.AccAddress {
	from, err := sdkTypes.AccAddressFromBech32(message.From)
	if err != nil || from.Empty() {
		// NOTE: This should never happen as the message is validated before it is sent
		return nil
	}
	return from
}
func (message *Message) GetFromIdentityID() ids.IdentityID {
	return message.FromID
}
func (message *Message) ValidateBasic() error {
	if message.GetFromAddress() == nil {
		return constants.InvalidMessage.Wrapf("from address %s is not a valid address", message.From)
	}
	if err := message.GetFromIdentityID().ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.MakerAssetID.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.TakerAssetID.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if _, ok := sdkTypes.NewIntFromString(message.MakerSplit); !ok {
		return constants.InvalidMessage.Wrapf("maker split %s is not a valid integer", message.MakerSplit)
	}
	if _, ok := sdkTypes.NewIntFromString(message.TakerSplit); !ok {
		return constants.InvalidMessage.Wrapf("taker split %s is not a valid integer", message.TakerSplit)
	}
	if err := message.ExpiryHeight.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	return nil
}
func (message *Message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.GetFromAddress()}
}
func (message *Message) RegisterInterface(interfaceRegistry types.InterfaceRegistry) {
	interfaceRegistry.RegisterImplementations((*sdkTypes.Msg)(nil), message)
}

func messagePrototype() helpers.Message {
	return &Message{}
}

func NewMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, makerAssetID ids.AssetID, takerAssetID ids.AssetID, makerSplit math.Int, takerSplit math.Int, expiryHeight typesSchema.Height) sdkTypes.Msg {
	return &Message{
		From:         from.String(),
		FromID:       fromID.(*baseIDs.IdentityID),
		MakerAssetID: makerAssetID.(*baseIDs.AssetID),
		TakerAssetID: takerAssetID.(*baseIDs.AssetID),
		MakerSplit:   makerSplit.String(),
		TakerSplit:   takerSplit.String(),
		ExpiryHeight: expiryHeight.(*baseTypes.Height),
	}
}
