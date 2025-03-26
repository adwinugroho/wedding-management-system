package middleware

import (
	"net/http"

	"github.com/adwinugroho/simple-wedding-management/internals/logger"
	"github.com/labstack/echo/v4"
)

// Validate, decode token, and save user info to context
func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			logger.LogError("Unauthorized cause authorization header is empty")
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		// TODO: Implement jwt decode and validation
		return next(c)
	}
}
