package main

import "github.com/gofiber/fiber/v2"

func main() {
	// create fiber instance
	server := fiber.New()

	// dummy response for /
	server.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, world",
		})
	})

	// start rest server
	server.Listen(":8000")
}
