package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name   string
	TypeID uint
	Type   ProductType `gorm:"foreignKey:TypeID"`
	Price  float64
}
