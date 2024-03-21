package controllers

import (
	"encoding/json"

	"github.com/farhanaltariq/fiberplate/common/codes"
	"github.com/farhanaltariq/fiberplate/common/status"
	"github.com/farhanaltariq/fiberplate/database/models"
	"github.com/farhanaltariq/fiberplate/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductTypeController interface {
	CreateOrUpdateType(c *fiber.Ctx) error
}

func NewProductTypeController(service middleware.Services) ProductTypeController {
	return &controller{service}
}

// @Summary Create Or Update Product Type
// @Description Create Or Update Product Type Data
// @Tags Product Type
// @Accept json
// @Security Authorization
// @Param data body models.ProductTypeRequest true "Product Type data"
// @Produce json
// @Success 200 {object} models.ProductResponseExample
// @Failure 400 {object} common.ResponseMessage
// @Router /type [post]
func (s *controller) CreateOrUpdateType(c *fiber.Ctx) error {
	logrus.Infoln("Dijalan yang benar")
	product := models.ProductTypeRequest{}
	if err := json.Unmarshal(c.Body(), &product); err != nil {
		return status.Errorf(c, codes.InternalServerError, err.Error())
	}

	productTypeOrm := models.ProductType{
		Model: gorm.Model{
			ID: uint(product.ID),
		},
		Name: product.Name,
	}

	logrus.Infoln(productTypeOrm)

	err := s.Services.ProductTypeService.InsertOrUpdate(productTypeOrm)
	if err != nil {
		return status.Errorf(c, codes.InternalServerError, err.Error())
	}
	return status.Successf(c, codes.OK, "OK")
}
