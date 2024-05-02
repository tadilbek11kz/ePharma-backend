package notification

import (
	"context"

	"github.com/tadilbek11kz/ePharma-backend/internal/app/store"
	model "github.com/tadilbek11kz/ePharma-backend/pkg/notification"
)

//go:generate mockery --name Service
type Service interface {
	CreateNotification(ctx context.Context, req model.CreateNotificationRequest) (notification model.Notification, err error)
	DeleteNotification(ctx context.Context, id string) (err error)
}

type service struct {
	st *store.RepositoryStore
}

func New(st *store.RepositoryStore) (srv Service) {
	srv = &service{
		st: st,
	}
	srv = WithLogging(srv)
	return
}

func (s *service) CreateNotification(ctx context.Context, req model.CreateNotificationRequest) (notification model.Notification, err error) {
	return s.st.NotificationRepository.CreateNotification(req)
}

func (s *service) DeleteNotification(ctx context.Context, id string) (err error) {
	return s.st.NotificationRepository.DeleteNotification(id)
}
