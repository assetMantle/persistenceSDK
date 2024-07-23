// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package mint

import (
	"context"
	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	errorConstants "github.com/AssetMantle/modules/helpers/constants"
	"github.com/AssetMantle/modules/x/assets/constants"
	"github.com/AssetMantle/modules/x/assets/key"
	"github.com/AssetMantle/modules/x/assets/mapper"
	"github.com/AssetMantle/modules/x/assets/parameters"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/bond"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
	"github.com/AssetMantle/modules/x/orders/record"
	"github.com/AssetMantle/modules/x/splits/auxiliaries/mint"
	baseData "github.com/AssetMantle/schema/data/base"
	baseDocuments "github.com/AssetMantle/schema/documents/base"
	"github.com/AssetMantle/schema/ids"
	baseIDs "github.com/AssetMantle/schema/ids/base"
	"github.com/AssetMantle/schema/lists"
	baseLists "github.com/AssetMantle/schema/lists/base"
	"github.com/AssetMantle/schema/parameters/base"
	baseProperties "github.com/AssetMantle/schema/properties/base"
	constantProperties "github.com/AssetMantle/schema/properties/constants"
	baseQualified "github.com/AssetMantle/schema/qualified/base"
	tendermintDB "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/cometbft/cometbft/libs/log"
	protoTendermintTypes "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/store"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

// start
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
	TestMinterModuleName = "testMinter"
	Denom                = "stake"
	ChainID              = "testChain"
	GenesisSupply        = 1000000000000
)

