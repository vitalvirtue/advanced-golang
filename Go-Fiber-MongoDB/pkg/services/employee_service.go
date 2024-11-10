package services

import (
    "context"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/models"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/db"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "time"
)

func CreateEmployee(employee models.Employee) (*mongo.InsertOneResult, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    return db.EmployeeCollection.InsertOne(ctx, employee)
}

func GetAllEmployees() ([]models.Employee, error) {
    var employees []models.Employee
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := db.EmployeeCollection.Find(ctx, bson.M{})
    if err != nil {
        return employees, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var employee models.Employee
        cursor.Decode(&employee)
        employees = append(employees, employee)
    }
    return employees, nil
}
