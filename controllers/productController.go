package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/farhanaltariq/fiberplate/common/codes"
	"github.com/farhanaltariq/fiberplate/common/status"
	"github.com/farhanaltariq/fiberplate/database/models"
	"github.com/farhanaltariq/fiberplate/middleware"
	"github.com/farhanaltariq/fiberplate/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductController interface {
	CreateOrUpdataProduct(c *fiber.Ctx) error
	GetListProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
	GetProductById(c *fiber.Ctx) error
}

func NewProductController(service middleware.Services) ProductController {
	return &controller{service}
}

// paginate performs pagination on a given GORM database connection

// @Summary Create Or Update Product
// @Description Create Or Update Product
// @Tags Product
// @Accept json
// @Security Authorization
// @Param data body models.ProductRequest true "Product data"
// @Produce json
// @Success 200 {object} models.ProductResponseExample
// @Failure 400 {object} common.ResponseMessage
// @Router /product [post]
func (s *controller) CreateOrUpdataProduct(c *fiber.Ctx) error {
	product := models.ProductRequest{}
	if err := json.Unmarshal(c.Body(), &product); err != nil {
		return status.Errorf(c, codes.InternalServerError, err.Error())
	}

	productOrm := models.Product{
		Model: gorm.Model{
			ID: uint(product.ID),
		},
		Name:   product.Name,
		TypeID: product.TypeID,
		Price:  product.Price,
	}

	err := s.Services.ProductService.CreateOrUpdate(productOrm)
	if err != nil {
		return status.Errorf(c, codes.InternalServerError, err.Error())
	}
	return status.Successf(c, codes.OK, "OK")
}

// @Summary Get Product
// @Description Get List Product
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} models.ProductResponseExample
// @Failure 400 {object} common.ResponseMessage
// @Router /product [get]
func (s *controller) GetListProduct(c *fiber.Ctx) error {
	res := models.ProductResponse{}
	pagination := utils.SetPagination(c)

	data, err := s.Services.ProductService.GetListProduct(c)
	if err != nil {
		return status.Errorf(c, codes.InternalServerError, err.Error())
	}

	res.Products = data
	res.Pagination = pagination
	return c.Status(codes.OK).JSON(res)
}

// @Summary Delete Product
// @Description Soft delete Product data
// @Tags Product
// @Security Authorization
// @Param id path int true "Product ID"
// @Produce json
// @Success 200 {object} common.ResponseMessage
// @Failure 400 {object} common.ResponseMessage
// @Router /product/{id} [delete]
func (s *controller) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return status.Errorf(c, codes.InternalServerError, err.Error())
	}

	err = s.Services.ProductService.DeleteProduct(id)
	if err != nil {
		return status.Errorf(c, codes.InternalServerError, err.Error())
	}
	return status.Successf(c, codes.OK, "OK")
}

// @Summary Get Product Details
// @Description Get Product By Id
// @Tags Product
// @Param id path int true "Product ID"
// @Produce json
// @Success 200 {object} models.ProductResponseExample
// @Failure 400 {object} common.ResponseMessage
// @Router /product/{id} [get]
func (s *controller) GetProductById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return status.Errorf(c, codes.InternalServerError, err.Error())
	}
	res, err := s.Services.ProductService.GetProductById(uint(id))
	if err != nil {
		return status.Errorf(c, codes.InternalServerError, err.Error())
	}
	return c.Status(codes.OK).JSON(res)
}
