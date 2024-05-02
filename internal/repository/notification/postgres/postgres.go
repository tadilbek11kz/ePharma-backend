package postgres

import (
	model "github.com/tadilbek11kz/ePharma-backend/pkg/notification"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (p *Repository) CreateNotification(data model.CreateNotificationRequest) (notification model.Notification, err error) {
	notification = model.Notification{
		Email:     data.Email,
		ProductID: data.ProductID,
	}
	err = p.db.Create(&notification).Error
	return
}

func (p *Repository) DeleteNotification(id string) (err error) {
	err = p.db.Delete(&model.Notification{}, "id = ?", id).Error
	return
}
