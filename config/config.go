package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func InitializeConfig() {
	db = initializeDB()
}

func GetDB() *gorm.DB {
	return db
}
