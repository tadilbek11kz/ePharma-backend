package product

import (
	pharmacyModel "github.com/tadilbek11kz/ePharma-backend/pkg/pharmacy"
	"github.com/tadilbek11kz/ePharma-backend/pkg/product"
)

type Repository interface {
	CreateProduct(data product.CreateProductRequest) (product product.Product, err error)
	GetAllProducts() (products []product.Product, err error)
	GetProduct(id string) (product product.Product, err error)
	GetProductAvailability(id string) (pharmacies []pharmacyModel.GetPharmacyAvailabilityRequest, err error)
	UpdateProduct(id string, data product.UpdateProductRequest) (product product.Product, err error)
	DeleteProduct(id string) (err error)
}
