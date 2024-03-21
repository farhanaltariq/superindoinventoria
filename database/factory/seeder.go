package factory

import "gorm.io/gorm"

// creata a function to call another function but using waitgroup

func Seed(db *gorm.DB) {
	seedUserAuth(db)
	seedProductType(db)
	SeedProduct(db)
}
