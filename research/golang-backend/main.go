package main

import (
	"github.com/gofiber/fiber/v2"
)

type fiberApp struct {
	fiberApp *fiber.App
}

var dbManager *DBManager

// constructor for fiber
func NewFiberApp() *fiberApp {
	return &fiberApp{fiberApp: fiber.New()}
}

func main() {
	initBackend()
}

func initBackend() {
	//new db manager object
	dbManager = NewDBManager("test")
	dbManager.GetInfo()
	dbManager.setupDb()

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
	weatherDataHandler := func(c *fiber.Ctx) error {
		// Read all weather data
		// Get the weather data
		data, err := dbManager.GetWeatherData()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error: " + err.Error())
		}
		if data == "" {
			return c.SendString("No weather data found")
		}
		return c.SendString(data)
	}

	// Register weather data route
	f.fiberApp.Get("/weather-data", weatherDataHandler)
}
