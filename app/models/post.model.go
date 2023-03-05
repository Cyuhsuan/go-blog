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

func NewPostModel(collection *mongo.Collection, ctx context.Context) PostModel {
	return PostModel{collection, ctx}
}

func (m *PostModel) FindAll() ([]Post, error) {
	cursor, err := m.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return []Post{}, errors.New("query error")
	}
	var results []Post
	for cursor.Next(context.TODO()) {
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
func (m *PostModel) UpdateById() error {

	return nil
}
func (m *PostModel) DeleteById() error {
	return nil
}
