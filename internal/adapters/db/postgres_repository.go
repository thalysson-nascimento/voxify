package db

import (
	"log"

	"github.com/thalysson-nascimento/voxify/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository() (*PostgresRepository, error) {
	dsn := "host=localhost user=thalysson password=151087 dbname=go_hex_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrar a estrutura do banco
	err = db.AutoMigrate(&models.UserModel{})
	if err != nil {
		log.Fatalf("Erro ao migrar banco de dados: %v", err)
		return nil, err
	}

	log.Println("Migração concluída com sucesso!")

	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Save(user *models.UserModel) error {
	// Cria o usuário no banco e preenche o ID automaticamente
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	// Aqui, o 'user' já terá o ID gerado pelo banco
	return nil
}

func (r *PostgresRepository) FindAll() ([]models.UserModel, error) {
	var users []models.UserModel
	err := r.db.Find(&users).Error
	return users, err
}
