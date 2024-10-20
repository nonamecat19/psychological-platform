package repository

import (
	"gorm.io/gorm"
	"server/internal/database"
	"server/internal/model"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{
		db: database.GetPostgresConnection(),
	}
}

func (r *MessageRepository) Create(data model.Message) (model.Message, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *MessageRepository) FindAll() ([]model.Message, error) {
	var messages []model.Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *MessageRepository) GetOne(id uint) (model.Message, error) {
	var message model.Message
	err := r.db.First(&message, id).Error
	return message, err
}

func (r *MessageRepository) UpdateOne(id uint, data model.Message) (model.Message, error) {
	var message model.Message
	err := r.db.First(&message, id).Error
	if err != nil {
		return message, err
	}
	err = r.db.Model(&message).Updates(data).Error
	return message, err
}

func (r *MessageRepository) DeleteOne(id uint) error {
	return r.db.Delete(&model.Message{}, id).Error
}
