// Copyright [2021] - [2025], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package queuing

import (
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/helpers/base"
)

func TestKafkaTopicConsumer(t *testing.T) {
	testConsumers := []string{"testConsumers"}

	require.Panics(t, func() {
		testKafkaState := newKafkaState(testConsumers)
		partitionConsumer := testKafkaState.Consumers["Topic"]

		var kafkaStore kafkaMsg
		if len(partitionConsumer.Messages()) == 0 {
			kafkaStore = kafkaMsg{Msg: nil}
		}

		kafkaMsg := <-partitionConsumer.Messages()

		err := base.CodecPrototype().GetLegacyAmino().UnmarshalJSON(kafkaMsg.Value, &kafkaStore)
		if err != nil {
			panic(err)
		}

		require.Equal(t, kafkaTopicConsumer("Topic", testKafkaState.Consumers, base.CodecPrototype().GetLegacyAmino()), kafkaStore)
	})
}

func TestNewConsumer(t *testing.T) {
	consumers := []string{"testConsumers"}
	config := sarama.NewConfig()

	consumer, _ := sarama.NewConsumer(consumers, config)

	// TODO: Add test cases.
	// require.Nil(t, err, "should not happened. err %v", err)

	require.Panics(t, func() {
		require.Equal(t, newConsumer(consumers), consumer)
	})
}
