package router

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mochganjarn/try-github-action/handler"
)

func Router(app *fiber.App) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	port := os.Getenv("PORT")
	hello_route(app)
	app.Listen(":" + port)
}

func hello_route(hello *fiber.App) {
	hello.Get("/", handler.CetakHello)
}
