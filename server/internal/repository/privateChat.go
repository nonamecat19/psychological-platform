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

func (r *PrivateChatRepository) FindById(id uint) ([]model.PrivateChat, error) {
	var chats []model.PrivateChat
	err := r.db.Preload("Messages").Preload("Specialist").Preload("Client").Where("id = ?", id).Find(&chats).Error
	return chats, err
}

func (r *PrivateChatRepository) FindAllByClient(id uint) ([]model.PrivateChat, error) {
	var chats []model.PrivateChat
	err := r.db.Where("client_id = ?", id).Preload("Specialist").Find(&chats).Error
	return chats, err
}

func (r *PrivateChatRepository) FindAllBySpecialist(id uint) ([]model.PrivateChat, error) {
	var chats []model.PrivateChat
	err := r.db.Where("specialist_id = ?", id).
		Preload("Client", func(db *gorm.DB) *gorm.DB {
			return db.Where("\"users\".\"is_anonymous\" = FALSE")
		}).
		Find(&chats).Error
	if err != nil {
		return nil, err
	}
	return chats, nil
}

func (r *PrivateChatRepository) FindConcreteChat(clientID, specialistID uint) (model.PrivateChat, error) {
	var chat model.PrivateChat
	err := r.db.Where("client_id = ? AND specialist_id = ?", clientID, specialistID).First(&chat).Error
	return chat, err
}

func (r *PrivateChatRepository) CreateConcreteChat(clientID, specialistID uint) (model.PrivateChat, error) {
	chat := model.PrivateChat{
		ClientID:     clientID,
		SpecialistID: specialistID,
	}
	err := r.db.Create(&chat).Error
	return chat, err
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
