package repository

import (
	"gorm.io/gorm"
	"server/internal/database"
	"server/internal/model"
)

type TherapyGroupRepository struct {
	db *gorm.DB
}

func NewTherapyGroupRepository() *TherapyGroupRepository {
	return &TherapyGroupRepository{
		db: database.GetPostgresConnection(),
	}
}

func (r *TherapyGroupRepository) Create(data model.TherapyGroup) (model.TherapyGroup, error) {
	err := r.db.Create(&data).Error
	return data, err
}

func (r *TherapyGroupRepository) FindAll() ([]model.TherapyGroup, error) {
	var groups []model.TherapyGroup
	err := r.db.Find(&groups).Error
	return groups, err
}

func (r *TherapyGroupRepository) FindAllBySpecialist(id uint) ([]model.TherapyGroup, error) {
	var groups []model.TherapyGroup
	err := r.db.Where("specialist_id = ?", id).Preload("Specialist").Find(&groups).Error
	return groups, err
}

func (r *TherapyGroupRepository) FindById(id uint) ([]model.TherapyGroup, error) {
	var users []model.User
	err := r.db.Where("is_anonymous = FALSE").Find(&users).Error
	if err != nil {
		return nil, err
	}

	var groups []model.TherapyGroup
	err = r.db.Preload("Messages").Preload("Specialist").Where("id = ?", id).Find(&groups).Error
	if err != nil {
		return nil, err
	}

	for _, group := range groups {
		userMap := make(map[uint]model.User)
		for _, user := range users {
			userMap[user.ID] = user
		}

		for i, message := range group.Messages {
			if user, exists := userMap[message.UserID]; exists {
				group.Messages[i].User = &user
			}
		}
	}

	return groups, nil
}

func (r *TherapyGroupRepository) GetOne(id uint) (model.TherapyGroup, error) {
	var group model.TherapyGroup
	err := r.db.First(&group, id).Error
	return group, err
}

func (r *TherapyGroupRepository) UpdateOne(data model.TherapyGroup) (model.TherapyGroup, error) {
	var group model.TherapyGroup
	err := r.db.First(&group, data.ID).Error
	if err != nil {
		return group, err
	}
	err = r.db.Model(&group).Updates(data).Error
	return group, err
}

func (r *TherapyGroupRepository) DeleteOne(id uint) error {
	return r.db.Delete(&model.TherapyGroup{}, id).Error
}
