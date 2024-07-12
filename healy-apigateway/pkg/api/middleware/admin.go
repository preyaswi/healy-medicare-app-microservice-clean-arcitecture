package middleware

import (
	"fmt"
	"healy-apigateway/pkg/api/response"
	"healy-apigateway/pkg/helper"
	"strings"

	"github.com/gofiber/fiber/v2"

	"net/http"
)

func AdminAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenHeader := c.Get("Authorization")
		fmt.Println(tokenHeader, "this is the token header")

		if tokenHeader == "" {
			response := response.ClientResponse("No auth header provided", nil, nil)
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response := response.ClientResponse("Invalid Token Format", nil, nil)
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		tokenpart := splitted[1]
		tokenClaims, err := helper.ValidateToken(tokenpart)
		if err != nil {
			response := response.ClientResponse("Invalid Token", nil, err.Error())
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		c.Locals("tokenClaims", tokenClaims)

		return c.Next()
	}
}
