// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package issue

import (
	"context"
	"fmt"
	"github.com/AssetMantle/modules/x/identities/mapper"
	"github.com/AssetMantle/modules/x/identities/record"
	"reflect"
	"testing"

	baseData "github.com/AssetMantle/schema/go/data/base"
	baseDocuments "github.com/AssetMantle/schema/go/documents/base"
	baseIDs "github.com/AssetMantle/schema/go/ids/base"
	baseLists "github.com/AssetMantle/schema/go/lists/base"
	baseProperties "github.com/AssetMantle/schema/go/properties/base"
	baseQualified "github.com/AssetMantle/schema/go/qualified/base"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	protoTendermintTypes "github.com/tendermint/tendermint/proto/tendermint/types"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/helpers"
	baseHelpers "github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/x/classifications/auxiliaries/conform"
	"github.com/AssetMantle/modules/x/identities/auxiliaries/authenticate"
	"github.com/AssetMantle/modules/x/identities/parameters"
	"github.com/AssetMantle/modules/x/maintainers/auxiliaries/authorize"
)

type TestKeepers struct {
	IssueKeeper helpers.TransactionKeeper
}

var (
	parameterManager      helpers.ParameterManager
	authenticateAuxiliary helpers.Auxiliary
	conformAuxiliary      helpers.Auxiliary
	authorizeAuxiliary    helpers.Auxiliary
)

func CreateTestInput(t *testing.T) (sdkTypes.Context, TestKeepers, helpers.Mapper, helpers.ParameterManager) {
	var legacyAmino = baseHelpers.CodecPrototype().GetLegacyAmino()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")
	Mapper := mapper.Prototype().Initialize(storeKey)
	encodingConfig := simapp.MakeTestEncodingConfig()
	appCodec := encodingConfig.Marshaler
	ParamsKeeper := paramsKeeper.NewKeeper(
		appCodec,
		legacyAmino,
		paramsStoreKey,
		paramsTransientStoreKeys,
	)
	parameterManager = parameters.Prototype().Initialize(ParamsKeeper.Subspace("test"))

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	Context := sdkTypes.NewContext(commitMultiStore, protoTendermintTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	authenticateAuxiliary = authenticate.Auxiliary.Initialize(Mapper, parameterManager)
	conformAuxiliary = conform.Auxiliary.Initialize(Mapper, parameterManager)
	authorizeAuxiliary = authorize.Auxiliary.Initialize(Mapper, parameterManager)
	keepers := TestKeepers{
		IssueKeeper: keeperPrototype().Initialize(Mapper, parameterManager, []interface{}{authenticateAuxiliary, conformAuxiliary, authorizeAuxiliary}).(helpers.TransactionKeeper),
	}

	return Context, keepers, Mapper, parameterManager
}

func Test_keeperPrototype(t *testing.T) {
	tests := []struct {
		name string
		want helpers.TransactionKeeper
	}{
		{"+ve", transactionKeeper{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keeperPrototype(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keeperPrototype() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Initialize(t *testing.T) {
	_, _, Mapper, parameterManager := CreateTestInput(t)
	type fields struct {
		mapper             helpers.Mapper
		authorizeAuxiliary helpers.Auxiliary
		bondAuxiliary      helpers.Auxiliary
		conformAuxiliary   helpers.Auxiliary
	}
	type args struct {
		mapper      helpers.Mapper
		in1         helpers.ParameterManager
		auxiliaries []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   helpers.Keeper
	}{
		{"+ve with nil", fields{}, args{}, transactionKeeper{}},
		{"+ve", fields{Mapper, conformAuxiliary, conformAuxiliary, authorizeAuxiliary}, args{Mapper, parameterManager, []interface{}{}}, transactionKeeper{Mapper, parameterManager, conformAuxiliary, authorizeAuxiliary, authorizeAuxiliary}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:             tt.fields.mapper,
				parameterManager:   parameterManager,
				authorizeAuxiliary: tt.fields.authorizeAuxiliary,
				bondAuxiliary:      tt.fields.bondAuxiliary,
				conformAuxiliary:   tt.fields.conformAuxiliary,
			}
			if got := transactionKeeper.Initialize(tt.args.mapper, tt.args.in1, tt.args.auxiliaries); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("Initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_transactionKeeper_Transact(t *testing.T) {
	Context, keepers, Mapper, parameterManager := CreateTestInput(t)
	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, err := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, err)
	immutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("ID1"), baseData.NewStringData("ImmutableData")))
	immutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("ID11"), baseData.NewStringData("ImmutableData")))
	mutableMetaProperties := baseLists.NewPropertyList(baseProperties.NewMetaProperty(baseIDs.NewStringID("authentication"), baseData.NewListData()))
	mutableProperties := baseLists.NewPropertyList(baseProperties.NewMesaProperty(baseIDs.NewStringID("authentication"), baseData.NewStringData("MutableData")))
	immutables := baseQualified.NewImmutables(immutableMetaProperties)
	mutables := baseQualified.NewMutables(mutableMetaProperties)
	classificationID := baseIDs.NewClassificationID(immutables, mutables)
	fromIdentityID := baseIDs.NewIdentityID(classificationID, immutables)
	identity := baseDocuments.NewIdentity(classificationID, immutables, mutables)
	identity = identity.ProvisionAddress([]sdkTypes.AccAddress{fromAccAddress}...)
	identity.Mutate(baseLists.AnyPropertiesToProperties(immutableMetaProperties.Get()...)...)
	keepers.IssueKeeper.(transactionKeeper).mapper.NewCollection(sdkTypes.WrapSDKContext(Context)).Add(record.NewRecord(identity))
	type fields struct {
		mapper             helpers.Mapper
		parameterManager   helpers.ParameterManager
		conformAuxiliary   helpers.Auxiliary
		bondAuxiliary      helpers.Auxiliary
		authorizeAuxiliary helpers.Auxiliary
	}
	type args struct {
		context context.Context
		message helpers.Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    helpers.TransactionResponse
		wantErr bool
	}{
		// NOTE: When test individually run 2nd test will fail
		{"+ve", fields{Mapper, parameterManager, authenticateAuxiliary, conformAuxiliary, authorizeAuxiliary}, args{Context.Context(), NewMessage(fromAccAddress, fromIdentityID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).(*Message)}, newTransactionResponse(nil), false},
		{"+ve Entity Already Exists", fields{Mapper, parameterManager, authenticateAuxiliary, conformAuxiliary, authorizeAuxiliary}, args{Context.Context(), NewMessage(fromAccAddress, fromIdentityID, classificationID, immutableMetaProperties, immutableProperties, mutableMetaProperties, mutableProperties).(*Message)}, newTransactionResponse(nil), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transactionKeeper := transactionKeeper{
				mapper:             tt.fields.mapper,
				parameterManager:   tt.fields.parameterManager,
				conformAuxiliary:   tt.fields.conformAuxiliary,
				bondAuxiliary:      tt.fields.bondAuxiliary,
				authorizeAuxiliary: tt.fields.authorizeAuxiliary,
			}
			got, err := transactionKeeper.Transact(tt.args.context, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Transact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transact() got = %v, want %v", got, tt.want)
			}
		})
	}
}
