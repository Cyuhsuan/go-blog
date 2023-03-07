package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PostReply Entities
type PostReply struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	PostId    primitive.ObjectID `json:"post_id" bson:"post_id" binding:"required"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}

type PostReplyRepository interface {
	FindAll(id string) ([]*PostReply, error) // 搜尋文章底下的留言
	FindById(id string) (*PostReply, error)  // 搜尋指定留言
	Create(data *PostReply) error
	UpdateById(data *PostReply, id string) error
	DeleteById(id string) error
}

type PostReplInteractor struct {
	repository PostReplyRepository
}

func NewPostReplyInteractor(prr PostReplyRepository) *PostReplInteractor {
	return &PostReplInteractor{prr}
}

func (pri *PostReplInteractor) CreateReply(data *PostReply) error {
	if err := pri.repository.Create(data); err != nil {
		return err
	}
	return nil
}

func (pri *PostReplInteractor) UpdateReply(data *PostReply, id string) error {
	if err := pri.repository.UpdateById(data, id); err != nil {
		return err
	}
	return nil
}

func (pri *PostReplInteractor) DeleteReply(id string) error {
	if err := pri.repository.DeleteById(id); err != nil {
		return err
	}
	return nil
}

func (pri *PostReplInteractor) FindPostReply(id string) ([]*PostReply, error) {
	if res, err := pri.repository.FindAll(id); err != nil {
		return []*PostReply{}, err
	} else {
		return res, nil
	}
}

func (pri *PostReplInteractor) FindReply(id string) (*PostReply, error) {
	if res, err := pri.repository.FindById(id); err != nil {
		return &PostReply{}, err
	} else {
		return res, nil
	}
}

// mongodb的 post reply model
type MongodbPostReplyModel struct {
	collection *mongo.Collection
}

func NewMongoPostReplyRepository(collection *mongo.Collection) PostReplyRepository {
	return &MongodbPostReplyModel{collection}
}

func (prm *MongodbPostReplyModel) FindAll(id string) ([]*PostReply, error) {
	ctx := context.TODO()
	oid, _ := primitive.ObjectIDFromHex(id)

	cursor, err := prm.collection.Find(ctx, bson.M{"post_id": oid})
	if err != nil {
		return []*PostReply{}, err
	}
	var results []*PostReply

	for cursor.Next(ctx) {
		var elem *PostReply
		err := cursor.Decode(&elem)
		if err != nil {
			return []*PostReply{}, errors.New("query data error")
		}

		results = append(results, elem)

	}
	return results, nil
}
func (prm *MongodbPostReplyModel) FindById(id string) (*PostReply, error)
func (prm *MongodbPostReplyModel) Create(data *PostReply) error
func (prm *MongodbPostReplyModel) UpdateById(data *PostReply, id string) error
func (prm *MongodbPostReplyModel) DeleteById(id string) error
