package notification

import (
	"github.com/google/uuid"
	"github.com/tadilbek11kz/ePharma-backend/pkg/util"
)

type Notification struct {
	util.BaseModel
	ProductID uuid.UUID `gorm:"type:uuid;not null" validate:"required" json:"product_id"`
	Email     string    `gorm:"size:255;not null" validate:"required,email" json:"email"`
}
