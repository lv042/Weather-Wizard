package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type FiberApp struct {
	fiberApp *fiber.App
}

// constructor for fiber
func NewFiberApp() *FiberApp {
	return &FiberApp{fiberApp: fiber.New()}
}

func (f *FiberApp) Listen(address string) {
	f.fiberApp.Listen(address)
}

func (f *FiberApp) InitFiber() {
	f.setupRoutes()
	f.Listen(":3000")
}

func (f *FiberApp) ListAllHandlers() {
	f.fiberApp.GetRoutes(true)
}

func logMiddleware(f *fiber.Ctx) error {
	// log input
	fmt.Println("Input:", string(f.Body()))

	// continue processing the request
	err := f.Next()

	// log output
	fmt.Println("Output:", string(f.Response().Body()))

	return err
}

func (f *FiberApp) setupRoutes() {

	// add middleware to log input and output for all routes
	f.fiberApp.Use(logMiddleware)

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
	f.fiberApp.Delete("/weather/delete", func(c *fiber.Ctx) error {
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
	f.fiberApp.Put("/weather/update", func(c *fiber.Ctx) error {
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
