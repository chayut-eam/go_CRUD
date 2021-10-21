package main

import (
	"fmt"

	"crud/app"
	log "crud/logger"
)

func main() {
	config, err := app.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error loading configuration: %w \n", err))
	}

	app.Bootstrap(config)
	defer app.TearDown()

	log.Logger().Infof("Starting %v...", config.AppInfo.Name)

	fiber := app.Fiber(config.AppInfo, config.FiberConfig)
	if err = fiber.Listen(config.FiberConfig.Address); err != nil {
		log.Logger().Fatalf("Error starting %v. Caused by: %v.", config.AppInfo.Name, err)
	}
}