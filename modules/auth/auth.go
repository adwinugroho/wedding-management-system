package auth

import (
	"github.com/adwinugroho/wedding-management-system/modules/auth/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthRoutes(e *echo.Echo, authHandler handlers.AuthHandler, authGoogleHandler handlers.AuthGoogleHandler) {
	authGroup := e.Group("/auth")
	authGroup.POST("/login", authHandler.Login)
	authGroup.GET("/login", authHandler.GetLogin)
	authGroup.POST("/logout", authHandler.Logout)
	authGroup.POST("/register", authHandler.Register, middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(5)))
	authGroup.GET("/register", authHandler.GetRegister)

	authGoogle := e.Group("/auth/google")
	authGoogle.GET("/login", authGoogleHandler.LoginWithGoogle)
	authGoogle.GET("/callback", authGoogleHandler.GoogleCallback)
}
