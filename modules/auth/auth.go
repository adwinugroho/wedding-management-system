package auth

import (
	"github.com/adwinugroho/wedding-management-system/modules/auth/handlers"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, authHandler handlers.AuthHandler, authGoogleHandler handlers.AuthGoogleHandler) {
	authGroup := e.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	// authGroup.POST("/logout", authHandler.Logout)
	// authGroup.POST("/register", authHandler.Register)

	authGoogle := e.Group("/auth/google")
	authGoogle.GET("/login", authGoogleHandler.LoginWithGoogle)
	authGoogle.GET("/callback", authGoogleHandler.GoogleCallback)
}
