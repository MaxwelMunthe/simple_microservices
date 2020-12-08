package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanParams struct {
	Name         string `json:"name"`
	BookName     string `json:"book_name"`
	LoanDuration string `json:"loan_duration"`
}

type LoanItem struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `json:"name"`
	BookName     string             `json:"book_name"`
	LoanDuration string             `json:"loan_duration"`
}

type UpdateLoanParams struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	BookName     string `json:"book_name"`
	LoanDuration string `json:"loan_duration"`
}
