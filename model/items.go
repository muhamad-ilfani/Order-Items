package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Item_code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_id    string `gorm:"primary_key" json:"order_id"`
}
