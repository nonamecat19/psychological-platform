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

func (r *UserRepository) GetOne(id uint) (model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *UserRepository) UpdateOne(data model.User) (model.User, error) {
	var user model.User
	err := r.db.First(&user, data.ID).Error
	if err != nil {
		return user, err
	}
	err = r.db.Model(&user).Updates(data).Error
	return user, err
}

func (r *UserRepository) DeleteOne(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

type MessageRepository struct {
	db *gorm.DB
}
