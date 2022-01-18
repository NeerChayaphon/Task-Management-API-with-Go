package database

import (
	"github.com/NeerChayaphon/Task-Management-API-with-Go/model"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		return err
	}
	return nil
}
