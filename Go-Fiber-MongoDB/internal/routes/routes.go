package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/internal/controllers"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")

    api.Post("/employees", controllers.CreateEmployee)
    api.Get("/employees", controllers.GetAllEmployees)

    api.Post("/departments", controllers.CreateDepartment)
    api.Get("/departments", controllers.GetAllDepartments)
}
