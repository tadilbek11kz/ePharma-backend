package product

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	model "github.com/tadilbek11kz/ePharma-backend/pkg/product"
)

func WithLogging(service Service) Service {
	return &middlewareLogging{
		next: service,
	}
}

type middlewareLogging struct {
	next Service
}

func (m *middlewareLogging) CreateProduct(ctx context.Context, req model.CreateProductRequest) (product model.Product, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method":  "CreateProduct",
		"msg":     "Create product",
		"logger":  "middlewareProductService",
		"payload": req,
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "CreateProduct",
			"msg":           "Create product",
			"logger":        "middlewareProductService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	product, err = m.next.CreateProduct(ctx, req)
	return
}

func (m *middlewareLogging) GetAllProducts(ctx context.Context) (products []model.Product, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method": "GetAllProducts",
		"msg":    "Get all products",
		"logger": "middlewareProductService",
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "GetAllProducts",
			"msg":           "Get all products",
			"logger":        "middlewareProductService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	products, err = m.next.GetAllProducts(ctx)
	return
}

func (m *middlewareLogging) GetProduct(ctx context.Context, id string) (product model.Product, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method": "GetProduct",
		"msg":    "Get product",
		"logger": "middlewareProductService",
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "GetProduct",
			"msg":           "Get product",
			"logger":        "middlewareProductService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	product, err = m.next.GetProduct(ctx, id)
	return
}

func (m *middlewareLogging) UpdateProduct(ctx context.Context, id string, req model.UpdateProductRequest) (product model.Product, err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method":  "UpdateProduct",
		"msg":     "Update product",
		"logger":  "middlewareProductService",
		"payload": req,
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "UpdateProduct",
			"msg":           "Update product",
			"logger":        "middlewareProductService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	product, err = m.next.UpdateProduct(ctx, id, req)
	return
}

func (m *middlewareLogging) DeleteProduct(ctx context.Context, id string) (err error) {
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"method": "DeleteProduct",
		"msg":    "Delete product",
		"logger": "middlewareProductService",
	}).Info()

	defer func(begin time.Time) {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"method":        "DeleteProduct",
			"msg":           "Delete product",
			"logger":        "middlewareProductService",
			"responseError": err,
			"elapsedTime":   time.Since(begin),
		}).Info()
	}(time.Now())

	err = m.next.DeleteProduct(ctx, id)
	return
}
