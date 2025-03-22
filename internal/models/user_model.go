package models

import (
	"time"

	"github.com/google/uuid"
)

// Representação do banco de dados (ORM)
type UserModel struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;not null"`
	CreatedAt time.Time `gorm:"default:now()"`
	UpdatedAt time.Time `gorm:"default:now()"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
}
// Define o nome correto da tabela
func (UserModel) TableName() string {
	return "users"
}