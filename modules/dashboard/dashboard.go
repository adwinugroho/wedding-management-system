package dashboard

import (
	"github.com/adwinugroho/wedding-management-system/modules/auth/services"
	"github.com/adwinugroho/wedding-management-system/modules/dashboard/handlers"
	"github.com/adwinugroho/wedding-management-system/modules/middleware"
	"github.com/labstack/echo/v4"
)

func DashboardRoutes(e *echo.Echo, dashboardHandler handlers.DashboardHandler, authService services.AuthService) {
	// Apply auth middleware to all dashboard routes
	dashboardGroup := e.Group("/user/dashboard")
	dashboardGroup.Use(middleware.AuthenticationMiddleware(authService))
	dashboardGroup.Use(middleware.CheckRole())

	dashboardGroup.GET("", dashboardHandler.GetDashboard)
	dashboardGroup.GET("/events", dashboardHandler.GetEvents)
	// dashboardGroup.POST("/events", dashboardHandler.CreateEvent)
	// dashboardGroup.GET("/guests", dashboardHandler.GetGuests)
	// Add more dashboard-related routes here
}
