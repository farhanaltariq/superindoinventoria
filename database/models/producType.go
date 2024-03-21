package models

import "gorm.io/gorm"

type ProductType struct {
	gorm.Model
	Name string
}

type ProductTypeRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductTypeResponse struct {
	Name string
}
