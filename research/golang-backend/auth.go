package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"log"
	"strings"
)

func BasicAuth(userFile string, realm string) fiber.Handler {
	// Read the user file into memory
	userBytes, err := ioutil.ReadFile(userFile)
	if err != nil {
		log.Fatal(err)
	}
	users := make(map[string]string)
	if err := json.Unmarshal(userBytes, &users); err != nil {
		log.Fatal(err)
	}

	// Return the middleware function
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		auth := c.Get("Authorization")
		if auth == "" {
			c.Status(fiber.StatusUnauthorized)
			c.Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
			return c.SendString("Unauthorized")
		}

		// Decode the Authorization header
		authBytes, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
		if err != nil {
			c.Status(fiber.StatusUnauthorized)
			c.Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
			return c.SendString("Unauthorized")
		}
		authString := string(authBytes)
		parts := strings.SplitN(authString, ":", 2)
		if len(parts) != 2 {
			c.Status(fiber.StatusUnauthorized)
			c.Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
			return c.SendString("Unauthorized")
		}

		// Check the username and password
		username := parts[0]
		password := parts[1]
		if storedPassword, ok := users[username]; !ok || storedPassword != password {
			c.Status(fiber.StatusUnauthorized)
			c.Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
			return c.SendString("Unauthorized")
		}

		// Call the next middleware function
		return c.Next()
	}
}
