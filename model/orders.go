package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Customer_name string    `json:"customer_name"`
	Ordered_at    time.Time `json:"ordered_at"`
	Item_user     Item      `gorm:"foreignKey:Order_id;reference:ID" json:"items_user"`
}
