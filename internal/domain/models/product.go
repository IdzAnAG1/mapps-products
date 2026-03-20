package models

import "github.com/google/uuid"

type Product struct {
	ID             string
	Name           string
	Description    string
	Price          float64
	Category       string
	VirtualImageID string
	ModelID        string
}

func NewProduct(name, description string, price float64, category, virtualImageID, modelID string) *Product {
	return &Product{
		ID:             uuid.NewString(),
		Name:           name,
		Description:    description,
		Price:          price,
		Category:       category,
		VirtualImageID: virtualImageID,
		ModelID:        modelID,
	}
}
