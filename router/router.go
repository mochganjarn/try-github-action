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
	if os.Getenv("APP_ENV") == "development" {
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")
	hello_route(app)
	user_route(app)
	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}

func hello_route(hello *fiber.App) {
	hello.Get("/", handler.CetakHello)
}

func user_route(app *fiber.App) {
	app.Get("/follower/:userid", handler.Follower)
	app.Get("/:userid/detail", handler.Detail)
}
