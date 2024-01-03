package middleware

import (
	"time"

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
	token, _ := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	// Check if the token is valid
	if !token.Valid {
		return status.Errorf(c, codes.Unauthorized, "Invalid token")
	}

	// If all checks pass, set the user claims in locals
	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		return status.Errorf(c, codes.Unauthorized, "Invalid token")
	}

	logrus.Infoln("Time Now: ", time.Now().Format("2006-01-02 15:04:05"))
	logrus.Infoln("Expired At : ", claims.ExpiresAt)
	return c.Next()
}
