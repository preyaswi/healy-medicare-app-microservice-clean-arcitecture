package client

import (
	"context"
	"fmt"
	interfaces "healy-apigateway/pkg/client/interface"
	"healy-apigateway/pkg/config"
	pb "healy-apigateway/pkg/pb/chat"
	models "healy-apigateway/pkg/utils"

	"google.golang.org/grpc"
)

type chatClient struct {
	Client pb.ChatClient
}

func NewChatClient(cfg config.Config) interfaces.ChatClient {
	grpcConnection, err := grpc.Dial(cfg.ChatSvc, grpc.WithInsecure())
	if err != nil {
		fmt.Println("could not connect", err)
	}
	grpcClient := pb.NewChatClient(grpcConnection)

	return &chatClient{
		Client: grpcClient,
	}
}

func (ad *chatClient) GetChat(userid string, req models.ChatRequest) ([]models.TempMessage, error) {
	fmt.Println("zzzzzzzzzzz", userid, req)
	data, err := ad.Client.GetFriendChat(context.Background(), &pb.GetFriendChatRequest{
		UserID:   userid,
		FriendID: req.FriendID,
		OffSet:   req.Offset,
		Limit:    req.Limit,
	})
	if err != nil {
		fmt.Println(data, "hello")
		return []models.TempMessage{}, err
	}
	var response []models.TempMessage

	for _, v := range data.FriendChat {
		chatResponse := models.TempMessage{
			SenderID:    v.SenderId,
			RecipientID: v.RecipientId,
			Content:     v.Content,
			Timestamp:   v.Timestamp,
		}
		response = append(response, chatResponse)
	}
	return response, nil
}
