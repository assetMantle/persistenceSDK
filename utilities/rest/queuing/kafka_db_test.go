// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queuing

import (
	"testing"

	baseIDs "github.com/AssetMantle/schema/ids/base"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers/base"
	"github.com/AssetMantle/modules/utilities/random"
)

func Test_Kafka_DB(t *testing.T) {
	require.Panics(t, func() {
		var legacyAmino = base.CodecPrototype().GetLegacyAmino()
		ticketID := TicketID(random.GenerateUniqueIdentifier("name"))
		kafkaDB, _ := dbm.NewGoLevelDB("KafkaDB", defaultCLIHome)
		setTicketIDtoDB(ticketID, kafkaDB, legacyAmino, []byte{})
		addResponseToDB(ticketID, baseIDs.NewStringID("").Bytes(), kafkaDB, legacyAmino)
		require.Equal(t, baseIDs.NewStringID("").Bytes(), getResponseFromDB(ticketID, kafkaDB, legacyAmino))
	})
}
