package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/"
	"github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/db"
	"github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/internals/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		Log.Fatal("Error loading .env file")
	}

	
}