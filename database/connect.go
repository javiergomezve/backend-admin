package database

import (
	"github.com/javiergomezve/backend-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_admin"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	database.AutoMigrate(&models.User{})
}
