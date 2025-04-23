package handlers

import (
	"fmt"
	"net/http"

	"github.com/adwinugroho/wedding-management-system/config"
	"github.com/adwinugroho/wedding-management-system/modules/auth/services"
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Login(c echo.Context) error
	GetLogin(c echo.Context) error
	GetRegister(c echo.Context) error
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{authService: authService}
}

func (h *authHandler) GetLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", map[string]any{
		"staticPath": "/static",
		"baseURL":    fmt.Sprintf("%s:%s", config.AppConfig.AppURL, config.AppConfig.Port),
	})
}

func (h *authHandler) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		// Return HTML response for HTMX
		// return c.HTML(http.StatusBadRequest, `
		// 	<div id="errorMessage" class="text-red-500 text-center text-sm mt-2">
		// 		Email and password are required
		// 	</div>
		// `)

		return c.String(http.StatusBadRequest, "Email and password are required.")
	}

	// user, err := h.authService.Login(c.Request().Context(), email, password)
	// if err != nil {
	// 	// Return HTML response for HTMX
	// 	// return c.HTML(http.StatusUnauthorized, `
	// 	// 	<div id="errorMessage" class="text-red-500 text-center text-sm mt-2">
	// 	// 		`+err.(*models.JsonResponse).Message+`
	// 	// 	</div>
	// 	// `)
	// 	return c.String(http.StatusUnauthorized, "Wrong email or password. Please try again.")
	// }

	// user.Password = nil

	// For successful login, return a redirect response
	c.Response().Header().Set("HX-Redirect", "/user/dashboard")
	return c.NoContent(http.StatusOK)
}

func (h *authHandler) GetRegister(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", map[string]any{
		"staticPath": "/static",
		"baseURL":    fmt.Sprintf("%s:%s", config.AppConfig.AppURL, config.AppConfig.Port),
	})
}
