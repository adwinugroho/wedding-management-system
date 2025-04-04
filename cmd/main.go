package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/adwinugroho/wedding-management-system/config"
	"github.com/adwinugroho/wedding-management-system/internals/logger"
	"github.com/adwinugroho/wedding-management-system/modules/auth"
	"github.com/adwinugroho/wedding-management-system/modules/auth/handlers"
	"github.com/adwinugroho/wedding-management-system/modules/auth/repository"
	"github.com/adwinugroho/wedding-management-system/modules/auth/services"
	"github.com/labstack/echo/v4"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logger.InitLogger()

	config.LoadConfig()
}

func main() {
	logger.LogInfo("Starting application...")
	logger.LogInfo("Application started on port:" + config.AppConfig.Port)

	ctx := context.Background()

	dbPort, err := strconv.Atoi(config.PostgreSQLConfig.PostgreSQLPort)
	if err != nil {
		logger.LogFatal("Failed to convert port to int:" + err.Error())
	}
	dbHandler, err := config.InitConnectDB(
		ctx,
		config.PostgreSQLConfig.PostgreSQLHost,
		config.PostgreSQLConfig.PostgreSQLUser,
		config.PostgreSQLConfig.PostgreSQLPassword,
		config.PostgreSQLConfig.PostgreSQLDBName,
		int32(dbPort),
	)
	if err != nil {
		logger.LogFatal("Failed to connect to database:" + err.Error())
	}

	var e = echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	authRepository := repository.NewAuthRepository(dbHandler.DB)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)

	authGoogleHandler := handlers.NewAuthGoogleHandler(authService)

	auth.AuthRoutes(e, authHandler, authGoogleHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.AppConfig.Port)))
}
