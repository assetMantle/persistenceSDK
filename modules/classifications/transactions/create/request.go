package create

import (
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"strings"
)

type transactionRequest struct {
	BaseReq       rest.BaseReq `json:"baseReq"`
	MaintainersID string       `json:"maintainersID" valid:"required~required field maintainersID missing matches(^[A-Za-z]$)~invalid field maintainersID"`
	Traits        string       `json:"traits" valid:"required~required field traits missing matches(^[A-Za-z]$)~invalid field traits"`
}

var _ helpers.TransactionRequest = (*transactionRequest)(nil)

func (transactionRequest transactionRequest) FromCLI(cliCommand helpers.CLICommand, cliContext context.CLIContext) helpers.TransactionRequest {
	return newTransactionRequest(
		cliCommand.ReadBaseReq(cliContext),
		cliCommand.ReadString(constants.MaintainersID),
		cliCommand.ReadString(constants.Traits),
	)
}

func (transactionRequest transactionRequest) GetBaseReq() rest.BaseReq {
	return transactionRequest.BaseReq
}

func (transactionRequest transactionRequest) MakeMsg() sdkTypes.Msg {
	from, Error := sdkTypes.AccAddressFromBech32(transactionRequest.GetBaseReq().From)
	if Error != nil {
		panic(errors.New(fmt.Sprintf("")))
	}

	traits := strings.Split(transactionRequest.Traits, constants.TraitsSeparator)
	if len(traits) > constants.MaxTraitCount {
		//TODO handle
		panic(errors.New(fmt.Sprintf("")))
	}

	var traitList []types.Trait
	for _, trait := range traits {
		traitIDAndProperty := strings.Split(trait, constants.TraitIDAndPropertySeparator)
		if len(traitIDAndProperty) == 2 && traitIDAndProperty[0] != "" {
			traitID := base.NewID(traitIDAndProperty[0])
			traitList = append(traitList, base.NewTrait(traitID, base.NewProperty(traitID, base.NewFact(traitIDAndProperty[1], base.NewSignatures(nil)))))
		}
	}

	return newMessage(
		from,
		base.NewID(transactionRequest.MaintainersID),
		base.NewTraits(traitList),
	)
}

func requestPrototype() helpers.TransactionRequest {
	return transactionRequest{}
}

func newTransactionRequest(baseReq rest.BaseReq, maintainersID string, traits string) helpers.TransactionRequest {
	return transactionRequest{
		BaseReq:       baseReq,
		MaintainersID: maintainersID,
		Traits:        traits,
	}
}
