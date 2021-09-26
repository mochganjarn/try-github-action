package handler

import (
	"github.com/gofiber/fiber/v2"
)

func CetakHello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
