package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mochganjarn/try-github-action/router"
)

func main() {
	app := fiber.New()

	router.Router(app)
}
