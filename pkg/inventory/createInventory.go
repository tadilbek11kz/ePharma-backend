package inventory

import "github.com/google/uuid"

type CreateInventoryRequest struct {
	ProductID  uuid.UUID `json:"product_id" binding:"required"`
	PharmacyID uuid.UUID `json:"pharmacy_id" binding:"required"`
	Price      float64   `json:"price" binding:"required"`
	Quantity   int       `json:"quantity" binding:"required"`
}
