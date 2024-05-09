package inventory

import (
	"github.com/google/uuid"
	"github.com/tadilbek11kz/ePharma-backend/pkg/util"
)

type Inventory struct {
	util.BaseModel
	ProductID  uuid.UUID `gorm:"type:uuid;not null" validate:"required" json:"product_id"`
	PharmacyID uuid.UUID `gorm:"type:uuid;not null" validate:"required" json:"pharmacy_id"`
	Price      float64   `gorm:"not null" validate:"required" json:"price"`
	Quantity   int       `gorm:"not null" validate:"required" json:"quantity"`
}

// func (inventory *Inventory) AfterSave(tx *gorm.DB) error {
// 	var emails []string
// 	err := tx.Model(notification.Notification{}).Select("email").Find(&emails, "product_id = ?", inventory.ProductID).Error
// 	if err != nil {
// 		return err
// 	}

// 	dispatcher := email.NewDispatcher()

// 	// err = dispatcher.SendEmail(emails)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	return nil
// }
