package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Debt struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Amount    float64            `json:"amount" bson:"amount"`
	DueDate   time.Time          `json:"dueDate" bson:"dueDate"`
	CPF       string             `json:"cpf" bson:"cpf"`
	UserEmail string             `json:"userEmail" bson:"userEmail"`
	Origin    string             `json:"origin" bson:"origin"`
}
