package main

import (
	"log"

	"github.com/adwinugroho/wedding-management-system/config"
	"github.com/adwinugroho/wedding-management-system/internals/logger"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	logger := logger.InitLogger()

	config.LoadConfig()

	logger.Info("Starting application...")
	logger.Info("Application started on port:", config.AppConfig.Port)
}
