package models

import "gorm.io/gorm"

type Role struct {
	Id          uint         `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permission"`
	Users       []User       `json:"users" gorm:"many2many:role_user"`

	gorm.Model
}
