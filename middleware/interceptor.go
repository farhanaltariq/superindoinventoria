package middleware

import (
	"github.com/farhanaltariq/fiberplate/common/codes"
	"github.com/farhanaltariq/fiberplate/common/status"
	"github.com/farhanaltariq/fiberplate/database/models"
	"github.com/farhanaltariq/fiberplate/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func CommonMiddleware(c *fiber.Ctx) error {
	endpoint := c.BaseURL() + c.Path()
	logrus.Infoln(utils.FormatMethod(c), endpoint)
	return c.Next()
}

func AuthInterceptor(c *fiber.Ctx) error {
	jwtSecret := []byte(utils.GetEnv("JWT_SECRET", "secret"))

	// Check if Authorization header is present
	if c.Get("Authorization") == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Extract the token from the Authorization header
	tokenString := c.Get("Authorization")[7:] // Remove "Bearer " prefix

	// Parse the token with custom claims and key
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return status.Errorf(c, codes.Unauthorized, "Invalid token")
	}

	// Check if the token is valid
	if !token.Valid {
		return status.Errorf(c, codes.Unauthorized, "Invalid token")
	}

	// If all checks pass, set the user claims in locals
	_, ok := token.Claims.(*models.Claims)
	if !ok {
		return status.Errorf(c, codes.Unauthorized, "Invalid token")
	}

	return c.Next()
}
