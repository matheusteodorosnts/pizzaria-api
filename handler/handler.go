package handler

import (
	"github.com/matheusteodorosnts/pizzaria-api/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandler() {
	db = config.GetDB()
}
