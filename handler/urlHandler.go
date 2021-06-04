package handler

import (
	"log"
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

	urlHash, err := service.CreateShortUrl(sURL)
	if err != nil {
		return err
	}

	// Also formulate a damn response
	return c.JSON(fiber.Map{
		"urlHash": urlHash,
	})
}

// ResolveHandler Resolves the hash and redirects to the url.
func ResolveHandler(c *fiber.Ctx) error {
	urlHash := c.Params("urlHash")
	log.Println(urlHash)

	redirectUrl, err := service.GetShortUrl(urlHash)
	if err != nil {
		return err
	}

	return c.Redirect(redirectUrl)
}
