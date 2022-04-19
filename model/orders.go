package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Customer_name string    `json:"customer_name"`
	Order_id      uint      `json:"order_id"`
	Ordered_at    time.Time `json:"ordered_at"`
	Item_user     []Item    `json:"item_user" gorm:"foreignKey:ID_ref;references:ID"`
}
