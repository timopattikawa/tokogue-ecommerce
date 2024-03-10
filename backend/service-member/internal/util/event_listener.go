package util

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type KafkaUsecase interface {
	ProduceMessage([]byte) bool
}

type KafkaConfig struct {
	KafkaEvent *kafka.Conn
}

func NewKafkaConfig(addr string, topic string) KafkaUsecase {
	conn, err := kafka.DialLeader(context.Background(), "tcp", addr, topic, 0)
	if err != nil {
		log.Printf("Kafka err : {%s}", err)
	}

	return &KafkaConfig{
		KafkaEvent: conn,
	}
}

func (k KafkaConfig) ProduceMessage(message []byte) bool {
	err := k.KafkaEvent.SetWriteDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Printf("Error kafka err {%s}", err)
		return false
	}
	_, err = k.KafkaEvent.WriteMessages(
		kafka.Message{Value: message})

	if err != nil {
		log.Printf("Error send message %s : err {%s}", message, err)
		return false
	}

	return true
}
