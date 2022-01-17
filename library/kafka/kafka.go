package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// SendMessage 发送消息
func SendMessage(brokerAddr []string, config *sarama.Config, topic string, value sarama.Encoder) (err error) {
	producer, err := sarama.NewSyncProducer(brokerAddr, config)
	if err != nil {
		return err
	}
	defer func() {
		if err = producer.Close(); err != nil {
			return
		}
	}()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: value,
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("send message err:%s", err)
	}
	return fmt.Errorf("partition:%d offset:%d\n", partition, offset)
}

// Consumer 消费消息
func Consumer(brokenAddr []string, topic string, partition int32, offset int64) (msg *sarama.ConsumerMessage, err error) {
	consumer, err := sarama.NewConsumer(brokenAddr, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = consumer.Close(); err != nil {
			return
		}
	}()
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = partitionConsumer.Close(); err != nil {
			return
		}
	}()
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			return msg, nil
		case err := <-partitionConsumer.Errors():
			return nil, err
		default:
			return
		}
	}
}
