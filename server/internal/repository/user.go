package repository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return data, err
	}

	data.Password = string(hashedPassword)
	data.Role = "user"
	err = r.db.Create(&data).Error

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

func (r *UserRepository) FindByCredentials(email string, password string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

type MessageRepository struct {
	db *gorm.DB
}
