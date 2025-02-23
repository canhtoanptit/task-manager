package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"task-management-api/internal/config"
	"task-management-api/internal/models"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.AppConfig.DatabaseURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Task{})
	DB = db
	return db
}
