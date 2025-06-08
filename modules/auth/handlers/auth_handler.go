package handlers

import (
	"fmt"
	"net/http"
	"time"

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
		return c.String(http.StatusBadRequest, "Email and password are required.")
	}

	user, token, err := h.authService.Login(c.Request().Context(), email, password)
	if err != nil {
		return c.String(http.StatusUnauthorized, "Wrong email or password. Please try again.")
	}

	user.Password = nil

	// Set JWT token in cookie
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	c.SetCookie(cookie)

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
