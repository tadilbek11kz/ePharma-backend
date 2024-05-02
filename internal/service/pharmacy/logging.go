package pharmacy

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	model "github.com/tadilbek11kz/ePharma-backend/pkg/pharmacy"
)

func WithLogging(service Service) Service {
	return &middlewareLogging{
		next: service,
	}
}

type middlewareLogging struct {
	next Service
}

func (m *middlewareLogging) CreatePharmacy(ctx context.Context, req model.CreatePharmacyRequest) (pharmacy model.Pharmacy, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method":  "CreatePharmacy",
		"msg":     "Create pharmacy",
		"logger":  "middlewarePharmacyService",
		"payload": req,
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "CreatePharmacy",
			"msg":           "Create pharmacy",
			"logger":        "middlewarePharmacyService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	pharmacy, err = m.next.CreatePharmacy(ctx, req)
	return
}

func (m *middlewareLogging) GetAllPharmacies(ctx context.Context) (pharmacies []model.Pharmacy, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method": "GetAllPharmacies",
		"msg":    "Get all pharmacies",
		"logger": "middlewarePharmacyService",
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "GetAllPharmacies",
			"msg":           "Get all pharmacies",
			"logger":        "middlewarePharmacyService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	pharmacies, err = m.next.GetAllPharmacies(ctx)
	return
}

func (m *middlewareLogging) GetPharmacy(ctx context.Context, id string) (pharmacy model.Pharmacy, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method": "GetPharmacy",
		"msg":    "Get pharmacy",
		"logger": "middlewarePharmacyService",
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "GetPharmacy",
			"msg":           "Get pharmacy",
			"logger":        "middlewarePharmacyService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	pharmacy, err = m.next.GetPharmacy(ctx, id)
	return
}

func (m *middlewareLogging) UpdatePharmacy(ctx context.Context, id string, req model.UpdatePharmacyRequest) (pharmacy model.Pharmacy, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method":  "UpdatePharmacy",
		"msg":     "Update pharmacy",
		"logger":  "middlewarePharmacyService",
		"payload": req,
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "UpdatePharmacy",
			"msg":           "Update pharmacy",
			"logger":        "middlewarePharmacyService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	pharmacy, err = m.next.UpdatePharmacy(ctx, id, req)
	return
}

func (m *middlewareLogging) DeletePharmacy(ctx context.Context, id string) (err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method": "DeletePharmacy",
		"msg":    "Delete pharmacy",
		"logger": "middlewarePharmacyService",
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "DeletePharmacy",
			"msg":           "Delete pharmacy",
			"logger":        "middlewarePharmacyService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	err = m.next.DeletePharmacy(ctx, id)
	return
}
