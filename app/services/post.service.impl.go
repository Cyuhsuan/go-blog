package services

import (
	"context"
	"go-blog/app/models"
	"go-blog/app/validation"

	"go.mongodb.org/mongo-driver/mongo"
)

type PostServiceImpl struct {
	model models.PostModel
}

func NewPostService(ctx context.Context, collection *mongo.Collection) PostService {
	return &PostServiceImpl{models.NewPostModel(collection, ctx)}
}

func (ps *PostServiceImpl) Index() ([]models.Post, error) {
	results, err := ps.model.FindAll()
	if err != nil {
		return []models.Post{}, err
	}
	return results, nil
}
func (ps *PostServiceImpl) Show(id string) (models.Post, error) {
	result, err := ps.model.FindById(id)
	if err != nil {
		return models.Post{}, err
	}
	return result, nil
}

func (ps *PostServiceImpl) Store(data validation.PostCreateForm) error {
	err := ps.model.Create(data)
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
