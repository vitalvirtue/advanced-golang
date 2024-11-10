package db

import (
    "go.mongodb.org/mongo-driver/mongo"
)

var (
    DepartmentCollection *mongo.Collection
    EmployeeCollection   *mongo.Collection
)

func InitCollections() {
    if HRMSDatabase != nil {
        DepartmentCollection = HRMSDatabase.Collection("departments")
        EmployeeCollection = HRMSDatabase.Collection("employees")
    }
}