var (
	moduleStoreKey = sdkTypes.NewKVStoreKey(constants.ModuleName)

	authenticateAuxiliaryKeeper               = new(MockAuxiliaryKeeper)
	authenticateAuxiliaryFailureAddress       = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_                                         = authenticateAuxiliaryKeeper.On("Help", mock.Anything, authenticate.NewAuxiliaryRequest(authenticateAuxiliaryFailureAddress, baseIDs.PrototypeIdentityID())).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                                         = authenticateAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authenticateAuxiliary                     = new(MockAuxiliary)
	_                                         = authenticateAuxiliary.On("GetKeeper").Return(authenticateAuxiliaryKeeper)
	authorizeAuxiliaryKeeper                  = new(MockAuxiliaryKeeper)
	authorizeAuxiliaryFailureClassificationID = baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList()))
	_                                         = authorizeAuxiliaryKeeper.On("Help", mock.Anything, authorize.NewAuxiliaryRequest(authorizeAuxiliaryFailureClassificationID, baseIDs.PrototypeIdentityID(), baseIDs.NewStringID(""))).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                                         = authorizeAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	authorizeAuxiliary                        = new(MockAuxiliary)
	_                                         = authorizeAuxiliary.On("GetKeeper").Return(authorizeAuxiliaryKeeper)
	mintAuxiliaryKeeper                       = new(MockAuxiliaryKeeper)
	mintAuxiliaryFailureID                    = baseIDs.NewIdentityID(baseIDs.NewClassificationID(baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.BondAmountProperty.GetID().GetKey(), baseData.NewNumberData(sdkTypes.NewInt(1)))))), baseQualified.NewImmutables(baseLists.NewPropertyList()))
	_                                         = mintAuxiliaryKeeper.On("Help", mock.Anything, mint.NewAuxiliaryRequest(mintAuxiliaryFailureID, baseDocuments.NewCoinAsset(Denom).GetCoinAssetID(), sdkTypes.OneInt())).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                                         = mintAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	mintAuxiliary                             = new(MockAuxiliary)
	_                                         = mintAuxiliary.On("GetKeeper").Return(mintAuxiliaryKeeper)
	bondAuxiliaryKeeper                       = new(MockAuxiliaryKeeper)
	bondAuxiliaryFailureAddress               = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_                                         = bondAuxiliaryKeeper.On("Help", mock.Anything, bond.NewAuxiliaryRequest(baseIDs.PrototypeClassificationID(), bondAuxiliaryFailureAddress, sdkTypes.OneInt())).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                                         = bondAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	_                                         = bondAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	bondAuxiliary                             = new(MockAuxiliary)
	_                                         = bondAuxiliary.On("GetKeeper").Return(bondAuxiliaryKeeper)
	conformAuxiliaryKeeper                    = new(MockAuxiliaryKeeper)
	conformAuxiliaryFailureID                 = baseIDs.PrototypeClassificationID()
	_                                         = conformAuxiliaryKeeper.On("Help", mock.Anything, conform.NewAuxiliaryRequest(conformAuxiliaryFailureID, baseQualified.NewImmutables(baseLists.NewPropertyList()), baseQualified.NewMutables(baseLists.NewPropertyList()))).Return(new(helpers.AuxiliaryResponse), errorConstants.MockError)
	_                                         = conformAuxiliaryKeeper.On("Help", mock.Anything, mock.Anything).Return(new(helpers.AuxiliaryResponse), nil)
	conformAuxiliary                          = new(MockAuxiliary)
	_                                         = conformAuxiliary.On("GetKeeper").Return(conformAuxiliaryKeeper)

	codec                    = baseHelpers.TestCodec()
	paramsStoreKey           = sdkTypes.NewKVStoreKey(paramsTypes.StoreKey)
	paramsTransientStoreKeys = sdkTypes.NewTransientStoreKey(paramsTypes.TStoreKey)
	ParamsKeeper             = paramsKeeper.NewKeeper(
		codec,
		codec.GetLegacyAmino(),
		paramsStoreKey,
		paramsTransientStoreKeys,
	)

	authStoreKey             = sdkTypes.NewKVStoreKey(authTypes.StoreKey)
	moduleAccountPermissions = map[string][]string{TestMinterModuleName: {authTypes.Minter}, constants.ModuleName: nil}
	AuthKeeper               = authKeeper.NewAccountKeeper(codec, authStoreKey, authTypes.ProtoBaseAccount, moduleAccountPermissions, sdkTypes.GetConfig().GetBech32AccountAddrPrefix(), authTypes.NewModuleAddress(govTypes.ModuleName).String())

	bankStoreKey         = sdkTypes.NewKVStoreKey(bankTypes.StoreKey)
	blacklistedAddresses = map[string]bool{authTypes.NewModuleAddress(TestMinterModuleName).String(): false, authTypes.NewModuleAddress(constants.ModuleName).String(): false}
	BankKeeper           = bankKeeper.NewBaseKeeper(codec, bankStoreKey, AuthKeeper, blacklistedAddresses, authTypes.NewModuleAddress(govTypes.ModuleName).String())

	Context = setContext()

	coinSupply = sdkTypes.NewCoins(sdkTypes.NewCoin(Denom, sdkTypes.NewInt(GenesisSupply)))
	_          = BankKeeper.MintCoins(Context, TestMinterModuleName, coinSupply)

	genesisAddress = sdkTypes.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	_              = BankKeeper.SendCoinsFromModuleToAccount(Context, TestMinterModuleName, genesisAddress, coinSupply)
	//parameterManager = parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))
	parameterManager = parameters.Prototype().Initialize(ParamsKeeper.Subspace(constants.ModuleName).WithKeyTable(parameters.Prototype().GetKeyTable())).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.WrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(Denom)))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.BurnEnabledProperty.GetKey(), baseData.NewBooleanData(true))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.RenumerateEnabledProperty.GetKey(), baseData.NewBooleanData(true))))).
				Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.UnwrapAllowedCoinsProperty.GetKey(), baseData.NewListData(baseData.NewStringData(Denom))))))

	TransactionKeeper = transactionKeeper{
		mapper:                mapper.Prototype().Initialize(moduleStoreKey),
		parameterManager:      parameterManager,
		authenticateAuxiliary: authenticateAuxiliary,
		authorizeAuxiliary:    authorizeAuxiliary,
		bondAuxiliary:         bondAuxiliary,
		conformAuxiliary:      conformAuxiliary,
		mintAuxiliary:         mintAuxiliary,
	}
)

