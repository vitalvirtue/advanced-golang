package controllers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/models"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/services"
)

func CreateDepartment(c *fiber.Ctx) error {
    var department models.Department
    if err := c.BodyParser(&department); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    result, err := services.CreateDepartment(department)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to create department"})
    }

    return c.Status(201).JSON(result)
}

func GetDepartments(c *fiber.Ctx) error {
    departments, err := services.GetAllDepartments()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch departments"})
    }

    return c.JSON(departments)
}
