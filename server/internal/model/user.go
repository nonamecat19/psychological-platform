package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          *string `gorm:"size:40"`
	Surname       *string `gorm:"size:40"`
	IsAnonymous   bool    `gorm:"default:false"`
	Role          string  `gorm:"size:50"`
	Email         string  `gorm:"size:255"`
	Password      string  `gorm:"size:255"`
	Messages      []Message
	TherapyGroups []TherapyGroup `gorm:"many2many:therapy_groups"`
}
