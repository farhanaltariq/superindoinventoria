package controllers

import (
	"github.com/farhanaltariq/fiberplate/common/codes"
	"github.com/farhanaltariq/fiberplate/common/status"
	"github.com/gofiber/fiber/v2"
)

type MiscController interface {
	HealthCheck(c *fiber.Ctx) error
}

type miscController struct {
}

func NewMiscController() MiscController {
	return &miscController{}
}
func (server *miscController) HealthCheck(c *fiber.Ctx) error {
	return status.Successf(c, codes.OK, "Server Running")
}
