package main

import (
	"os"
	"url-shortener/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	route.Configure(app)

	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = ":3000"
	}
	app.Listen(port)
}
