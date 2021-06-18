package handler

import (
	"fmt"
	"url-shortener/config"
	"url-shortener/model"
	"url-shortener/service"

	"github.com/gofiber/fiber/v2"
)

// CreateHandler Creates and stores a hash alongside the provided
// url and returns said hash for future referencing.
func CreateHandler(c *fiber.Ctx) error {
	sURL := new(model.ShortURL)
	if err := c.BodyParser(sURL); err != nil {
		return err
	}

	err := service.CreateShortUrl(sURL)
	if err != nil {
		return err
	}

	shortUrl := fmt.Sprintf("%s/%s", config.GetConfig().Domain, sURL.Hash)
	return c.JSON(fiber.Map{
		"shortUrl": shortUrl,
	})
}

// ResolveHandler Resolves the hash and redirects to the url.
func ResolveHandler(c *fiber.Ctx) error {
	urlHash := c.Params("urlHash")

	shortURL, err := service.GetShortUrl(urlHash)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if shortURL == nil {
		return fiber.NewError(404, "Short URL Not Found")
	}

	return c.Redirect(shortURL.URL)
}
