package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strings"
	"testing"
)

const (
	brokerAddr = "localhost:9092"
	topic      = "my_topic"
	msg        = "test_message"
)

func TestKafkaSendMessage(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	addr := strings.Split(brokerAddr, "m")
	t.Log("send message")
	if err := SendMessage(addr, config, topic, sarama.ByteEncoder(msg)); err != nil {
		t.Error(err)
	}

}

func TestKafkaReceiveMessage(t *testing.T) {
	addr := strings.Split(brokerAddr, "m")
	t.Log("receive message")
	for {
		// sarama.OffsetNewest：从当前的偏移量开始消费，sarama.OffsetOldest：从最老的偏移量开始消费
		msg, err := Consumer(addr, topic, 0, sarama.OffsetNewest)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(fmt.Sprintf("msg offset: %d, partition: %d, timestamp: %s, value: %s", msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value)))
		}
	}
}
