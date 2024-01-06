package controllers

import (
	"github.com/farhanaltariq/fiberplate/common/codes"
	"github.com/farhanaltariq/fiberplate/common/status"
	"github.com/farhanaltariq/fiberplate/middleware"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetListUser(c *fiber.Ctx) error
}

type userController struct {
	middleware.Services
}

func NewUserController(service middleware.Services) UserController {
	return &userController{service}
}

func (s *userController) GetListUser(c *fiber.Ctx) error {
	return status.Successf(c, codes.OK, "OK")
}
