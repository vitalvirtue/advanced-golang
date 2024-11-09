package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Department struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Description string      `json:"description" bson:"description"`
	ManagerID primitive.ObjectID `json:"manager_id" bson:"manager_id"`
}
