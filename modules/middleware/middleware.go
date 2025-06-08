package middleware

import (
	"net/http"

	"github.com/adwinugroho/wedding-management-system/config"
	"github.com/adwinugroho/wedding-management-system/internals/logger"
	"github.com/adwinugroho/wedding-management-system/internals/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Validate, decode token, and save user info to context
func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get token from cookie
		cookie, err := c.Cookie("token")
		if err != nil {
			logger.LogError("Unauthorized cause token cookie not found")
			return c.JSON(http.StatusUnauthorized, models.NewError("401-Unauthorized", "Unauthorized"))
		}

		token := cookie.Value
		if token == "" {
			logger.LogError("Unauthorized cause token is empty")
			return c.JSON(http.StatusUnauthorized, models.NewError("401-Unauthorized", "Unauthorized"))
		}

		// Parse and validate token
		claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JWTSecret), nil
		})

		if err != nil {
			logger.LogError("Unauthorized cause token validation failed: " + err.Error())
			return c.JSON(http.StatusUnauthorized, models.NewError("401-Unauthorized", "Unauthorized"))
		}

		if !claims.Valid {
			logger.LogError("Unauthorized cause token is invalid")
			return c.JSON(http.StatusUnauthorized, models.NewError("401-Unauthorized", "Unauthorized"))
		}

		// Get claims
		mapClaims, ok := claims.Claims.(jwt.MapClaims)
		if !ok {
			logger.LogError("Unauthorized cause invalid token claims")
			return c.JSON(http.StatusUnauthorized, models.NewError("401-Unauthorized", "Unauthorized"))
		}

		// Set user info in context
		c.Set("user_id", mapClaims["id"])
		c.Set("user_email", mapClaims["email"])
		c.Set("user_name", mapClaims["name"])
		c.Set("user_role", mapClaims["role"])

		return next(c)
	}
}
