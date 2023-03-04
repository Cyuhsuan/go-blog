package models

import (
	"context"
	"errors"
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
}

func (p *Post) FindAll(collection *mongo.Collection) ([]Post, error) {
	cursor, err := collection.Find(context.TODO(), bson.D{})
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
func (p *Post) FindById(collection *mongo.Collection, id string) (Post, error) {
	oid, _ := primitive.ObjectIDFromHex(id)

	var post Post

	query := bson.M{"_id": oid}
	var ctx context.Context
	err := collection.FindOne(ctx, query).Decode(&post)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Post{}, err
		}
		return Post{}, err
	}

	return post, nil
}
func (p *Post) Create(collection *mongo.Collection) error {
	var ctx context.Context
	_, err := collection.InsertOne(ctx, &p)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return errors.New("user with that email already exist")
		}
		return err
	}
	return nil
}
func (p *Post) UpdateById() error {

	return nil
}
func (p *Post) DeleteById() error {
	return nil
}

// func (p *Post) Index() {

// }
// func (p *Post) Show(id primitive.ObjectID) {

// }
// func (p *Post) Store() {

// }
// func (p *Post) Update() {

// }
// func (p *Post) Delete(id primitive.ObjectID) {

// }
