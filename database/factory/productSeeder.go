package factory

import (
	"math/rand"

	"github.com/bxcodec/faker/v3"
	"github.com/farhanaltariq/fiberplate/database/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func getProductTypeID(db *gorm.DB) (typeIds []uint) {
	var productType models.ProductType
	db.Model(&productType).Select("id").Find(&typeIds)
	return typeIds
}

func SeedProduct(db *gorm.DB) {

	// do nothing if data is already over 100
	var count int64
	sql := db.Model(&models.Product{}).Count(&count)
	if sql.Error != nil {
		logrus.Errorln(sql.Error)
		return
	}

	if count >= 100 {
		logrus.Infoln("Data is already over 100, skipping product seeding")
		return
	}

	typeID := getProductTypeID(db)
	products := []models.Product{}

	// implement faker to generate random food
	// for example

	for i := 0; i < 100; i++ {
		product := models.Product{
			Name:  faker.Word(),
			Price: rand.Float64() * 1000,
			// get random typeid
			TypeID: typeID[rand.Intn(len(typeID))],
		}
		products = append(products, product)
	}

	for _, product := range products {
		db.Create(&product)
	}

}
