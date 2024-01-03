package routes

import (
	"github.com/gofiber/fiber/v2"
)

func User(router fiber.Router) {
	router.Post("/login", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	// router.Post("/register", controllers.Register)
}
