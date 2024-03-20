package routes

import (
	"github.com/farhanaltariq/fiberplate/controllers"
	"github.com/farhanaltariq/fiberplate/middleware"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	services := middleware.InitServices()

	api := app.Group("/api")
	api.Get("/", controllers.NewMiscController(services).HealthCheck)

	Authentications(api.Group("/auth"), services)

	api.Use(middleware.AuthInterceptor)
	User(api.Group("/user"), services)
}
