package controllers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/services"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/types"
)

func CreateEmployee(c *fiber.Ctx) error {
    var req types.CreateEmployeeRequest

    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request payload",
        })
    }

    // Employee oluşturma işlemi
    result, err := services.CreateEmployee(req)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create employee",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(result)
}

func GetEmployees(c *fiber.Ctx) error {
    employees, err := services.GetAllEmployees()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch employees"})
    }

    return c.JSON(employees)
}
