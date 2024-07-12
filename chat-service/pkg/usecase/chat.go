package usecase

import (
	"chatservice/pkg/config"
	"chatservice/pkg/helper"
	"chatservice/pkg/models"
	"chatservice/pkg/repository/interfaces"
	services "chatservice/pkg/usecase/interfaces"
	"encoding/json"
	"fmt"
	"time"

	"github.com/IBM/sarama"
)

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
		fmt.Println("errorn loading config", err)
		return
	}

	configs := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{cfg.KafkaPort}, configs)
	if err != nil {
		fmt.Println("Error creatig Kafka consumer:", err)
		return
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(cfg.KafkaTpic, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("Error creating partition consumer:", err)
		return
	}
	defer partitionConsumer.Close()
	fmt.Println("Kafka consumer started")
	for {
		select {
		case message := <-partitionConsumer.Messages():
			msg, err := c.UnmarshelChatMessage(message.Value)
			if err != nil {
				fmt.Println("Error unmarshalling message", err)
				continue
			}
			err = c.chatRepository.StoreFriendsChat(*msg)

			if err != nil {
				fmt.Println("Error storing message in repository:", err)
				continue
			}
		case err := <-partitionConsumer.Errors():
			fmt.Println("kafka consumer error", err)
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
