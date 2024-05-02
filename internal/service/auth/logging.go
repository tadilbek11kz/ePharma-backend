package auth

import (
	"context"
	"time"

	"github.com/tadilbek11kz/ePharma-backend/internal/util/token"
	model "github.com/tadilbek11kz/ePharma-backend/pkg/user"

	"github.com/sirupsen/logrus"
)

func WithLogging(service Service) Service {
	return &middlewareLogging{
		next: service,
	}
}

type middlewareLogging struct {
	next Service
}

func (m *middlewareLogging) CreateUser(ctx context.Context, req model.CreateUserRequest) (user model.User, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method":  "CreateUser",
		"msg":     "Create user",
		"logger":  "middlewareUserService",
		"payload": req,
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "CreateUser",
			"msg":           "Crate user",
			"logger":        "middlewareUserService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	user, err = m.next.CreateUser(ctx, req)
	return
}

func (m *middlewareLogging) LoginUser(ctx context.Context, req model.LoginUserRequest) (jwt token.Token, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method":  "LoginUser",
		"msg":     "Login user",
		"logger":  "middlewareUserService",
		"payload": req,
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "LoginUser",
			"msg":           "Login user",
			"logger":        "middlewareUserService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	jwt, err = m.next.LoginUser(ctx, req)
	return
}

func (m *middlewareLogging) RefreshToken(ctx context.Context, cookie string) (jwt token.Token, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method":  "RefreshToken",
		"msg":     "Refresh token",
		"logger":  "middlewareUserService",
		"payload": cookie,
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "RefreshToken",
			"msg":           "Refresh token",
			"logger":        "middlewareUserService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	jwt, err = m.next.RefreshToken(ctx, cookie)
	return
}
