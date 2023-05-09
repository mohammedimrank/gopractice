package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

var userMap map[string]int

func main() {
	app := fiber.New()

	userMap = make(map[string]int)

	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("Date:", time.Now())

		return c.Next()
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!!")
	})

	app.Get("/query", func(c *fiber.Ctx) error {
		fmt.Println(string(c.Request().URI().QueryString()))
		return c.SendString("hello query string")
	})

	app.Get("/param/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		fmt.Println("Id recieved with param is:", id)
		return c.SendString("bye!!")
	})

	app.Get("/user/:name", func(c *fiber.Ctx) error {
		return getCounter(c)
	})

	log.Fatal(app.Listen(":3000"))
}

func getCounter(c *fiber.Ctx) error {

	count := 1
	username := c.Params("name")
	if user, ok := userMap[username]; ok {
		count = user
	}
	userMap[username] = count + 1

	return c.Status(200).JSON(&fiber.Map{
		"user":  username,
		"count": count,
	})
}
