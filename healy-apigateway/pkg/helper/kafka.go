package helper

import (
	"encoding/json"
	"fmt"
	"log"

	"healy-apigateway/pkg/config"
	models "healy-apigateway/pkg/utils"

	"github.com/IBM/sarama"
	"github.com/gofiber/websocket/v2"
)

// SendMessageToUser sends a message to a specific user using WebSocket and Kafka for messaging.
func SendMessageToUser(User map[string]*websocket.Conn, msg []byte, userID string) {
	fmt.Println("Processing message...")

	var message models.Message
	if err := json.Unmarshal(msg, &message); err != nil {
		fmt.Println("Error while unmarshaling message:", err)
		return
	}

	message.SenderID = userID
	fmt.Printf("Recipient ID: %s, User ID: %s\n", message.RecipientID, userID)
	recipientConn, ok := User[message.RecipientID]
	if ok {
		if err := recipientConn.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println("Error sending message to recipient:", err)
		}
	}

	if err := KafkaProducer(message); err != nil {
		fmt.Println("Error producing message to Kafka:", err)
	}
}

// KafkaProducer sends a message to a Kafka topic.
func KafkaProducer(message models.Message) error {
	fmt.Println("Sending message to Kafka...")

	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	configs := sarama.NewConfig()
	configs.Producer.Return.Successes = true
	configs.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{cfg.KafkaPort}, configs)
	if err != nil {
		return fmt.Errorf("failed to create Kafka producer: %w", err)
	}
	defer producer.Close()

	result, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: cfg.KafkaTpic,
		Key:   sarama.StringEncoder("Friend message"),
		Value: sarama.StringEncoder(result),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error sending message to Kafka:", err)
	}

	log.Printf("[producer] partition id:%d; offset:%d, value:%v\n", partition, offset, message)
	return nil
}
