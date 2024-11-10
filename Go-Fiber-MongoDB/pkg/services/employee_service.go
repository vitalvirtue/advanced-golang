package services

import (
    "context"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/models"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/db"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/types"
    "time"
    "log"

    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    
)

func CreateEmployee(req types.CreateEmployeeRequest) (*mongo.InsertOneResult, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    employee := models.Employee{
        ID:        primitive.NewObjectID(), 
        FirstName: req.FirstName,
        LastName:  req.LastName,
        Position:  req.Position,
        Salary:    req.Salary,
    }

    result, err := db.EmployeeCollection.InsertOne(ctx, employee)
    if err != nil {
        log.Printf("Error creating employee: %v", err)
        return nil, err
    }
    return result, nil
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
