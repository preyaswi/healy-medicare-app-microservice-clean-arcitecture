package models

import "time"

type MessageReq struct {
	SenderID    string    `bson:"senderid"`
	RecipientID string    `bson:"recipientid"`
	Content     string    `bson:"content"`
	Timestamp   time.Time `bson:"timestamp"`
}

type Pagination struct {
	Limit  string
	OffSet string
}

type Message struct {
	ID          string    `bson:"_id"`
	SenderID    string    `bson:"senderid"`
	RecipientID string    `bson:"recipientid"`
	Content     string    `bson:"content"`
	Timestamp   time.Time `bson:"timestamp"`
}
