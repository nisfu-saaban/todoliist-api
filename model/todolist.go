package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name string `gorm:"column:name"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Todo{})
}
