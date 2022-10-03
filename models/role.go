package models

import "time"

type Role struct {
	Id          uint         `json:"id" gorm:"primary_key"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permission"`
	Users       []User       `json:"users" gorm:"many2many:role_user"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"autoUpdateTime:true"`
}
