package repository

import (
	"server/internal/database"
	"server/internal/model"
)

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

func (r *MessageRepository) UpdateOne(data model.Message) (model.Message, error) {
	var message model.Message
	err := r.db.First(&message, data.ID).Error
	if err != nil {
		return message, err
	}
	err = r.db.Model(&message).Updates(data).Error
	return message, err
}

func (r *MessageRepository) DeleteOne(id uint) error {
	return r.db.Delete(&model.Message{}, id).Error
}
