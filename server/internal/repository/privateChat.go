package repository

import (
	"gorm.io/gorm"
	"server/internal/database"
	"server/internal/model"
)

type PrivateChatRepository struct {
	db *gorm.DB
}

func NewPrivateChatRepository() *PrivateChatRepository {
	return &PrivateChatRepository{
		db: database.GetPostgresConnection(),
	}
}

func (r *PrivateChatRepository) Create(data model.PrivateChat) (model.PrivateChat, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *PrivateChatRepository) FindAll() ([]model.PrivateChat, error) {
	var chats []model.PrivateChat
	err := r.db.Find(&chats).Error
	return chats, err
}

func (r *PrivateChatRepository) GetOne(id uint) (model.PrivateChat, error) {
	var chat model.PrivateChat
	err := r.db.First(&chat, id).Error
	return chat, err
}

func (r *PrivateChatRepository) UpdateOne(data model.PrivateChat) (model.PrivateChat, error) {
	var chat model.PrivateChat
	err := r.db.First(&chat, data.ID).Error
	if err != nil {
		return chat, err
	}
	err = r.db.Model(&chat).Updates(data).Error
	return chat, err
}

func (r *PrivateChatRepository) DeleteOne(id uint) error {
	return r.db.Delete(&model.PrivateChat{}, id).Error
}