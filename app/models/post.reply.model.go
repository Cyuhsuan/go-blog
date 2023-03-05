package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostReply struct {
	PostId      primitive.ObjectID
	PostReplyId primitive.ObjectID
	Content     string
	CreatedAt   time.Time
}
