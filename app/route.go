package app

import (
	"os"
	"os/signal"
	"time"

	"crud/handler"
	"crud/health"
	"crud/model"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"

	log "crud/logger"
)

func Fiber(appInfo model.AppInfo, config model.FiberConfig) *fiber.App {
	fiberConfig := fiber.Config{
		ErrorHandler:          handler.ErrorResponseHandler,
		ReadTimeout:           time.Millisecond * time.Duration(config.ReadTimeout),
		WriteTimeout:          time.Millisecond * time.Duration(config.WriteTimeout),
		IdleTimeout:           time.Millisecond * time.Duration(config.IdleTimeout),
		AppName:               appInfo.Name,
		DisableStartupMessage: !appInfo.Banner,
	}

	fiber := fiber.New(fiberConfig)
	withGracefulShutdown(fiber)

	fiber.Get("/health", adaptor.HTTPHandlerFunc(health.HealthCheckHandler()))

	person := fiber.Group("/person")
	person.Get("/:id", handler.GetPerson)
	person.Get("/", handler.GetAllPerson)
	person.Post("/", handler.CreatePerson)
	person.Put("/:id", handler.UpdatePerson)
	person.Delete("/:id", handler.DeletePerson)
	return fiber
}

func withGracefulShutdown(fiber *fiber.App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		var _ = <-c
		log.Logger().Info("Tearing down fiber...")
		_ = fiber.Shutdown()
	}()
}
