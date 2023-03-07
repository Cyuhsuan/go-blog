package models

import (
	"context"
	"errors"
	"go-blog/app/validation"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
	Title     string    `bson:"title"`
	Content   string    `bson:"content"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type PostModel struct {
	collection *mongo.Collection
}

func NewPostModel(collection *mongo.Collection) *PostModel {
	return &PostModel{collection}
}

func (m *PostModel) FindAll(ctx context.Context) ([]Post, error) {
	cursor, err := m.collection.Find(ctx, bson.D{})
	if err != nil {
		return []Post{}, errors.New("query error")
	}
	var results []Post
	for cursor.Next(ctx) {
		var elem Post
		err := cursor.Decode(&elem)
		if err != nil {
			return []Post{}, errors.New("query data error")
		}

		results = append(results, elem)

	}
	return results, nil
}
func (m *PostModel) FindById(ctx context.Context, id string) (Post, error) {
	oid, _ := primitive.ObjectIDFromHex(id)

	var post Post

	query := bson.M{"_id": oid}
	err := m.collection.FindOne(ctx, query).Decode(&post)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Post{}, err
		}
		return Post{}, err
	}

	return post, nil
}
func (m *PostModel) Create(ctx context.Context, data validation.PostCreateForm) error {
	_, err := m.collection.InsertOne(ctx, data)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return errors.New("user with that email already exist")
		}
		return err
	}
	return nil
}
func (m *PostModel) UpdateById(ctx context.Context, data validation.PostCreateForm, id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", oid}}
	var updateFields bson.D
	conv, _ := bson.Marshal(data)
	bson.Unmarshal(conv, &updateFields)
	update := bson.D{{"$set", updateFields}}

	_, err := m.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
func (m *PostModel) DeleteById(ctx context.Context, id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)

	_, err := m.collection.DeleteOne(ctx, bson.D{{"_id", oid}})
	if err != nil {
		return err
	}
	return nil
}
