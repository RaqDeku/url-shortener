package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"url-shorter/server/internals/handlers"
	"url-shorter/server/internals/store/dbstore"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	// Initialize template engine.
	engine := handlebars.New("./views", ".hbs")

	// Initialise fiber app an set template engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(helmet.New())
	app.Use(logger.New())

	// Initialize database
	urlStore := dbstore.NewUrlStore()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// Passed database as a dependency into both post and get handlers
	app.Post("/", handlers.NewUrlShortenHandler(handlers.UrlShortenHandlerParam{
		UrlStore: urlStore,
	}).ShortenUrl)

	app.Get("/:shortUrl", handlers.GetOrignalUrlHandler(handlers.GetUrlHandlerParams{
		UrlStore: urlStore,
	}).GetUrl)

	// Kill signals
	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-killSig
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	// serve app
	if err := app.Listen(":8080"); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
	// TODO: run cleanup tasks

}
