package main

import (
	"url-shorter/server/internals/handlers"
	"url-shorter/server/internals/store/dbstore"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars/v2"
)

func main() {
	// Initialize template engine.
	engine := handlebars.New("./views", ".hbs")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

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

	app.Listen(":8080")
}
