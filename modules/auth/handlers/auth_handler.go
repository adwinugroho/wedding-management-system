package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adwinugroho/wedding-management-system/config"
	"github.com/adwinugroho/wedding-management-system/internals/logger"
	"github.com/adwinugroho/wedding-management-system/internals/models"
	"github.com/adwinugroho/wedding-management-system/modules/auth/services"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler interface {
	Login(c echo.Context) error
	GetLogin(c echo.Context) error
	GetRegister(c echo.Context) error
	Register(c echo.Context) error
	Logout(c echo.Context) error
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

	user, err := h.authService.GetUserByEmail(c.Request().Context(), email)
	if err != nil {
		logger.LogError("Error while get user by email, cause: " + err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error, Please Contact Customer Service.")
	}

	if user == nil {
		return c.String(http.StatusUnauthorized, "Wrong email or password. Please try again.")
	}

	if user.Password == nil {
		return c.String(http.StatusUnauthorized, "Wrong email or password. Please try again.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		logger.LogError("Error while compare hash and password, cause: " + err.Error())
		return c.String(http.StatusUnauthorized, "Wrong email or password. Please try again.")
	}

	user, token, err := h.authService.GenerateJWTToken(c.Request().Context(), *user)
	if err != nil {
		logger.LogError("Error while generate JWT token, cause: " + err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error, Please Contact Customer Service.")
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

func (h *authHandler) Register(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	if name == "" || email == "" || password == "" {
		return c.String(http.StatusBadRequest, "Name, email, and password are required.")
	}

	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: &password,
	}

	user, err := h.authService.RegisterUser(c.Request().Context(), newUser)
	if err != nil {
		logger.LogError("Error while registering user, cause: " + err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error, Please Contact Customer Service.")
	}

	// Generate JWT token for the new user
	user, token, err := h.authService.GenerateJWTToken(c.Request().Context(), *user)
	if err != nil {
		logger.LogError("Error while generate JWT token, cause: " + err.Error())
		return c.String(http.StatusInternalServerError, "Internal Server Error, Please Contact Customer Service.")
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

	c.Response().Header().Set("HX-Redirect", "/user/dashboard")
	return c.NoContent(http.StatusOK)
}

func (h *authHandler) Logout(c echo.Context) error {
	// Remove JWT token cookie
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	c.SetCookie(cookie)

	c.Response().Header().Set("HX-Redirect", "/auth/login")
	return c.NoContent(http.StatusOK)
}
