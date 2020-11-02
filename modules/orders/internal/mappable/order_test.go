package mappable

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Order_Methods(t *testing.T) {

	classificationID := base.NewID("classificationID")
	makerOwnableID := base.NewID("makerOwnableID")
	takerOwnableID := base.NewID("takerOwnableID")
	makerID := base.NewID("makerID")

	takerIDImmutableProperty := base.NewProperty(base.NewID(properties.TakerID), base.NewFact(base.NewStringData("takerIDImmutableProperty")))
	exchangeRateImmutableProperty := base.NewProperty(base.NewID(properties.ExchangeRate), base.NewFact(base.NewStringData("exchangeRateImmutableProperty")))
	creationImmutableProperty := base.NewProperty(base.NewID(properties.Creation), base.NewFact(base.NewStringData("creationImmutableProperty")))
	expiryImmutableProperty := base.NewProperty(base.NewID(properties.Expiry), base.NewFact(base.NewStringData("expiryImmutableProperty")))
	makerOwnableSplitImmutableProperty := base.NewProperty(base.NewID(properties.MakerOwnableSplit), base.NewFact(base.NewStringData("makerOwnableSplitImmutableProperty")))

	takerIDMutableProperty := base.NewProperty(base.NewID(properties.TakerID), base.NewFact(base.NewStringData("takerIDMutableProperty")))
	exchangeRateMutableProperty := base.NewProperty(base.NewID(properties.ExchangeRate), base.NewFact(base.NewStringData("exchangeRateMutableProperty")))
	creationMutableProperty := base.NewProperty(base.NewID(properties.Creation), base.NewFact(base.NewStringData("creationMutableProperty")))
	expiryMutableProperty := base.NewProperty(base.NewID(properties.Expiry), base.NewFact(base.NewStringData("expiryMutableProperty")))
	makerOwnableSplitMutableProperty := base.NewProperty(base.NewID(properties.MakerOwnableSplit), base.NewFact(base.NewStringData("makerOwnableSplitMutableProperty")))

	immutables := base.NewImmutables(base.NewProperties(takerIDImmutableProperty, exchangeRateImmutableProperty, creationImmutableProperty, expiryImmutableProperty, makerOwnableSplitImmutableProperty))
	mutables := base.NewMutables(base.NewProperties(takerIDMutableProperty, exchangeRateMutableProperty, creationMutableProperty, expiryMutableProperty, makerOwnableSplitMutableProperty))
	testOrderID := key.NewOrderID(classificationID, makerOwnableID, takerOwnableID, makerID, immutables)
	testOrder := NewOrder(testOrderID, immutables, base.NewMutables(base.NewProperties())).(order)
	testOrder2 := NewOrder(testOrderID, base.NewImmutables(base.NewProperties()), mutables).(order)
	testOrder3 := NewOrder(testOrderID, base.NewImmutables(base.NewProperties()), base.NewMutables(base.NewProperties())).(order)

	data, _ := base.ReadIDData("")
	defaultTakerProperty := base.NewProperty(base.NewID(properties.TakerID), base.NewFact(data))
	defaultExchangeRateProperty := base.NewProperty(base.NewID(properties.ExchangeRate), base.NewFact(base.NewDecData(sdkTypes.OneDec())))
	data, _ = base.ReadHeightData("")
	defaultCreationProperty := base.NewProperty(base.NewID(properties.Creation), base.NewFact(data))
	defaultExpiryProperty := base.NewProperty(base.NewID(properties.Expiry), base.NewFact(data))
	data, _ = base.ReadDecData("")
	defaultMakerOwnableSplitProperty := base.NewProperty(base.NewID(properties.MakerOwnableSplit), base.NewFact(data))

	require.Equal(t, order{ID: testOrderID, Immutables: immutables, Mutables: base.NewMutables(base.NewProperties())}, testOrder)
	require.Equal(t, testOrderID, testOrder.GetID())
	require.Equal(t, testOrderID, testOrder.GetKey())
	require.Equal(t, classificationID, testOrder.GetClassificationID())
	require.Equal(t, makerOwnableID, testOrder.GetMakerOwnableID())
	require.Equal(t, takerOwnableID, testOrder.GetTakerOwnableID())
	require.Equal(t, makerID, testOrder.GetMakerID())

	//GetTakerID
	require.Equal(t, takerIDImmutableProperty, testOrder.GetTakerID())
	require.Equal(t, takerIDMutableProperty, testOrder2.GetTakerID())
	require.Equal(t, defaultTakerProperty, testOrder3.GetTakerID())
	//GetExchangeRate
	require.Equal(t, exchangeRateImmutableProperty, testOrder.GetExchangeRate())
	require.Equal(t, exchangeRateMutableProperty, testOrder2.GetExchangeRate())
	require.Equal(t, defaultExchangeRateProperty, testOrder3.GetExchangeRate())

	//GetCreation
	require.Equal(t, creationImmutableProperty, testOrder.GetCreation())
	require.Equal(t, creationMutableProperty, testOrder2.GetCreation())
	require.Equal(t, defaultCreationProperty, testOrder3.GetCreation())

	//GetExpiry
	require.Equal(t, expiryImmutableProperty, testOrder.GetExpiry())
	require.Equal(t, expiryMutableProperty, testOrder2.GetExpiry())
	require.Equal(t, defaultExpiryProperty, testOrder3.GetExpiry())

	//GetMakerOwnableSplit
	require.Equal(t, makerOwnableSplitImmutableProperty, testOrder.GetMakerOwnableSplit())
	require.Equal(t, makerOwnableSplitMutableProperty, testOrder2.GetMakerOwnableSplit())
	require.Equal(t, defaultMakerOwnableSplitProperty, testOrder3.GetMakerOwnableSplit())

	require.Equal(t, immutables, testOrder.GetImmutables())
	require.Equal(t, mutables, testOrder2.GetMutables())
	require.Equal(t, testOrderID, testOrder2.GetID())
	require.Equal(t, testOrderID, testOrder2.GetKey())

}
