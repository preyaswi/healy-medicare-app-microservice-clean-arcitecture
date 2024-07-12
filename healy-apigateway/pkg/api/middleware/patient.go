package middleware

import (
	"fmt"
	"healy-apigateway/pkg/helper"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func UserAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		tokenString := helper.GetTokenFromHeader(authHeader)
		if tokenString == "" {
			tokenString = c.Cookies("Authorization")
			if tokenString == "" {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"error":   "Unauthorized",
					"message": "Authorization token not provided",
				})
			}
		}
		userID, userEmail, err := helper.ExtractUserIDFromToken(tokenString)
		fmt.Println(userID, "userid", userEmail, "useremail")
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error":   "Unauthorized",
				"message": "Invalid token",
			})
		}
		c.Locals("user_id", userID)
		c.Locals("user_email", userEmail)
		c.Locals("patient_id", userID)
		c.Locals("doctor_id", userID)
		fmt.Println(userID, "userid")
		return c.Next()
	}
}
