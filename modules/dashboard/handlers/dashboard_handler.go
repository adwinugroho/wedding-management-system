package handlers

import (
	"fmt"
	"net/http"

	"github.com/adwinugroho/wedding-management-system/config"
	"github.com/adwinugroho/wedding-management-system/modules/auth/services"
	"github.com/labstack/echo/v4"
)

type DashboardHandler interface {
	GetDashboard(c echo.Context) error
}

type dashboardHandler struct {
	authService services.AuthService
}

func NewDashboardHandler(authService services.AuthService) DashboardHandler {
	return &dashboardHandler{authService: authService}
}

func (h *dashboardHandler) GetDashboard(c echo.Context) error {
	return c.Render(http.StatusOK, "dashboard.html", map[string]any{
		"StaticPath": "/static",
		"BaseURL":    fmt.Sprintf("%s:%s/user/dashboard", config.AppConfig.AppURL, config.AppConfig.Port),
	})
}
