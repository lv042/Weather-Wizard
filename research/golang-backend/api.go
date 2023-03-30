package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/url"
)

type FiberApp struct {
	fiberApp *fiber.App
}

func (f *FiberApp) Log(message string) {
	log.Default().Println("FiberApp: ", message)
}

func (f *FiberApp) LogError(message string) {
	red := color.New(color.FgRed).SprintFunc()
	log.Fatal("FiberApp Error: ", red(message))
}

func (f *FiberApp) GetInfo() {
	f.Log(fmt.Sprintf("%+v", f))
}

// NewFiberApp constructor for fiber
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

func (f *FiberApp) logMiddleware(c *fiber.Ctx) error {
	// Log request method and URL
	f.Log(fmt.Sprintf("Request: %s %s", c.Method(), c.Path()))

	// Log request headers
	headers, _ := json.Marshal(c.Request())
	f.Log(fmt.Sprintf("Request Headers: %s", headers))

	// Log request body
	body := c.Request().Body()
	f.Log(fmt.Sprintf("Request Body: %s", body))

	// Restore the request body for further processing
	c.Request().SetBody(body)

	// Continue processing the request
	err := c.Next()

	// Log response status code
	f.Log(fmt.Sprintf("Response: %d", c.Response().StatusCode()))

	// Log response headers
	headers, _ = json.Marshal(c.Response().Header)
	f.Log(fmt.Sprintf("Response Headers: %s", headers))

	// Log response body
	body = c.Response().Body()
	f.Log(fmt.Sprintf("Response Body: %s \n\n\n", body))

	// Restore the response body for further processing
	c.Response().SetBody(body)

	return err
}

func (f *FiberApp) setupRoutes() {

	// add middleware to Log input and output for all routes
	f.fiberApp.Use(f.logMiddleware)

	// GET request to retrieve weather data by timestamp
	f.fiberApp.Get("/weather/:timestamp", func(c *fiber.Ctx) error {
		timestamp := c.Params("timestamp")

		// URL-decode the timestamp
		decodedTimestamp, err := url.QueryUnescape(timestamp)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid timestamp format")
		}

		// call GetWeatherDataByTimestampJSON method from dbManager object
		weatherData, err := dbManager.GetWeatherDataByTimestampJSON(decodedTimestamp)
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
