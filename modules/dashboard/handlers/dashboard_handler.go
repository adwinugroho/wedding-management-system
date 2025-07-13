package handlers

import (
	"fmt"
	"net/http"

	"github.com/adwinugroho/wedding-management-system/config"
	"github.com/adwinugroho/wedding-management-system/internals/logger"
	"github.com/adwinugroho/wedding-management-system/modules/auth/services"
	"github.com/labstack/echo/v4"
)

type DashboardHandler interface {
	GetDashboard(c echo.Context) error
	GetEvents(c echo.Context) error
}

type dashboardHandler struct {
	authService services.AuthService
}

func NewDashboardHandler(authService services.AuthService) DashboardHandler {
	return &dashboardHandler{authService: authService}
}

func (h *dashboardHandler) GetDashboard(c echo.Context) error {
	role, ok := c.Get("user_role").(string)
	if !ok {
		return c.String(http.StatusForbidden, "Error, invalid user")
	}
	err := c.Render(http.StatusOK, "dashboard.html", map[string]any{
		"staticPath":     "/static",
		"baseURL":        fmt.Sprintf("%s:%s/user/dashboard", config.AppConfig.AppURL, config.AppConfig.Port),
		"frontURL":       fmt.Sprintf("%s:%s", config.AppConfig.AppURL, config.AppConfig.Port),
		"title":          "Dashboard | Wedding Planner",
		"page":           "dashboard",
		"role":           role,
		"hrefDashboard":  "text-pink-500 hover:text-pink-600",
		"ihrefDashboard": "opacity-75",
		"hrefEvent":      "text-blueGray-700 hover:text-blueGray-500",
		"ihrefEvent":     "text-blueGray-300",
	})
	if err != nil {
		logger.LogError("[GetDashboard] Error while rendering page, cause:" + err.Error())
		return c.String(http.StatusForbidden, "Error, page not found")
	}
	return nil
}

func (h *dashboardHandler) GetEvents(c echo.Context) error {
	role, ok := c.Get("user_role").(string)
	if !ok {
		return c.String(http.StatusForbidden, "Error, invalid user")
	}
	err := c.Render(http.StatusOK, "events.html", map[string]any{
		"staticPath":     "/static",
		"baseURL":        fmt.Sprintf("%s:%s/user/dashboard", config.AppConfig.AppURL, config.AppConfig.Port),
		"frontURL":       fmt.Sprintf("%s:%s", config.AppConfig.AppURL, config.AppConfig.Port),
		"title":          "Events | Wedding Planner",
		"page":           "Events",
		"role":           role,
		"hrefEvent":      "text-pink-500 hover:text-pink-600",
		"ihrefEvent":     "opacity-75",
		"hrefDashboard":  "text-blueGray-700 hover:text-blueGray-500",
		"ihrefDashboard": "text-blueGray-300",
	})
	if err != nil {
		logger.LogError("[GetEvents] Error while rendering page, cause:" + err.Error())
		return c.String(http.StatusForbidden, "Error, page not found")
	}

	return nil
}

// func (h *dashboardHandler) GetVenues(c echo.Context) error {
// 	err := c.Render(http.StatusOK, "events.html", map[string]any{
// 		"staticPath":     "/static",
// 		"baseURL":        fmt.Sprintf("%s:%s/user/dashboard", config.AppConfig.AppURL, config.AppConfig.Port),
// 		"frontURL":       fmt.Sprintf("%s:%s", config.AppConfig.AppURL, config.AppConfig.Port),
// 		"title":          "Venues | Wedding Planner",
// 		"page":           "Venues",
// 		"hrefDashboard":  "text-blueGray-700 hover:text-blueGray-500",
// 		"ihrefDashboard": "text-blueGray-300",
// 		"hrefEvent":      "text-blueGray-700 hover:text-blueGray-500",
// 		"ihrefEvent":     "text-blueGray-300",
// 		"hrefVenue":      "text-pink-500 hover:text-pink-600",
// 		"ihrefVenue":     "opacity-75",
// 	})
// 	if err != nil {
// 		logger.LogError("[GetVenues] Error while rendering page, cause:" + err.Error())
// 		return c.String(http.StatusForbidden, "Error, page not found")
// 	}

// 	return nil
// }
