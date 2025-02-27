// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package maintainer

import (
	"net/http"

	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/cosmos/cosmos-sdk/client"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/maintainers/key"
)

var _ helpers.QueryRequest = (*QueryRequest)(nil)

// Validate godoc
// @Summary Search for a maintainer by maintainer ID
// @Description Able to query the maintainers details
// @Accept json
// @Produce json
// @Tags Maintainers
// @Param maintainerID path string true "Unique identifier of a maintainer."
// @Success 200 {object} queryResponse "Message for a successful query response"
// @Failure default  {object}  queryResponse "Message for an unexpected error response."
// @Router /maintainers/maintainers/{maintainerID} [get]
func (queryRequest *QueryRequest) Validate() error {
	if err := queryRequest.Key.ValidateBasic(); err != nil {
		return constants.InvalidRequest.Wrapf(err.Error())
	}

	return nil
}

func (*QueryRequest) FromCLI(cliCommand helpers.CLICommand, _ client.Context) (helpers.QueryRequest, error) {
	if maintainerID, err := baseIDs.PrototypeMaintainerID().FromString(cliCommand.ReadString(constants.MaintainerID)); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(maintainerID.(ids.MaintainerID)), nil
	}
}

func (*QueryRequest) FromHTTPRequest(httpRequest *http.Request) (helpers.QueryRequest, error) {
	if maintainerID, err := baseIDs.PrototypeMaintainerID().FromString(httpRequest.URL.Query().Get(constants.Key.GetName())); err != nil {
		return &QueryRequest{}, err
	} else {
		return newQueryRequest(maintainerID.(ids.MaintainerID)), nil
	}
}

func requestPrototype() helpers.QueryRequest {
	return &QueryRequest{}
}

func newQueryRequest(maintainerID ids.MaintainerID) helpers.QueryRequest {
	return &QueryRequest{Key: key.NewKey(maintainerID).(*key.Key)}
}
