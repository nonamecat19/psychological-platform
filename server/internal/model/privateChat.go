package model

import "gorm.io/gorm"

type PrivateChat struct {
	gorm.Model
	SpecialistID uint
	Specialist   User `gorm:"foreignKey:SpecialistID"`
	ClientID     uint
	Client       User `gorm:"foreignKey:ClientID"`
	Messages     []Message
}
