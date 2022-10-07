package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	Id         uint        `json:"id"`
	FirstName  string      `json:"-"`
	LastName   string      `json:"-"`
	Name       string      `json:"name" gorm:"-"`
	Email      string      `json:"email"`
	Total      float32     `json:"total" gorm:"-"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime:true"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Id           uint      `json:"id"`
	OrderId      uint      `json:"order_id"`
	ProductTitle string    `json:"product_title"`
	Price        float32   `json:"price"`
	Quantity     uint      `json:"quantity"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}

func (order *Order) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Order{}).Count(&total)

	return total
}

func (order *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var orders []Order

	db.Debug().Offset(offset).Limit(limit).Preload("OrderItems").Find(&orders)

	for i := range orders {
		var total float32 = 0

		for _, orderItem := range orders[i].OrderItems {
			total += orderItem.Price * float32(orderItem.Quantity)
		}

		orders[i].Name = orders[i].FirstName + " " + orders[i].LastName
		orders[i].Total = total
	}

	return orders
}
