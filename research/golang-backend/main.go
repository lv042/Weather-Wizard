package main

import (
	"github.com/gofiber/fiber/v2"
)

var fiberApp *fiber.App

func main() {
	setup_db()
	fiberApp = fiber.New()

	// Define routes
	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fiberApp.Listen(":3000")

}

func initBackend() {

}

func initFiber() {

}
