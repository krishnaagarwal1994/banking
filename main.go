package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	// Here we are defining the routes

	// log.Print("Starting our Application")
	logger.Info("Starting our Application")
	app.Start()
}
