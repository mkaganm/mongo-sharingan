package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"mongo-sharingan/pkg/api"
	"os"
)

// main is the entrypoint of the application
func main() {

	app := fiber.New(fiber.Config{
		//ReadTimeout:   time.Second * 15,
		//WriteTimeout:  time.Second * 15,
		Concurrency:  10,
		ServerHeader: "MONGO_SHARINGAN_V1",
		AppName:      "MONGO_SHARINGAN_V1",
	})

	// Register routes
	api.RegisterRoutes(app)

	// Listen on port
	err := app.Listen(os.Getenv("LOCAL_SERVER_PORT"))
	if err != nil {
		log.Fatal("Failed to listen on port from env!")
	}
}
