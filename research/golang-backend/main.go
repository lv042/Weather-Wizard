package main

import (
	"github.com/gofiber/fiber/v2"
)

var fiberApp *fiber.App

func main() {
	initBackend()
	fiberApp = fiber.New()

	// Define routes
	fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	fiberApp.Listen(":3000")

}

func initBackend() {
	//new db manager object
	dbManager := NewDBManager("Postgres")
	print(dbManager.ToString())

	initFiber()
}

func initFiber() {

}
