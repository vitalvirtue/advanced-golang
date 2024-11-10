package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/db"
	"github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.ConnectDB()
	db.InitCollections()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}