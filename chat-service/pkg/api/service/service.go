package service

import (
	"chatservice/pkg/logging"
	"chatservice/pkg/models"
	"chatservice/pkg/pb"
	"chatservice/pkg/usecase/interfaces"
	"context"
	"time"
)

type ChatServer struct {
	chatUseCase interfaces.ChatUseCase
	pb.UnimplementedChatServer
}

func NewChatServer(usecase interfaces.ChatUseCase) pb.ChatServer {
	return &ChatServer{
		chatUseCase: usecase,
	}
}
func (ad *ChatServer) GetFriendChat(ctx context.Context, req *pb.GetFriendChatRequest) (*pb.GetFriendChatResponse, error) {
	logEntry := logging.Logger().WithField("method", "GetFriendChat")
	logEntry.Info("Processing GetFriendChat request with limit:", req.GetLimit(), "and offset:", req.GetOffSet())
	ind, _ := time.LoadLocation("Asia/Kolkata")
	result, err := ad.chatUseCase.GetFriendChat(req.UserID, req.FriendID, models.Pagination{Limit: req.Limit, OffSet: req.OffSet})
	if err != nil {
		logEntry.WithError(err).Errorf("Error getting friend chat ")

		return nil, err
	}
	var finalResult []*pb.Message
	for _, val := range result {
		finalResult = append(finalResult, &pb.Message{
			MessageID:   val.ID,
			SenderId:    val.SenderID,
			RecipientId: val.RecipientID,
			Content:     val.Content,
			Timestamp:   val.Timestamp.In(ind).String(),
		})
	}
	return &pb.GetFriendChatResponse{FriendChat: finalResult}, nil
}
