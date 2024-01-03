package routes

import (
	"github.com/farhanaltariq/fiberplate/controllers"
	"github.com/farhanaltariq/fiberplate/database"
	"github.com/farhanaltariq/fiberplate/middleware"
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	// add /api prefix to all routes

	api := app.Group("/api")

	api.Get("/", controllers.NewMiscController().HealthCheck)

	db := database.GetDBConnection()
	Authentications(api.Group("/auth"), db)

	userGroup := api.Group("/user/:id")
	userGroup.Use(middleware.AuthInterceptor) // Use JwtMiddleware for the "/user" route group
	User(userGroup)
}
