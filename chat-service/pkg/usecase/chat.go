package usecase

import (
	"chatservice/pkg/config"
	"chatservice/pkg/helper"
	"chatservice/pkg/models"
	"chatservice/pkg/repository/interfaces"
	services "chatservice/pkg/usecase/interfaces"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/IBM/sarama"
)

type exampleConsumerGroupHandler struct {
	chatUseCase *chatusecase
}

func (h exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (h exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Received message: %s\n", string(msg.Value))
		chatMsg, err := h.chatUseCase.UnmarshelChatMessage(msg.Value)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}
		err = h.chatUseCase.chatRepository.StoreFriendsChat(*chatMsg)
		if err != nil {
			log.Printf("Error storing message in repository: %v", err)
			continue
		}
		sess.MarkMessage(msg, "")
	}
	return nil
}

type chatusecase struct {
	chatRepository interfaces.ChatRepository
}

func NewChatUseCase(repository interfaces.ChatRepository) services.ChatUseCase {
	return &chatusecase{
		chatRepository: repository,
	}
}

func (c *chatusecase) MessageConsumer() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading config: %v", err)
		return
	}

	brokers := []string{cfg.KafkaPort}
	topic := []string{"CHAT-TOPIC"}

	configs := sarama.NewConfig()
	configs.Version = sarama.V2_0_0_0
	configs.Consumer.Offsets.AutoCommit.Enable = true
	configs.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRange()
	configs.Consumer.Offsets.Initial = sarama.OffsetOldest

	log.Printf("Attempting to create consumer group with brokers: %v", brokers)

	admin, err := sarama.NewClusterAdmin(brokers, configs)
	if err != nil {
		log.Printf("Error creating cluster admin: %v", err)
	}
	defer admin.Close()

	// CREATE TOPIC USING
	err = admin.CreateTopic("CHAT-TOPIC", &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false)

	if err != nil {
		// If the topic already exists, this is not an error
		if err != sarama.ErrTopicAlreadyExists {
			log.Printf("Error creating topic: %v", err)
		}
	}

	consumerGroup, err := sarama.NewConsumerGroup(brokers, "chat-consumer-group", configs)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}
	defer consumerGroup.Close()

	ctx := context.Background()
	handler := exampleConsumerGroupHandler{chatUseCase: c}

	log.Println("Starting to consume messages")

	for {
		err := consumerGroup.Consume(ctx, topic, handler)
		if err != nil {
			log.Printf("Error from consumer: %v", err)
			time.Sleep(time.Second * 5)
		}
	}
}

func (c *chatusecase) UnmarshelChatMessage(data []byte) (*models.MessageReq, error) {
	var message models.MessageReq
	err := json.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}
	message.Timestamp = time.Now()
	return &message, nil
}

func (ad *chatusecase) GetFriendChat(userid, friendid string, pagination models.Pagination) ([]models.Message, error) {
	var err error
	pagination.OffSet, err = helper.Pagination(pagination.Limit, pagination.OffSet)
	if err != nil {
		return nil, err
	}
	_ = ad.chatRepository.UpdateReadAsMessage(userid, friendid)
	return ad.chatRepository.GetFriendChat(userid, friendid, pagination)
}
