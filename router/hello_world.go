package router

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/mochganjarn/try-github-action/handler"
)

func Router(app *fiber.App) {
	err := godotenv.Load()
	if os.Getenv("APP_ENV") != "production" {
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	port := os.Getenv("APP_PORT")
	hello_route(app)
	app.Listen(":" + port)
}

func hello_route(hello *fiber.App) {
	hello.Get("/", handler.CetakHello)
}
