package services

import (
	"context"
	"go-blog/app/models"
	"go-blog/app/validation"

	"go.mongodb.org/mongo-driver/mongo"
)

type PostServiceImpl struct {
	model *models.PostModel
}

var ctx context.Context

func NewPostService(collection *mongo.Collection) PostService {
	return &PostServiceImpl{models.NewPostModel(collection)}
}

func (ps *PostServiceImpl) Index() ([]models.Post, error) {
	results, err := ps.model.FindAll(ctx)
	if err != nil {
		return []models.Post{}, err
	}
	return results, nil
}
func (ps *PostServiceImpl) Show(id string) (models.Post, error) {
	result, err := ps.model.FindById(ctx, id)
	if err != nil {
		return models.Post{}, err
	}
	return result, nil
}

func (ps *PostServiceImpl) Store(data validation.PostCreateForm) error {
	err := ps.model.Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
func (ps *PostServiceImpl) Update(data validation.PostCreateForm, id string) error {
	err := ps.model.UpdateById(ctx, data, id)
	if err != nil {
		return err
	}
	return nil
}
func (ps *PostServiceImpl) Delete(id string) error {
	err := ps.model.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
