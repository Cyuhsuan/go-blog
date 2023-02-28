package services

import (
	"go-blog/app/models"
	"go-blog/app/validation"
)

type AuthService interface {
	Regitster(validation.RegitsterForm) (*models.DBResponse, error)
	SignInUser(*models.SignInInput) (*models.DBResponse, error)
}
