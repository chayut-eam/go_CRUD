package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Create Struct
type Person struct {
	_id       primitive.ObjectID `json:”id,omitempty”`
	FirstName string             `json:”firstname,omitempty”`
	LastName  string             `json:”lastname,omitempty”`
	Email     string             `json:”email,omitempty”`
	Age       int                `json:”age,omitempty”`
}
