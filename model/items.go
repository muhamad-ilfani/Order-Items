package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Item_code   string `json:"item_code" gorm:"primaryKey"`
	Description string `json:"description"`
	Quantity    int64  `json:"quantity"`
	ID_ref      uint
}
