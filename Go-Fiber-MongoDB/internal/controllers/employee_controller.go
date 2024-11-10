package controllers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/models"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/services"
)

func CreateEmployee(c *fiber.Ctx) error {
    var employee models.Employee
    if err := c.BodyParser(&employee); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    result, err := services.CreateEmployee(employee)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to create employee"})
    }

    return c.Status(201).JSON(result)
}

func GetEmployees(c *fiber.Ctx) error {
    employees, err := services.GetAllEmployees()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch employees"})
    }

    return c.JSON(employees)
}
