package model

import "gorm.io/gorm"

type TherapyGroup struct {
	gorm.Model
	Name         string `gorm:"size:100"`
	Description  string `gorm:"type:text"`
	Users        []User `gorm:"many2many:group_users"`
	SpecialistID uint
	Specialist   User `gorm:"foreignKey:SpecialistID"`
	Messages     []Message
}
