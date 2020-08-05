package mapper

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/traits"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type Identity struct {
	ID                       types.ID              `json:"id" valid:"required~required field id missing"`
	ProvisionedAddressList   []sdkTypes.AccAddress `json:"provisionedAddressList" valid:"required~required field provisionedAddressList missing"`
	UnprovisionedAddressList []sdkTypes.AccAddress `json:"unprovisionedAddressList" valid:"required~required field unprovisionedAddressList missing"`
	Immutables               types.Immutables      `json:"immutables" valid:"required~required field immutables missing"`
	Mutables                 types.Mutables        `json:"mutables" valid:"required~required field mutables missing"`
}

var _ mappables.InterIdentity = (*Identity)(nil)

func (identity Identity) GetID() types.ID { return identity.ID }
func (identity Identity) GetChainID() types.ID {
	return identityIDFromInterface(identity.ID).ChainID
}

func (identity Identity) GetClassificationID() types.ID {
	return identityIDFromInterface(identity.ID).ClassificationID
}
func (identity Identity) GetProvisionedAddressList() []sdkTypes.AccAddress {
	return identity.ProvisionedAddressList
}
func (identity Identity) GetUnprovisionedAddressList() []sdkTypes.AccAddress {
	return identity.UnprovisionedAddressList
}
func (identity Identity) ProvisionAddress(accAddress sdkTypes.AccAddress) mappables.InterIdentity {
	identity.ProvisionedAddressList = append(identity.ProvisionedAddressList, accAddress)
	return identity
}
func (identity Identity) UnprovisionAddress(accAddress sdkTypes.AccAddress) mappables.InterIdentity {
	for i, provisionedAddress := range identity.ProvisionedAddressList {
		if provisionedAddress.Equals(accAddress) {
			identity.ProvisionedAddressList = append(identity.ProvisionedAddressList[:i], identity.ProvisionedAddressList[i+1:]...)
			identity.UnprovisionedAddressList = append(identity.UnprovisionedAddressList, accAddress)
			return identity
		}
	}
	return identity
}
func (identity Identity) GetImmutables() types.Immutables { return identity.Immutables }
func (identity Identity) GetMutables() types.Mutables     { return identity.Mutables }
func (identity Identity) IsProvisioned(accAddress sdkTypes.AccAddress) bool {
	for _, provisionedAddress := range identity.ProvisionedAddressList {
		if provisionedAddress.Equals(accAddress) {
			return true
		}
	}
	return false
}
func (identity Identity) IsUnprovisioned(accAddress sdkTypes.AccAddress) bool {
	for _, unprovisionedAddress := range identity.UnprovisionedAddressList {
		if unprovisionedAddress.Equals(accAddress) {
			return true
		}
	}
	return false
}
func (identity Identity) Encode() []byte {
	return packageCodec.MustMarshalBinaryBare(identity)
}
func (identity Identity) Decode(bytes []byte) traits.Mappable {
	packageCodec.MustUnmarshalBinaryBare(bytes, &identity)
	return identity
}
func identityPrototype() traits.Mappable {
	return Identity{}
}
func NewIdentity(identityID types.ID, provisionedAddressList []sdkTypes.AccAddress, unprovisionedAddressList []sdkTypes.AccAddress, immutables types.Immutables, mutables types.Mutables) mappables.InterIdentity {
	return Identity{
		ID:                       identityID,
		ProvisionedAddressList:   provisionedAddressList,
		UnprovisionedAddressList: unprovisionedAddressList,
		Immutables:               immutables,
		Mutables:                 mutables,
	}
}
