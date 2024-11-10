package services

import (
    "context"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/models"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/db"
    "github.com/vitalvirtue/advanced-golang/Go-Fiber-MongoDB/pkg/types"
    "time"
    "log"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

func CreateDepartment(req types.CreateDepartmentRequest) (*mongo.InsertOneResult, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    department := models.Department{
        ID:          primitive.NewObjectID(),
        Name:        req.Name,
        Description: req.Description,
        ManagerID:   primitive.NewObjectID(),
    }

    result, err := db.DepartmentCollection.InsertOne(ctx, department)
    if err != nil {
        log.Printf("Error creating department: %v", err)
        return nil, err
    }

    return result, nil
}

func GetAllDepartments() ([]types.CreateDepartmentResponse, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    var departmentResponses []types.CreateDepartmentResponse
    cursor, err := db.DepartmentCollection.Find(ctx, bson.M{})
    if err != nil {
        log.Printf("Error fetching departments: %v", err)
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var department models.Department
        if err := cursor.Decode(&department); err != nil {
            log.Printf("Error decoding department: %v", err)
            return nil, err
        }

        departmentResponses = append(departmentResponses, types.CreateDepartmentResponse{
            ID:          department.ID.Hex(),
            Name:        department.Name,
            Description: department.Description,
            ManagerID:   department.ManagerID.Hex(),
        })
    }

    return departmentResponses, nil
}
