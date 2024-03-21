package models

import "gorm.io/gorm"

type ProductType struct {
	gorm.Model
	Name string
}

type ProductTypeResponse struct {
	Name string
}
