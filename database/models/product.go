package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name   string
	TypeID uint
	Type   ProductType `gorm:"foreignKey:TypeID"`
	Price  float64
}

type ProductRequest struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	TypeID uint    `json:"typeId"`
	Price  float64 `json:"price"`
}

type ProductData struct {
	Name   string
	TypeID uint
	Type   ProductTypeResponse
	Price  float64
}
type ProductResponseExample struct {
	Products   []ProductData
	Pagination Pagination
}

type ProductResponse struct {
	Products   []Product
	Pagination Pagination
}
