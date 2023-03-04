package services

import (
	"go-blog/app/models"
	"go-blog/app/validation"
)

type PostService interface {
	Index() ([]models.Post, error)
	Show(id string) (models.Post, error)
	Store(data validation.PostCreateForm) error
	Update() error
	Delete() error
}
