package dashboard

import (
	"net/http"
	"strings"

	"github.com/adwinugroho/wedding-management-system/modules/dashboard/handlers"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func DashboardRoutes(e *echo.Echo, dashboardHandler handlers.DashboardHandler) {
	// Apply auth middleware to all dashboard routes
	dashboardGroup := e.Group("/user/dashboard")
	// dashboardGroup.Use(CheckRole())

	dashboardGroup.GET("", dashboardHandler.GetDashboard)
	dashboardGroup.GET("/events", dashboardHandler.GetEvents)
	// dashboardGroup.POST("/events", dashboardHandler.CreateEvent)
	// dashboardGroup.GET("/guests", dashboardHandler.GetGuests)
	// Add more dashboard-related routes here
}

func CheckRole() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorizationHeader := c.Request().Header.Get("Authorization")
			if authorizationHeader == "" {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}
			token := strings.Split(authorizationHeader, " ")[1]
			claims, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
				return []byte("your-secret-key"), nil
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
