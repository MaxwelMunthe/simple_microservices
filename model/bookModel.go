package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookParams struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type BookItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Author      string             `json:"author"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
}

type UpdateBookParams struct {
	ID          string `json:"id"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
