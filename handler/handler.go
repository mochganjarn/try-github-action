package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Person struct {
	Sammy UserID `json:"Sammy"`
	Jesse UserID `json:"Jesse"`
	Drew  UserID `json:"Drew"`
	Jamie UserID `json:"Jamie"`
}

type UserID struct {
	Username  string `json:"Username"`
	Followers int    `json:"Followers"`
}

type FollowerResult struct {
	Followers int
}

func Follower(c *fiber.Ctx) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://jsonkeeper.com/b/DMXK", nil)
	if err != nil {
		return c.SendString(err.Error())
	}

	resp, err := client.Do(req)
	if err != nil {
		return c.SendString(err.Error())
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var p Person
	err = decoder.Decode(&p)
	if err != nil {
		return c.SendString(err.Error())
	}

	switch c.Params("userid") {
	case "SammyShark":
		f := FollowerResult{p.Sammy.Followers}
		return c.JSON(f)
	case "JesseOctopus":
		f := FollowerResult{p.Jesse.Followers}
		return c.JSON(f)
	case "DrewSquid":
		f := FollowerResult{p.Drew.Followers}
		return c.JSON(f)
	case "JamieMantisShrimp":
		f := FollowerResult{p.Jamie.Followers}
		return c.JSON(f)
	default:
		return fiber.ErrNotFound
	}
}

func Detail(c *fiber.Ctx) error {
	resp, err := http.Get("https://jsonkeeper.com/b/DMXK")
	if err != nil {
		return c.SendString(err.Error())
	}

	decoder := json.NewDecoder(resp.Body)
	var p Person
	err = decoder.Decode(&p)
	if err != nil {
		return c.SendString(err.Error())
	}

	switch c.Params("userid") {
	case "sammy":
		return c.JSON(p.Sammy)
	case "jesse":
		return c.JSON(p.Jesse)
	case "drew":
		return c.JSON(p.Drew)
	case "jamie":
		return c.JSON(p.Jamie)
	default:
		return fiber.ErrNotFound
	}
}

func CetakHello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
