package middleware

import (
	"net/http"

	"github.com/adwinugroho/wedding-management-system/config"
	"github.com/adwinugroho/wedding-management-system/internals/logger"
	"github.com/adwinugroho/wedding-management-system/internals/models"
	"github.com/adwinugroho/wedding-management-system/modules/auth/services"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Validate, decode token, and save user info to context
func AuthenticationMiddleware(authService services.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
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
				return c.JSON(http.StatusUnauthorized, models.NewError("401-Unauthorized", "Unauthorized Token empty"))
			}

			/* TODO: checking if user already logout and use that token
			// Check if token is blacklisted
			// isBlacklisted, err := authService.IsTokenBlacklisted(c.Request().Context(), token)
			// if err != nil {
			// 	logger.LogError("Error checking token blacklist: " + err.Error())
			// 	return c.JSON(http.StatusInternalServerError, models.NewError("500-Internal-Error", "Internal server error"))
			// }

			// if isBlacklisted {
			// 	logger.LogError("Unauthorized cause token is blacklisted")
			// 	return c.JSON(http.StatusUnauthorized, models.NewError("401-Unauthorized", "Token has been invalidated"))
			// }
			*/

			// Parse and validate token
			claims, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
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
			c.Set("user_role", mapClaims["role"])

			return next(c)
		}
	}
}

func CheckRole() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
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
				return c.JSON(http.StatusUnauthorized, models.NewError("401-Unauthorized", "Unauthorized Token empty"))
			}
			claims, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
				return []byte(config.AppConfig.JWTSecret), nil
			})
			if err != nil {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}

			if claims == nil {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}

			mapClaims, ok := claims.Claims.(jwt.MapClaims)
			if !ok {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}

			role, ok := mapClaims["role"].(string)
			if !ok {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}

			if role != "ADMIN" && role != "USER" {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}
			return next(c)
		}
	}
}
