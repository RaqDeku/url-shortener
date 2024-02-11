package handlers

import (
	"url-shorter/server/internals/store"
	"url-shorter/server/internals/util"

	"github.com/gofiber/fiber/v2"
)

type UrlShortenHandler struct {
	urlStore store.UrlStore
}

type UrlShortenHandlerParam struct {
	UrlStore store.UrlStore
}

// Recieve database as a dependency
func NewUrlShortenHandler(params UrlShortenHandlerParam) *UrlShortenHandler {
	return &UrlShortenHandler{
		urlStore: params.UrlStore,
	}
}

func (h *UrlShortenHandler) ShortenUrl(c *fiber.Ctx) error {
	longUrl := c.FormValue("longUrl")

	shortUrl := util.GenerateShortUrl(6)

	err := h.urlStore.StoreUrl(longUrl, shortUrl)

	if err != nil {
		return err
	}

	host := c.Hostname()

	return c.Status(201).Render("home", fiber.Map{
		"shortUrl": host + "/" + shortUrl,
	})
}
