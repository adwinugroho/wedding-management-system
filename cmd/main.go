package main

import (
	"log"

	"github.com/adwinugroho/wedding-management-system/internals/logger"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	logger := logger.InitLogger()
	logger.Info("Starting application...")
}
