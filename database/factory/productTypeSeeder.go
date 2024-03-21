package factory

import (
	"github.com/farhanaltariq/fiberplate/database/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func seedProductType(db *gorm.DB) {
	// Define product types to be seeded
	productTypes := []models.ProductType{
		{
			Name: "Sayuran",
		},
		{
			Name: "Protein",
		},
		{
			Name: "Buah",
		},
		{
			Name: "Snack",
		},
	}

	// Check if product types already exist in the database
	var existingTypes []models.ProductType
	result := db.Find(&existingTypes).Debug()
	if result.Error != nil {
		// Handle error if any
		// For simplicity, you can log the error
		logrus.Println("Error:", result.Error)
		return
	}

	// If no product types found, seed the data
	if len(existingTypes) == 0 {
		// Seed product types
		db.Debug().Create(&productTypes)
	}
}
