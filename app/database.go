package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func ConnectionDatabase() *gorm.DB {
	dns := "root:tryHackme@tcp(localhost:3306)/todolist?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.DB()
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}
