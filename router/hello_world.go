package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mochganjarn/try-github-action/handler"
)

func Router(app *fiber.App) {
	hello_route(app)
	app.Listen(":3000")
}

func hello_route(hello *fiber.App) {
	hello.Get("/", handler.CetakHello)
}
