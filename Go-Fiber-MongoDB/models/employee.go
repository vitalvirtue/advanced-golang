package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Position  string             `json:"position" bson:"position"`
	Salary    float64            `json:"salary" bson:"salary"`
}
