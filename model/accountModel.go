package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccountParams struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type AccountItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Address     string             `json:"address"`
	PhoneNumber string             `json:"phone_number"`
}

type UpdateAccountParams struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
