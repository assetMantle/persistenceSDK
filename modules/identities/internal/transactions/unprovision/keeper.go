// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package unprovision

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/modules/identities/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
)

type transactionKeeper struct {
	mapper helpers.Mapper
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)
	identityID := message.IdentityID
	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(identityID))

	identity := identities.Get(key.FromID(identityID)).(mappables.Identity)
	if identity == nil {
		return newTransactionResponse(errors.EntityNotFound)
	}

	if !identity.IsProvisioned(message.From) {
		return newTransactionResponse(errors.NotAuthorized)
	}

	if !identity.IsProvisioned(message.To) {
		return newTransactionResponse(errors.EntityNotFound)
	}

	identities.Mutate(identity.UnprovisionAddress(message.To))

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, _ []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper
	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
