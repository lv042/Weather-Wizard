package main

import (
	"github.com/gofiber/fiber/v2"
)

type fiberApp struct {
	fiberApp *fiber.App
}

//constructor for fiber

func NewFiberApp() *fiberApp {
	return &fiberApp{fiberApp: fiber.New()}
}

func main() {
	initBackend()
}

func initBackend() {
	//new db manager object
	dbManager := NewDBManager("Postgres")
	dbManager.GetInfo()
	dbManager.setupDb()
	result, error := dbManager.getWeatherData()

	//new fiber app object
	fiberApp := NewFiberApp()
	fiberApp.InitFiber()
}

func (f *fiberApp) Listen(address string) {
	f.fiberApp.Listen(address)
}

func (f *fiberApp) InitFiber() {
	f.setupRoutes()
	f.Listen(":3000")
}

func (f *fiberApp) setupRoutes() {
	//Get requests
	f.fiberApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}
