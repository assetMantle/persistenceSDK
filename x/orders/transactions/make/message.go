// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package make

import (
	"cosmossdk.io/math"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
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
	if err := message.ClassificationID.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.TakerID.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.MakerAssetID.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.TakerAssetID.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.ExpiresIn.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.ImmutableMetaProperties.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.MutableMetaProperties.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.ImmutableProperties.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if err := message.MutableProperties.ValidateBasic(); err != nil {
		return constants.InvalidMessage.Wrapf(err.Error())
	}
	if _, ok := sdkTypes.NewIntFromString(message.MakerSplit); !ok {
		return constants.InvalidMessage.Wrapf("maker split %s is not a valid integer", message.MakerSplit)
	} else if _, ok := sdkTypes.NewIntFromString(message.TakerSplit); !ok {
		return constants.InvalidMessage.Wrapf("taker split %s is not a valid integer", message.TakerSplit)
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
func NewMessage(from sdkTypes.AccAddress, fromID ids.IdentityID, classificationID ids.ClassificationID, takerID ids.IdentityID, makerAssetID ids.AssetID, takerAssetID ids.AssetID, expiresIn typesSchema.Height, makerSplit math.Int, takerSplit math.Int, immutableMetaProperties lists.PropertyList, immutableProperties lists.PropertyList, mutableMetaProperties lists.PropertyList, mutableProperties lists.PropertyList) sdkTypes.Msg {

	return &Message{
		From:                    from.String(),
		FromID:                  fromID.(*baseIDs.IdentityID),
		ClassificationID:        classificationID.(*baseIDs.ClassificationID),
		TakerID:                 takerID.(*baseIDs.IdentityID),
		MakerAssetID:            makerAssetID.(*baseIDs.AssetID),
		TakerAssetID:            takerAssetID.(*baseIDs.AssetID),
		ExpiresIn:               expiresIn.(*baseTypes.Height),
		MakerSplit:              makerSplit.String(),
		TakerSplit:              takerSplit.String(),
		ImmutableMetaProperties: immutableMetaProperties.(*baseLists.PropertyList),
		ImmutableProperties:     immutableProperties.(*baseLists.PropertyList),
		MutableMetaProperties:   mutableMetaProperties.(*baseLists.PropertyList),
		MutableProperties:       mutableProperties.(*baseLists.PropertyList),
	}
}
