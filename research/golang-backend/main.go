package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
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
		switch c.Method() {
		case "GET":
			// Read all weather data
			// Get the weather data
			data, err := dbManager.ReadWeatherData()
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Error: " + err.Error())
			}
			if data == "" {
				return c.SendString("No weather data found")
			}
			return c.SendString(data)
		case "DELETE":
			// Delete weather data
			weatherData := new(WeatherData)
			if err := c.BodyParser(weatherData); err != nil {
				fmt.Println("Error: " + err.Error())
				return c.Status(fiber.StatusBadRequest).SendString("Error: " + err.Error())
			}
			if err := dbManager.DeleteWeatherData(fmt.Sprintf("%v", weatherData.Timestamp)); err != nil {
				fmt.Println("Error: " + err.Error())
				return c.Status(fiber.StatusInternalServerError).SendString("Error: " + err.Error())
			}
			return c.SendString("Weather data deleted")
		default:
			return c.Status(fiber.StatusMethodNotAllowed).SendString("Method not allowed")
		}
	}

	// Register weather data route
	f.fiberApp.Route("/weather-data", func(r fiber.Router) {
		r.Use("/", func(c *fiber.Ctx) error {
			// Middleware function to log requests
			log.Printf("Request received: %s %s", c.Method(), c.Path())
			return c.Next()
		})

		r.Get("/", weatherDataHandler)
		r.Post("/", weatherDataHandler)
		r.Put("/", weatherDataHandler)
		r.Delete("/", weatherDataHandler)
	})
}
