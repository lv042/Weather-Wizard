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
	// Get all weather data
	f.fiberApp.Get("/api/weather_data", func(c *fiber.Ctx) error {
		data, err := dbManager.GetAllWeatherDataJSON()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.Send(data)
	})

	// Get weather data by timestamp
	f.fiberApp.Post("/api/weather_data/timestamp", func(c *fiber.Ctx) error {
		data := new(struct {
			Timestamp string `json:"timestamp"`
		})
		if err := c.BodyParser(data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}
		result, err := dbManager.GetWeatherDataByTimestampJSON(string(c.Body()))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.SendString(result)
	})

	// Create weather data
	f.fiberApp.Post("/api/weather_data", func(c *fiber.Ctx) error {
		result, err := dbManager.CreateWeatherDataJSON(string(c.Body()))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": result,
		})
	})
}
