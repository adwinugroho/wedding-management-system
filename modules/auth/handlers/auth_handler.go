package handlers

import (
	"net/http"

	"github.com/adwinugroho/wedding-management-system/internals/models"
	"github.com/adwinugroho/wedding-management-system/modules/auth/services"
	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	Login(c echo.Context) error
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{authService: authService}
}

func (h *authHandler) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.JSON(http.StatusBadRequest, models.NewJsonResponse(false).SetError("400-Bad-Request", "Email and password are required"))
	}

	user, err := h.authService.Login(c.Request().Context(), email, password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.(*models.JsonResponse))
	}

	user.Password = nil

	return c.JSON(http.StatusOK, models.NewJsonResponse(true).SetData(user).SetMessage("Login successful"))
}
