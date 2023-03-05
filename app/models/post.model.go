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
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostModel struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewPostModel(collection *mongo.Collection, ctx context.Context) *PostModel {
	return &PostModel{collection, ctx}
}

func (m *PostModel) FindAll() ([]Post, error) {
	cursor, err := m.collection.Find(m.ctx, bson.D{})
	if err != nil {
		return []Post{}, errors.New("query error")
	}
	var results []Post
	for cursor.Next(m.ctx) {
		var elem Post
		err := cursor.Decode(&elem)
		if err != nil {
			return []Post{}, errors.New("query data error")
		}

		results = append(results, elem)

	}
	return results, nil
}
func (m *PostModel) FindById(id string) (Post, error) {
	oid, _ := primitive.ObjectIDFromHex(id)

	var post Post

	query := bson.M{"_id": oid}
	err := m.collection.FindOne(m.ctx, query).Decode(&post)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Post{}, err
		}
		return Post{}, err
	}

	return post, nil
}
func (m *PostModel) Create(data validation.PostCreateForm) error {
	_, err := m.collection.InsertOne(m.ctx, data)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return errors.New("user with that email already exist")
		}
		return err
	}
	return nil
}
func (m *PostModel) UpdateById(data validation.PostCreateForm, id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", oid}}
	var updateFields bson.D
	conv, _ := bson.Marshal(data)
	bson.Unmarshal(conv, &updateFields)
	update := bson.D{{"$set", updateFields}}

	_, err := m.collection.UpdateOne(m.ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
func (m *PostModel) DeleteById(id string) error {
	oid, _ := primitive.ObjectIDFromHex(id)

	_, err := m.collection.DeleteOne(m.ctx, bson.D{{"_id", oid}})
	if err != nil {
		return err
	}
	return nil
}
