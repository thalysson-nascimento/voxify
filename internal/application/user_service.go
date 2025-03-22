package application

import (
	"log"

	"github.com/jinzhu/copier"
	"github.com/thalysson-nascimento/voxify/internal/core"
	"github.com/thalysson-nascimento/voxify/internal/models"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(coreUser core.User) (*models.UserModel, error) {
	// Converte core.User para models.UserModel
	var modelUser models.UserModel
	err := copier.Copy(&modelUser, &coreUser)
	if err != nil {
		log.Fatalf("Erro ao copiar: %v", err)
		return nil, err
	}

	// Salva o modelo no banco de dados (agora passa o ponteiro)
	if err := s.repo.Save(&modelUser); err != nil {
		log.Printf("Erro ao salvar o usuário: %v", err)
		return nil, err
	}

	// Retorna o modelo do usuário criado, que já tem o ID, CreatedAt e UpdatedAt preenchidos
	return &modelUser, nil
}


func (service *UserService) GetAllUsers() ([]models.UserModel, error) {
	return service.repo.FindAll()
}
