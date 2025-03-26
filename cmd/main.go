package main

import (
	"log"

	"github.com/adwinugroho/simple-wedding-management/internals/logger"
)

func main() {
	log.Println("simple wedding management")
	logger := logger.InitLogger()
	logger.Info("Starting application...")
}
