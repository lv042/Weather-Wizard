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

func (f *fiberApp) ListAllHandlers() {
	f.fiberApp.GetRoutes(true)
}

func (f *fiberApp) setupRoutes() {
	// GET request to retrieve weather data by timestamp
	f.fiberApp.Get("/weather/:timestamp", func(c *fiber.Ctx) error {
		timestamp := c.Params("timestamp")

		// call GetWeatherDataByTimestampJSON method from dbManager object
		weatherData, err := dbManager.GetWeatherDataByTimestampJSON(timestamp)
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString(weatherData)
	})

	// GET request to retrieve all weather data
	f.fiberApp.Get("/weather", func(c *fiber.Ctx) error {
		// call GetAllWeatherDataJSON method from dbManager object
		weatherData, err := dbManager.GetAllWeatherDataJSON()
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.Send(weatherData)
	})

	// POST request to delete weather data by timestamp
	f.fiberApp.Post("/weather/delete", func(c *fiber.Ctx) error {
		// get JSON data from request body
		jsonStr := string(c.Body())

		// call DeleteWeatherDataJSON method from dbManager object
		result, err := dbManager.DeleteWeatherDataJSON(jsonStr)
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString(result)
	})

	// POST request to update weather data by timestamp
	f.fiberApp.Post("/weather/update", func(c *fiber.Ctx) error {
		// get JSON data from request body
		jsonStr := string(c.Body())

		// call UpdateWeatherDataJSON method from dbManager object
		result, err := dbManager.UpdateWeatherDataJSON(jsonStr)
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString(result)
	})

	// POST request to create weather data
	f.fiberApp.Post("/weather/create", func(c *fiber.Ctx) error {
		// get JSON data from request body
		jsonStr := string(c.Body())

		// call CreateWeatherDataJSON method from dbManager object
		result, err := dbManager.CreateWeatherDataJSON(jsonStr)
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString(result)
	})
}
