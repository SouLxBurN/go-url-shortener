package route

import (
	"url-shortener/handler"

	"github.com/gofiber/fiber/v2"
)

// Configure
// Configures all routes for the *fiber.App
func Configure(app *fiber.App) {

	// Create new url and return hash.
	app.Post("/new", handler.CreateHandler)

	// Handle Redirects for short url
	app.Get("/:urlHash", handler.ResolveHandler)

}
