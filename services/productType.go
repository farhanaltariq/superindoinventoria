package services

import (
	"github.com/farhanaltariq/fiberplate/database/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductTypeService interface {
	InsertOrUpdate(models.ProductType) error
}

type productTypeService struct {
	db *gorm.DB
}

func NewProductTypeService(db *gorm.DB) ProductTypeService {
	return &productTypeService{db}
}

func (server *productTypeService) InsertOrUpdate(productType models.ProductType) error {
	logrus.Infoln("InsertOrUpdate: ", productType)
	query := server.db.Debug()
	err := query.Save(&productType).Error
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}
