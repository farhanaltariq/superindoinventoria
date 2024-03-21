package routes

import (
	"github.com/farhanaltariq/fiberplate/controllers"
	"github.com/farhanaltariq/fiberplate/middleware"
	"github.com/gofiber/fiber/v2"
)

func Product(router fiber.Router, service middleware.Services) {
	productController := controllers.NewProductController(service)

	router.Get("/", productController.GetListProduct)
	router.Get("/:id", productController.GetProductById)

	router.Use(middleware.AuthInterceptor)
	router.Post("/", productController.CreateOrUpdataProduct)
	router.Delete("/:id", productController.DeleteProduct)

}
