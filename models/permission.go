package models

import "time"

type Permission struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Roles     []Role    `json:"roles" gorm:"many2many:role_permission"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}
