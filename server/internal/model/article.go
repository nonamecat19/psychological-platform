package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title *string `gorm:"size:40"`
	Text  *string `gorm:"size:2048"`
}
