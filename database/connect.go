package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_admin"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}
}
