// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transaction

import (
	schema "github.com/AssetMantle/schema/x"
	"github.com/AssetMantle/schema/x/helpers"
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterLegacyAminoCodec(messagePrototype func() helpers.Message) *codec.LegacyAmino {
	Codec := codec.NewLegacyAmino()
	messagePrototype().RegisterLegacyAminoCodec(Codec)
	schema.RegisterLegacyAminoCodec(Codec)
	Codec.Seal()

	return Codec
}
