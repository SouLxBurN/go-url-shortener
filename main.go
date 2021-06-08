package main

import (
	"url-shortener/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	route.Configure(app)
	app.Listen(":3000")
}
