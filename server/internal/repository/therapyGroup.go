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
