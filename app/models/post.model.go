package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Post Entities
type Post struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title" binding:"required"`
	Content   string             `json:"content" bson:"content" binding:"required"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}

type PostRepository interface {
	FindAll() ([]*Post, error)
	FindById(id string) (*Post, error)
	Create(data *Post) error
	UpdateById(data *Post, id string) error
	DeleteById(id string) error
}

type PostInteractor struct {
	repository PostRepository
}

func NewPostInteractor(pr PostRepository) *PostInteractor {
	return &PostInteractor{pr}
}

func (pi *PostInteractor) CreatePost(data *Post) error {
	if err := pi.repository.Create(data); err != nil {
		return err
	}
	return nil
}

func (pi *PostInteractor) UpdatePost(data *Post, id string) error {
	if err := pi.repository.UpdateById(data, id); err != nil {
		return err
	}
	return nil
}

func (pi *PostInteractor) GetAllPost() ([]*Post, error) {
	data, err := pi.repository.FindAll()
	if err != nil {
		return []*Post{}, err
	}
	return data, nil
}

func (pi *PostInteractor) GetPostById(id string) (*Post, error) {
	data, err := pi.repository.FindById(id)
	if err != nil {
		return &Post{}, err
	}
	return data, nil
}

func (pi *PostInteractor) DeletePostById(id string) error {
	if err := pi.repository.DeleteById(id); err != nil {
		return err
	}
	return nil
}

// mongodb的post model
type MongodbPostModel struct {
	collection *mongo.Collection
}

func NewMongoPostRepository(collection *mongo.Collection) PostRepository {
	return &MongodbPostModel{collection}
}

func (pm *MongodbPostModel) FindAll() ([]*Post, error) {
	ctx := context.TODO()
	cursor, err := pm.collection.Find(ctx, bson.D{})
	if err != nil {
		return []*Post{}, errors.New("query error")
	}
	var results []*Post
	for cursor.Next(ctx) {
		var elem *Post
		err := cursor.Decode(&elem)
		if err != nil {
			return []*Post{}, errors.New("query data error")
		}

		results = append(results, elem)

	}
	return results, nil
}
func (pm *MongodbPostModel) FindById(id string) (*Post, error) {
	ctx := context.TODO()
	oid, _ := primitive.ObjectIDFromHex(id)

	var post *Post
	if err := pm.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&post); err != nil {
		return &Post{}, err
	}

	return post, nil
}
func (pm *MongodbPostModel) Create(data *Post) error {
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	ctx := context.TODO()
	if _, err := pm.collection.InsertOne(ctx, data); err != nil {
		return err
	}
	return nil
}
func (pm *MongodbPostModel) UpdateById(data *Post, id string) error {
	ctx := context.TODO()
	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", oid}}
	var updateFields bson.D
	data.UpdatedAt = time.Now()
	conv, _ := bson.Marshal(data)
	bson.Unmarshal(conv, &updateFields)
	update := bson.D{{"$set", updateFields}}
	if _, err := pm.collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}
	return nil
}
func (pm *MongodbPostModel) DeleteById(id string) error {
	ctx := context.TODO()
	oid, _ := primitive.ObjectIDFromHex(id)

	_, err := pm.collection.DeleteOne(ctx, bson.D{{"_id", oid}})
	if err != nil {
		return err
	}
	return nil
}
