package services

import (
	"context"
	"go-blog/app/models"
	"go-blog/app/validation"
	"time"

	// "github.com/example/golang-test/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type PostServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
	model      models.Post
}

func NewPostService(db *mongo.Database, ctx context.Context, model models.Post) PostService {
	return &PostServiceImpl{db.Collection("post"), ctx, model}
}

func (ps *PostServiceImpl) Index() ([]models.Post, error) {
	results, err := ps.model.FindAll(ps.collection)
	if err != nil {
		return []models.Post{}, err
	}
	return results, nil
}
func (ps *PostServiceImpl) Show(id string) (models.Post, error) {
	result, err := ps.model.FindById(ps.collection, id)
	if err != nil {
		return models.Post{}, err
	}
	return result, nil
}

func (ps *PostServiceImpl) Store(data validation.PostCreateForm) error {
	ps.model.Title = data.Title
	ps.model.Content = data.Content
	ps.model.CreatedAt = time.Now()
	ps.model.UpdatedAt = time.Now()
	err := ps.model.Create(ps.collection)
	if err != nil {
		return err
	}
	return nil
}
func (ps *PostServiceImpl) Update() error {
	return nil
}
func (ps *PostServiceImpl) Delete() error {
	return nil
}
