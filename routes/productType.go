package routes

import (
	"github.com/farhanaltariq/fiberplate/controllers"
	"github.com/farhanaltariq/fiberplate/middleware"
	"github.com/gofiber/fiber/v2"
)

func ProductType(router fiber.Router, service middleware.Services) {
	productTypeController := controllers.NewProductTypeController(service)

	router.Use(middleware.AuthInterceptor)
	router.Post("/", productTypeController.CreateOrUpdateType)
}
