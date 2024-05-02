package notification

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	model "github.com/tadilbek11kz/ePharma-backend/pkg/notification"
)

func WithLogging(service Service) Service {
	return &middlewareLogging{
		next: service,
	}
}

type middlewareLogging struct {
	next Service
}

func (m *middlewareLogging) CreateNotification(ctx context.Context, req model.CreateNotificationRequest) (notification model.Notification, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method":  "CreateNotification",
		"msg":     "Create notification",
		"logger":  "middlewareNotificationService",
		"payload": req,
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "CreateNotification",
			"msg":           "Create notification",
			"logger":        "middlewareNotificationService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	notification, err = m.next.CreateNotification(ctx, req)
	return
}

func (m *middlewareLogging) DeleteNotification(ctx context.Context, id string) (err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method": "DeleteNotification",
		"msg":    "Delete notification",
		"logger": "middlewareNotificationService",
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "DeleteNotification",
			"msg":           "Delete notification",
			"logger":        "middlewareNotificationService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	err = m.next.DeleteNotification(ctx, id)
	return
}
