package handlers

import (
	"url-shorter/server/internals/store"

	"github.com/gofiber/fiber/v2"
)

type GetUrlHandler struct {
	urlStore store.UrlStore
}

type GetUrlHandlerParams struct {
	UrlStore store.UrlStore
}

// Recieve Database as dependency
func GetOrignalUrlHandler(params GetUrlHandlerParams) *GetUrlHandler {
	return &GetUrlHandler{
		urlStore: params.UrlStore,
	}
}

func (h *GetUrlHandler) GetUrl(c *fiber.Ctx) error {
	shortUrl := c.Params("shortUrl")

	url, err := h.urlStore.GetUrl(shortUrl)
	if err != nil {
		return err
	}

	return c.Status(302).Redirect(url.LongUrl)
}
