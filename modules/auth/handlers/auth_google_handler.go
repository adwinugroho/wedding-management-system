package handlers

import (
	"net/http"

	"github.com/adwinugroho/wedding-management-system/internals/models"
	"github.com/adwinugroho/wedding-management-system/internals/sso"
	"github.com/adwinugroho/wedding-management-system/modules/auth/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AuthGoogleHandler interface {
	LoginWithGoogle(c echo.Context) error
	GoogleCallback(c echo.Context) error
}

type authGoogleHandler struct {
	authService services.AuthService
}

func NewAuthGoogleHandler(authService services.AuthService) AuthGoogleHandler {
	return &authGoogleHandler{authService: authService}
}

func (h *authGoogleHandler) LoginWithGoogle(c echo.Context) error {
	googleOAuthURL := sso.GetGoogleOAuthURL(uuid.New().String()) // uuid as state
	return c.Redirect(http.StatusTemporaryRedirect, googleOAuthURL)
}

func (h *authGoogleHandler) GoogleCallback(c echo.Context) error {
	if c.QueryParam("state") == "" {
		return c.JSON(http.StatusBadRequest, models.NewError("403-Invalid-Google-Callback", "Invalid state"))
	}

	code := c.QueryParam("code")
	token, err := sso.GetGoogleOAuthToken(code, c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewError("500-Google-OAuth-Token-Error", err.Error()))
	}

	userInfo, err := sso.GetGoogleOAuthUserInfo(token, c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewError("500-Google-OAuth-User-Info-Error", err.Error()))
	}

	user, err := h.authService.GetUserByEmail(c.Request().Context(), userInfo.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewError("500-Get-User-By-Email-Error", err.Error()))
	}

	return c.JSON(http.StatusOK, models.NewJsonResponse(true).SetData(user))
}
