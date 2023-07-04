package genesis

import (
	"context"

	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	"github.com/AssetMantle/schema/go/lists"
	"github.com/AssetMantle/schema/go/lists/base"
	sdkCodec "github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/x/maintainers/mappable"
	"github.com/AssetMantle/modules/x/maintainers/parameters"
)

var _ helpers.Genesis = (*Genesis)(nil)

func (genesis *Genesis) Default() helpers.Genesis {
	return Prototype()
}
func (genesis *Genesis) ValidateBasic(parameterManager helpers.ParameterManager) error {
	if len(genesis.ParameterList.Get()) != len(genesis.Default().(*Genesis).ParameterList.Get()) {
		return errorConstants.IncorrectFormat.Wrapf("expected %d parameters, got %d", len(genesis.Default().(*Genesis).ParameterList.Get()), len(genesis.ParameterList.Get()))
	}

	for _, parameter := range genesis.ParameterList.Get() {
		var isPresent bool
		for _, defaultParameter := range genesis.Default().(*Genesis).ParameterList.Get() {
			isPresent = false
			if defaultParameter.GetMetaProperty().Compare(parameter.GetMetaProperty()) == 0 {
				isPresent = true
				break
			}
		}

		if !isPresent {
			return errorConstants.IncorrectFormat.Wrapf("expected parameter %s not found", parameter.GetMetaProperty().GetKey().AsString())
		}

		if err := parameterManager.ValidateParameter(parameter); err != nil {
			return err
		}
	}

	for _, mappable := range genesis.Mappables {
		if err := mappable.ValidateBasic(); err != nil {
			return err
		}
	}

	return nil
}
func (genesis *Genesis) Import(context context.Context, mapper helpers.Mapper, parameterManager helpers.ParameterManager) {
	for _, Mappable := range genesis.Mappables {
		mapper.NewCollection(context).Add(Mappable)
	}

	parameterManager.Set(context, genesis.ParameterList)
}
func (genesis *Genesis) Export(context context.Context, mapper helpers.Mapper, parameterManager helpers.ParameterManager) helpers.Genesis {
	return genesis.Initialize(mapper.NewCollection(context).FetchAll().GetMappables(), parameterManager.Fetch(context).Get())
}
func (genesis *Genesis) Encode(jsonCodec sdkCodec.JSONCodec) []byte {
	bytes, err := jsonCodec.MarshalJSON(genesis)
	if err != nil {
		panic(err)
	}

	return bytes
}
func (genesis *Genesis) Decode(jsonCodec sdkCodec.JSONCodec, byte []byte) helpers.Genesis {
	if err := jsonCodec.UnmarshalJSON(byte, genesis); err != nil {
		panic(err)
	}

	return genesis
}
func (genesis *Genesis) Initialize(mappables []helpers.Mappable, parameterList lists.ParameterList) helpers.Genesis {
	if len(mappables) == 0 {
		genesis.Mappables = genesis.Default().(*Genesis).Mappables
	} else {
		genesis.Mappables = mappable.MappablesFromInterface(mappables)
	}

	if len(parameterList.Get()) == 0 {
		genesis.ParameterList = genesis.Default().(*Genesis).ParameterList
	} else {
		parameters := parameterList.Get()
		for _, defaultParameter := range genesis.Default().(*Genesis).ParameterList.Get() {
			for i, parameter := range parameters {
				if defaultParameter.GetMetaProperty().GetID().Compare(parameter.GetMetaProperty().GetID()) == 0 {
					parameters[i] = defaultParameter.Mutate(parameter.GetMetaProperty().GetData())
				}
			}
		}
		genesis.ParameterList = base.NewParameterList(parameters...).(*base.ParameterList)
	}

	return genesis
}

func Prototype() helpers.Genesis {
	return &Genesis{
		Mappables:     []*mappable.Mappable{},
		ParameterList: parameters.Prototype().Get().(*base.ParameterList),
	}
}
