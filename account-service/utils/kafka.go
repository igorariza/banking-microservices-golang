package utils

import (
	"banking-system/account-service/data/models"
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	kafkaBroker = os.Getenv("KAFKA_BROKER")
	topic       = os.Getenv("CREATE_ACCOUNT_TOPIC")
)

func CreateTopic(brokerAddress, topicName string, numPartitions, replicationFactor int) error {
	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(10 * time.Second))

	err = conn.CreateTopics(kafka.TopicConfig{
		Topic:             topicName,
		NumPartitions:     numPartitions,
		ReplicationFactor: replicationFactor,
	})
	if err != nil {
        log.Printf("failed to create topic: %v", err)
		return err
	}

	log.Printf("Topico '%s' creado exitosamente", topicName)
	return nil
}

func PublishAccountEvent(ctx context.Context, accountID string) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaBroker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	jsonAccount, _ := json.Marshal(models.Account{ID: accountID})

	message := kafka.Message{
		Key:   []byte(topic),
		Value: []byte(jsonAccount),
	}

	err := writer.WriteMessages(ctx, message)
	if err != nil {
		log.Printf("failed to write messages: %v", err)
		return err
	}

	log.Printf("published create account event: %s", accountID)
	return nil
}
