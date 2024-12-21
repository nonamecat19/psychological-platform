package model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID         uint
	Content        string `gorm:"type:text"`
	PrivateChatID  *uint
	TherapyGroupID *uint
	User           *User `gorm:"-"`
}
