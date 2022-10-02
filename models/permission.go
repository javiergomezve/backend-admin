package models

import "gorm.io/gorm"

type Permission struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Roles []Role `json:"roles" gorm:"many2many:role_permission"`

	gorm.Model
}
