package application

import "github.com/thalysson-nascimento/voxify/internal/models"

type UserRepository interface {
	Save(user *models.UserModel) error
	FindAll() ([]models.UserModel, error)
}