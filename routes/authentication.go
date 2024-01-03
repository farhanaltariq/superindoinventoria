package routes

import (
	"github.com/farhanaltariq/fiberplate/controllers"
	"github.com/farhanaltariq/fiberplate/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Authentications(router fiber.Router, db *gorm.DB) {
	authService := services.NewAuthService(db)
	userService := services.NewUserService(db)
	authController := controllers.NewAuthController(authService, userService)

	router.Post("/login", authController.Login)
	router.Post("/register", authController.Register)
}
