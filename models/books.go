package models

import (
	"github.com/jinzhu/gorm"
)

type Books struct {
	ID        uint    `gorm:"primary key;autoincrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err.Error
}
