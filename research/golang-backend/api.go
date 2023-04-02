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
	metrics  *Metrics
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
	return &FiberApp{
		fiberApp: fiber.New(),
		metrics:  NewMetrics(),
	}
}

func (f *FiberApp) Listen(address string) {
	f.fiberApp.Listen(address)
}

func (f *FiberApp) InitFiber() {
	f.setupRoutes()
	f.Listen(":3001")
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

	// add middleware to log input and output for all routes
	f.fiberApp.Use(f.logMiddleware)

	// Serve static files from the "static" directory
	f.fiberApp.Static("/", "./web")

	// GET request to retrieve weather data by timestamp
	f.fiberApp.Get("api/weather/:timestamp", func(c *fiber.Ctx) error {
		timestamp := c.Params("timestamp")

		// URL-decode the timestamp
		decodedTimestamp, err := url.QueryUnescape(timestamp)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Invalid timestamp format")
		}

		// call GetWeatherDataByTimestampJSON method from dbManager object
		weatherData, err := dbManager.GetWeatherDataByTimestampJSON(decodedTimestamp)
		if err != nil {
			f.metrics.IncrementErrorCount(c.Route().Path)
			return c.SendString(err.Error())
		}

		f.metrics.IncrementRequestCount(c.Route().Path)
		return c.SendString(weatherData)
	})

	// GET request to retrieve all weather data
	f.fiberApp.Get("api/weather", func(c *fiber.Ctx) error {
		// call GetAllWeatherDataJSON method from dbManager object
		weatherData, err := dbManager.GetAllWeatherDataJSON()
		if err != nil {
			f.metrics.IncrementErrorCount(c.Route().Path)
			return c.SendString(err.Error())
		}

		f.metrics.IncrementRequestCount(c.Route().Path)
		return c.Send(weatherData)
	})

	// POST request to delete weather data by timestamp
	f.fiberApp.Delete("api/weather/delete", func(c *fiber.Ctx) error {
		// get JSON data from request body
		jsonStr := string(c.Body())

		// call DeleteWeatherDataJSON method from dbManager object
		result, err := dbManager.DeleteWeatherDataJSON(jsonStr)
		if err != nil {
			f.metrics.IncrementErrorCount(c.Route().Path)
			fmt.Println("Failed to delete weather data:", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		if result == "No weather data found for the specified timestamp" {
			f.metrics.IncrementErrorCount(c.Route().Path)
			return c.SendStatus(fiber.StatusNotFound)
		}

		f.metrics.IncrementRequestCount(c.Route().Path)
		return c.SendString("Weather data deleted successfully.")
	})

	// POST request to update weather data by timestamp
	f.fiberApp.Put("api/weather/update", func(c *fiber.Ctx) error {
		// get JSON data from request body
		jsonStr := string(c.Body())

		// call UpdateWeatherDataJSON method from dbManager object
		result, err := dbManager.UpdateWeatherDataJSON(jsonStr)
		if err != nil {
			f.metrics.IncrementErrorCount(c.Route().Path)
			return c.SendString(err.Error())
		}

		f.metrics.IncrementRequestCount(c.Route().Path)
		return c.SendString(result)
	})

	// POST request to create weather data
	f.fiberApp.Post("api/weather/create", func(c *fiber.Ctx) error {
		// get JSON data from request body
		jsonStr := string(c.Body())

		// call CreateWeatherDataJSON method from dbManager object
		result, err := dbManager.CreateWeatherDataJSON(jsonStr)
		if err != nil {
			f.metrics.IncrementErrorCount(c.Route().Path)
			return c.SendString(err.Error())
		}

		f.metrics.IncrementRequestCount(c.Route().Path)
		return c.SendString(result)
	})

	// GET request to retrieve metrics
	f.fiberApp.Get("api/metrics", func(c *fiber.Ctx) error {
		metrics := f.metrics.GetMetrics()

		f.metrics.IncrementRequestCount(c.Route().Path)
		return c.JSON(metrics)
	})

	// Protect the admin routes with basic authentication
	f.fiberApp.Use(BasicAuth("users.json", "Admin Access"))
	f.fiberApp.Get("/admin", func(c *fiber.Ctx) error {
		return c.SendString("Admin Page")
	})

	//main page
	f.fiberApp.Get("/", func(c *fiber.Ctx) error {
		f.metrics.IncrementRequestCount(c.Route().Path)
		//return main html page
		return c.SendFile("./web/index.html")
	})
}
