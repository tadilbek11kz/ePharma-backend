package inventory

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	model "github.com/tadilbek11kz/ePharma-backend/pkg/inventory"
)

func WithLogging(service Service) Service {
	return &middlewareLogging{
		next: service,
	}
}

type middlewareLogging struct {
	next Service
}

func (m *middlewareLogging) CreateInventory(ctx context.Context, req model.CreateInventoryRequest) (inventory model.Inventory, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method":  "CreateInventory",
		"msg":     "Create inventory",
		"logger":  "middlewareInventoryService",
		"payload": req,
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "CreateInventory",
			"msg":           "Create inventory",
			"logger":        "middlewareInventoryService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	inventory, err = m.next.CreateInventory(ctx, req)
	return
}

func (m *middlewareLogging) GetAllInventories(ctx context.Context) (inventories []model.Inventory, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method": "GetAllInventories",
		"msg":    "Get all inventories",
		"logger": "middlewareInventoryService",
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "GetAllInventories",
			"msg":           "Get all inventories",
			"logger":        "middlewareInventoryService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	inventories, err = m.next.GetAllInventories(ctx)
	return
}

func (m *middlewareLogging) GetInventory(ctx context.Context, id string) (inventory model.Inventory, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method": "GetInventory",
		"msg":    "Get inventory",
		"logger": "middlewareInventoryService",
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "GetInventory",
			"msg":           "Get inventory",
			"logger":        "middlewareInventoryService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	inventory, err = m.next.GetInventory(ctx, id)
	return
}

func (m *middlewareLogging) UpdateInventory(ctx context.Context, id string, req model.UpdateInventoryRequest) (inventory model.Inventory, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method":  "UpdateInventory",
		"msg":     "Update inventory",
		"logger":  "middlewareInventoryService",
		"payload": req,
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "UpdateInventory",
			"msg":           "Update inventory",
			"logger":        "middlewareInventoryService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	inventory, err = m.next.UpdateInventory(ctx, id, req)
	return
}

func (m *middlewareLogging) DeleteInventory(ctx context.Context, id string) (err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method": "DeleteInventory",
		"msg":    "Delete inventory",
		"logger": "middlewareInventoryService",
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "DeleteInventory",
			"msg":           "Delete inventory",
			"logger":        "middlewareInventoryService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	err = m.next.DeleteInventory(ctx, id)
	return
}
