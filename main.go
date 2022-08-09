package main

import (
	"github.com/dannielss/banking/app"
	"github.com/dannielss/banking/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
