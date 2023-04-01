package docs

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/AssetMantle/modules/schema/data"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/properties"
	"github.com/AssetMantle/modules/schema/properties/constants"
)

func assetClassificationHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		id, immutables, _ := ReadAndProcess(context, false, true, responseWriter, httpRequest)
		rest.PostProcessResponse(responseWriter, context, newResponse(id.AsString(), immutables.GetProperty(constants.BondAmountProperty.GetID()).Get().(properties.MetaProperty).GetData().Get().(data.NumberData).AsString(), nil))
	}
}
func assetIDHandler(context client.Context) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, httpRequest *http.Request) {
		classificationID, immutables, _ := ReadAndProcess(context, false, false, responseWriter, httpRequest)

		rest.PostProcessResponse(responseWriter, context, newResponse(baseIDs.NewAssetID(classificationID, immutables).AsString(), "", nil))
	}
}
