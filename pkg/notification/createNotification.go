package notification

import "github.com/google/uuid"

type CreateNotificationRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Email     string    `json:"email" binding:"required"`
}
