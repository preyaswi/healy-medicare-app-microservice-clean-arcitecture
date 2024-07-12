package repository

import (
	"chatservice/pkg/models"
	"chatservice/pkg/repository/interfaces"
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type chatRepository struct {
	MessageCollection *mongo.Collection
}

func NewChatRepository(DB *mongo.Database) interfaces.ChatRepository {
	return &chatRepository{
		MessageCollection: DB.Collection("messages"),
	}
}
func (c *chatRepository) StoreFriendsChat(message models.MessageReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.MessageCollection.InsertOne(ctx, message)
	if err != nil {
		return err
	}
	return nil
}

func (c *chatRepository) UpdateReadAsMessage(userid, friendid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.MessageCollection.UpdateMany(ctx, bson.M{"senderid": bson.M{"$in": bson.A{friendid}}, "recipientid": bson.M{"$in": bson.A{userid}}}, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "send"}}}})
	if err != nil {
		return err
	}
	return nil
}

func (c *chatRepository) GetFriendChat(userID, friendID string, pagination models.Pagination) ([]models.Message, error) {

	var messages []models.Message
	filter := bson.M{"senderid": bson.M{"$in": bson.A{userID, friendID}}, "recipientid": bson.M{"$in": bson.A{friendID, userID}}}
	limit, _ := strconv.Atoi(pagination.Limit)
	offset, _ := strconv.Atoi(pagination.OffSet)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	option := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	cursor, err := c.MessageCollection.Find(ctx, filter, options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}}), option)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var message models.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
