package database

import (
	"gorm.io/gorm"
	"server/internal/model"
)

func migrateDb(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.PrivateChat{},
		&model.Message{},
		&model.TherapyGroup{},
	)
}