func setContext() sdkTypes.Context {
	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(moduleStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(authStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(bankStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, storeTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, storeTypes.StoreTypeTransient, memDB)
	_ = commitMultiStore.LoadLatestVersion()
	return sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{ChainID: ChainID}, false, log.NewNopLogger())
}

func TestTransactionKeeperTransact(t *testing.T) {

	type args struct {
		from             sdkTypes.AccAddress
		fromID           ids.IdentityID
		toID             ids.IdentityID
		classificationID ids.ClassificationID
		immutableProps   lists.PropertyList
		mutableProps     lists.PropertyList
	}
	tests := []struct {
		name    string
		args    args
		setup   func(*testing.T)
		want    *TransactionResponse
		wantErr helpers.Error
	}{
		{
			name: "minting not enabled",
			args: args{
				from:             genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			setup: func(t *testing.T) {
				parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(false)))))
			},
			wantErr: errorConstants.NotAuthorized,
		},
		{
			name: "authorization failure",
			args: args{
				from:             genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: authorizeAuxiliaryFailureClassificationID,
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			setup: func(t *testing.T) {
				parameterManager.Set(sdkTypes.WrapSDKContext(Context), baseLists.NewParameterList(base.NewParameter(baseProperties.NewMetaProperty(constantProperties.MintEnabledProperty.GetKey(), baseData.NewBooleanData(true)))))
			},
			wantErr: errorConstants.MockError,
		},
		{
			name: "authentication failure",
			args: args{
				from:             authenticateAuxiliaryFailureAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			setup:   func(t *testing.T) {},
			wantErr: errorConstants.MockError,
		},
		{
			name: "asset already exists",
			args: args{
				from:             genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("prop1"), baseData.NewStringData("val1"))),
				mutableProps:     baseLists.NewPropertyList(),
			},
			setup: func(t *testing.T) {
				assetID := baseIDs.NewAssetID(baseIDs.PrototypeClassificationID(), baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("prop1"), baseData.NewStringData("val1")))))
				assets := TransactionKeeper.mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Fetch(key.NewKey(assetID))
				assets.Add(record.NewRecord(
					baseDocuments.NewOrder(baseIDs.PrototypeClassificationID(),
						baseQualified.NewImmutables(baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("prop1"), baseData.NewStringData("val1")))),
						baseQualified.NewMutables(baseLists.NewPropertyList()))))
			},
			wantErr: errorConstants.EntityAlreadyExists,
		},
		{
			name: "conform auxiliary failure",
			args: args{
				from:             genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: conformAuxiliaryFailureID,
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(),
			},
			setup:   func(t *testing.T) {},
			wantErr: errorConstants.MockError,
		},
		{
			name: "asset supply negative",
			args: args{
				from:             genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.SupplyProperty.GetID().GetKey(), baseData.NewNumberData(sdkTypes.NewInt(-1)))),
			},
			setup:   func(t *testing.T) {},
			wantErr: errorConstants.IncorrectFormat,
		},
		{
			name: "bond auxiliary failure",
			args: args{
				from:             bondAuxiliaryFailureAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(baseProperties.NewMetaProperty(constantProperties.BondAmountProperty.GetID().GetKey(), baseData.NewNumberData(sdkTypes.NewInt(1)))),
			},
			setup: func(t *testing.T) {
			},
			wantErr: errorConstants.MockError,
		},
		{
			name: "no bond amount property",
			args: args{
				from:             genesisAddress,
				toID:             baseIDs.PrototypeIdentityID(),
				classificationID: baseIDs.PrototypeClassificationID(),
				immutableProps:   baseLists.NewPropertyList(),
				mutableProps:     baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("prop1"), baseData.NewNumberData(sdkTypes.NewInt(1)))),
			},
			setup:   func(t *testing.T) {},
			wantErr: errorConstants.MetaDataError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.setup(t)
			//fmt.Print(tt.args.classificationID)
			got, err := TransactionKeeper.Transact(sdkTypes.WrapSDKContext(Context), NewMessage(
				tt.args.from,
				baseIDs.PrototypeIdentityID(),
				tt.args.toID,
				tt.args.classificationID,
				baseLists.NewPropertyList(),
				tt.args.immutableProps,
				baseLists.NewPropertyList(),
				tt.args.mutableProps,
			).(helpers.Message))

			if (err != nil) && !tt.wantErr.Is(err) {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Error("unexpected response")
			}
		})
	}
}
