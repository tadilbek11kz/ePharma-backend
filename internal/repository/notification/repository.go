package notification

import (
	"github.com/tadilbek11kz/ePharma-backend/pkg/notification"
)

type Repository interface {
	CreateNotification(data notification.CreateNotificationRequest) (notification.Notification, error)
	DeleteNotification(id string) error
}
