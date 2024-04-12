// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package authenticate

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	"github.com/AssetMantle/modules/utilities/random"
	"github.com/AssetMantle/modules/x/identities/constants"
	"github.com/AssetMantle/modules/x/identities/mapper"
	"github.com/AssetMantle/modules/x/identities/parameters"
	"github.com/AssetMantle/modules/x/identities/record"
	baseData "github.com/AssetMantle/schema/go/data/base"
	"github.com/AssetMantle/schema/go/documents"
	"github.com/AssetMantle/schema/go/documents/base"
	"github.com/AssetMantle/schema/go/errors"
	errorConstants "github.com/AssetMantle/schema/go/errors/constants"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/mock"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"
	"reflect"
	"testing"
)

type MockAuxiliary struct {
	mock.Mock
}

var _ helpers.Auxiliary = (*MockAuxiliary)(nil)

func (mockAuxiliary *MockAuxiliary) GetName() string { panic(mockAuxiliary) }
func (mockAuxiliary *MockAuxiliary) GetKeeper() helpers.AuxiliaryKeeper {
	args := mockAuxiliary.Called()
	return args.Get(0).(helpers.AuxiliaryKeeper)
}
func (mockAuxiliary *MockAuxiliary) Initialize(_ helpers.Mapper, _ helpers.ParameterManager, _ ...interface{}) helpers.Auxiliary {
	panic(mockAuxiliary)
}

type MockAuxiliaryKeeper struct {
	mock.Mock
}

var _ helpers.AuxiliaryKeeper = (*MockAuxiliaryKeeper)(nil)

func (mockAuxiliaryKeeper *MockAuxiliaryKeeper) Help(context context.Context, request helpers.AuxiliaryRequest) (helpers.AuxiliaryResponse, error) {
	args := mockAuxiliaryKeeper.Called(context, request)
	return args.Get(0).(helpers.AuxiliaryResponse), args.Error(1)
}
func (mockAuxiliaryKeeper *MockAuxiliaryKeeper) Initialize(m2 helpers.Mapper, manager helpers.ParameterManager, i []interface{}) helpers.Keeper {
	args := mockAuxiliaryKeeper.Called(m2, manager, i)
	return args.Get(0).(helpers.Keeper)
}

const (
	ChainID = "testChain"
)

var (
	moduleStoreKey = sdkTypes.NewKVStoreKey(constants.ModuleName)

	supplementAuxiliaryKeeper = new(MockAuxiliaryKeeper)
	_                         = supplementAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(newAuxiliaryResponse(), nil)

	supplementAuxiliary = new(MockAuxiliary)
	_                   = supplementAuxiliary.On("GetKeeper").Return(supplementAuxiliaryKeeper)

	provisionedAddress   = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	unprovisionedAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	testIdentity         = base.PrototypeNameIdentity().ProvisionAddress(provisionedAddress)
	testIdentityID       = testIdentity.(documents.NameIdentity).GetNameIdentityID()

	paramsStoreKey           = sdkTypes.NewKVStoreKey(paramsTypes.StoreKey)
	paramsTransientStoreKeys = sdkTypes.NewTransientStoreKey(paramsTypes.TStoreKey)
	ParamsKeeper             = paramsKeeper.NewKeeper(encodingConfig.Marshaler, encodingConfig.Amino, paramsStoreKey, paramsTransientStoreKeys)

	encodingConfig = simapp.MakeTestEncodingConfig()

	parameterManager = parameters.Prototype().Initialize(ParamsKeeper.Subspace(constants.ModuleName).WithKeyTable(parameters.Prototype().GetKeyTable()))

	AuxiliaryKeeper = auxiliaryKeeper{mapper.Prototype().Initialize(moduleStoreKey), parameterManager, supplementAuxiliary}

	setContext = func() sdkTypes.Context {
		memDB := tendermintDB.NewMemDB()
		commitMultiStore := store.NewCommitMultiStore(memDB)
		commitMultiStore.MountStoreWithDB(moduleStoreKey, sdkTypes.StoreTypeIAVL, memDB)
		commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
		commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
		_ = commitMultiStore.LoadLatestVersion()
		return sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{ChainID: ChainID}, false, log.NewNopLogger())

	}

	Context = setContext()

	_ = AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).
		Add(record.NewRecord(testIdentity))
)

func Test_auxiliaryKeeper_Help(t *testing.T) {
	tests := []struct {
		name    string
		setup   func()
		request helpers.AuxiliaryRequest
		want    helpers.AuxiliaryResponse
		wantErr errors.Error
	}{
		{
			"valid request",
			func() {},
			auxiliaryRequest{
				Address:    provisionedAddress,
				IdentityID: testIdentityID,
			},
			newAuxiliaryResponse(),
			nil,
		},
		{
			"invalid request",
			func() {},
			auxiliaryRequest{
				Address:    unprovisionedAddress,
				IdentityID: testIdentityID,
			},
			nil,
			errorConstants.NotAuthorized,
		},
		{
			"identity not found",
			func() {},
			auxiliaryRequest{
				Address:    provisionedAddress,
				IdentityID: base.NewNameIdentity(baseIDs.NewStringID("not found"), baseData.NewListData()).GetNameIdentityID(),
			},
			nil,
			errorConstants.EntityNotFound,
		},
		{
			"many identities present",
			func() {
				for i := 0; i < 10000; i++ {
					_ = AuxiliaryKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).
						Add(record.NewRecord(base.NewNameIdentity(baseIDs.NewStringID(random.GenerateUniqueIdentifier()), baseData.NewListData(baseData.NewAccAddressData(sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())), baseData.NewAccAddressData(sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())), baseData.NewAccAddressData(sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())), baseData.NewAccAddressData(sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address()))))))
				}
			},
			auxiliaryRequest{
				Address:    provisionedAddress,
				IdentityID: testIdentityID,
			},
			newAuxiliaryResponse(),
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			got, err := AuxiliaryKeeper.Help(sdkTypes.WrapSDKContext(Context), tt.request)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Help() got = %v, want %v", got, tt.want)
			}

			if err != nil && tt.wantErr == nil || err == nil && tt.wantErr != nil || err != nil && tt.wantErr != nil && !tt.wantErr.Is(err) {
				t.Errorf("Help() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
