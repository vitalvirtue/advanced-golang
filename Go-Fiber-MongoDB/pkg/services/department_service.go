package services

import (
    "context"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/models"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/db"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "time"
)

func CreateDepartment(department models.Department) (*mongo.InsertOneResult, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    return db.DepartmentCollection.InsertOne(ctx, department)
}

func GetAllDepartments() ([]models.Department, error) {
    var departments []models.Department
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    cursor, err := db.DepartmentCollection.Find(ctx, bson.M{})
    if err != nil {
        return departments, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var department models.Department
        cursor.Decode(&department)
        departments = append(departments, department)
    }
    return departments, nil
}
