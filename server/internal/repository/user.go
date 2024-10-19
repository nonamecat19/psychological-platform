package repository

import (
	"gorm.io/gorm"
	"server/internal/database"
	"server/internal/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetPostgresConnection(),
	}
}

func (r *UserRepository) Create(data model.User) (model.User, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}
