package controllers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/types"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/services"
)

func CreateDepartment(c *fiber.Ctx) error {
    var req types.CreateDepartmentRequest

    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request payload",
        })
    }

    result, err := services.CreateDepartment(req)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create department",
        })
    }

    return c.Status(201).JSON(result)
}

func GetAllDepartments(c *fiber.Ctx) error {
    departments, err := services.GetAllDepartments()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch departments"})
    }

    return c.JSON(departments)
}
