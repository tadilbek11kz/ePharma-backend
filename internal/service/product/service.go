package product

import (
	"context"

	"github.com/tadilbek11kz/ePharma-backend/internal/app/store"
	pharmacyModel "github.com/tadilbek11kz/ePharma-backend/pkg/pharmacy"
	model "github.com/tadilbek11kz/ePharma-backend/pkg/product"
)

//go:generate mockery --name Service
type Service interface {
	CreateProduct(ctx context.Context, req model.CreateProductRequest) (product model.Product, err error)
	GetAllProducts(ctx context.Context) (products []model.Product, err error)
	GetProduct(ctx context.Context, id string) (product model.Product, err error)
	GetProductAvailability(ctx context.Context, id string) (pharmacies []pharmacyModel.GetPharmacyAvailabilityRequest, err error)
	UpdateProduct(ctx context.Context, id string, req model.UpdateProductRequest) (product model.Product, err error)
	DeleteProduct(ctx context.Context, id string) (err error)
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

func (s *service) CreateProduct(ctx context.Context, req model.CreateProductRequest) (product model.Product, err error) {
	return s.st.ProductRepository.CreateProduct(req)
}

func (s *service) GetAllProducts(ctx context.Context) (products []model.Product, err error) {
	return s.st.ProductRepository.GetAllProducts()
}

func (s *service) GetProduct(ctx context.Context, id string) (product model.Product, err error) {
	return s.st.ProductRepository.GetProduct(id)
}

func (s *service) UpdateProduct(ctx context.Context, id string, req model.UpdateProductRequest) (product model.Product, err error) {
	return s.st.ProductRepository.UpdateProduct(id, req)
}

func (s *service) DeleteProduct(ctx context.Context, id string) (err error) {
	return s.st.ProductRepository.DeleteProduct(id)
}

func (s *service) GetProductAvailability(ctx context.Context, id string) (pharmacies []pharmacyModel.GetPharmacyAvailabilityRequest, err error) {
	return s.st.ProductRepository.GetProductAvailability(id)
}
