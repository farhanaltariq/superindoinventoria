package services

import (
	"github.com/farhanaltariq/fiberplate/database/models"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductService interface {
	CreateOrUpdate(models.Product) error
	GetListProduct(c *fiber.Ctx, productTypeId int, searchKeyword, sortField string) ([]models.Product, error)
	DeleteProduct(int) error
	GetProductById(uint) (models.ProductResponse, error)
}

type productService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{db}
}

func (server *productService) CreateOrUpdate(product models.Product) error {
	query := server.db.Debug()
	err := query.Save(&product).Error
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

func (server *productService) GetListProduct(c *fiber.Ctx, productTypeId int, searchKeyword, sortField string) ([]models.Product, error) {
	var products []models.Product
	db := server.db

	// Menerapkan filter berdasarkan query parameter yang diberikan
	if productTypeId != 0 {
		db = db.Where("type_id = ?", productTypeId)
	}
	if searchKeyword != "" {
		db = db.Where("name ILIKE ?", "%"+searchKeyword+"%")
	}
	if sortField != "" {
		db = db.Order(sortField)
	}

	// Melakukan query ke database
	err := db.Debug().Find(&products).Error
	if err != nil {
		logrus.Errorln(err)
		return products, err
	}

	return products, nil
}

func (server *productService) DeleteProduct(id int) error {
	query := server.db.Debug()
	err := query.Delete(&models.Product{}, id).Error
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	return nil
}

func (server *productService) GetProductById(id uint) (models.ProductResponse, error) {
	var product models.Product
	// preload("Type")
	query := server.db.Debug().Preload("Type").Where("id = ?", id)
	err := query.First(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.ProductResponse{}, nil
		}
		logrus.Errorln(err)
		return models.ProductResponse{}, err
	}

	return models.ProductResponse{Products: []models.Product{product}}, nil
}
