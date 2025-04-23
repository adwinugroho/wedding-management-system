package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/adwinugroho/wedding-management-system/config"
	"github.com/adwinugroho/wedding-management-system/internals/logger"
	routeAuth "github.com/adwinugroho/wedding-management-system/modules/auth"
	handlerAuth "github.com/adwinugroho/wedding-management-system/modules/auth/handlers"
	repoAuth "github.com/adwinugroho/wedding-management-system/modules/auth/repository"
	serviceAuth "github.com/adwinugroho/wedding-management-system/modules/auth/services"
	routeDashboard "github.com/adwinugroho/wedding-management-system/modules/dashboard"
	handlerDashboard "github.com/adwinugroho/wedding-management-system/modules/dashboard/handlers"
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
	logger.LogInfo("Application URL:" + config.AppConfig.AppURL)

	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 60*time.Second)
	defer cancel()

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

	// Add template renderer
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// Add security headers middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Security headers
			c.Response().Header().Set("X-Content-Type-Options", "nosniff")
			c.Response().Header().Set("X-Frame-Options", "DENY")
			c.Response().Header().Set("X-XSS-Protection", "1; mode=block")
			c.Response().Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
			c.Response().Header().Set("Content-Security-Policy", "default-src 'self' https: 'unsafe-inline'")

			// Cache control for static files
			if strings.HasPrefix(c.Path(), "/static") {
				c.Response().Header().Set("Cache-Control", "public, max-age=31536000")
			}

			return next(c)
		}
	})

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/", serveHTML)

	authRepository := repoAuth.NewAuthRepository(dbHandler.DB)
	authService := serviceAuth.NewAuthService(authRepository)
	authHandler := handlerAuth.NewAuthHandler(authService)
	dashboardHandler := handlerDashboard.NewDashboardHandler(authService)

	authGoogleHandler := handlerAuth.NewAuthGoogleHandler(authService)

	routeAuth.AuthRoutes(e, authHandler, authGoogleHandler)
	routeDashboard.DashboardRoutes(e, dashboardHandler)

	// Serve static files with proper MIME types
	e.Static("/static", "static")
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Set proper content type for CSS files
			if strings.HasSuffix(c.Path(), ".css") {
				c.Response().Header().Set(echo.HeaderContentType, "text/css")
			}
			// Set proper content type for JavaScript files
			if strings.HasSuffix(c.Path(), ".js") {
				c.Response().Header().Set(echo.HeaderContentType, "application/javascript")
			}
			return next(c)
		}
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.AppConfig.Port)))
}

func serveHTML(c echo.Context) error {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error loading template")
	}

	// Render index.html explicitly
	return tmpl.ExecuteTemplate(c.Response().Writer, "index.html", map[string]any{
		"staticPath": "/static",
		"baseURL":    fmt.Sprintf("%s:%s", config.AppConfig.AppURL, config.AppConfig.Port),
	})
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
