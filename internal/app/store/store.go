package store

import (
	"github.com/tadilbek11kz/ePharma-backend/internal/connections"
	"github.com/tadilbek11kz/ePharma-backend/internal/repository/pharmacy"
	pharmacyRepository "github.com/tadilbek11kz/ePharma-backend/internal/repository/pharmacy/postgres"
	"github.com/tadilbek11kz/ePharma-backend/internal/repository/product"
	productRepository "github.com/tadilbek11kz/ePharma-backend/internal/repository/product/postgres"
	"github.com/tadilbek11kz/ePharma-backend/internal/repository/user"
	userRepository "github.com/tadilbek11kz/ePharma-backend/internal/repository/user/postgres"
)

type RepositoryStore struct {
	UserRepository     user.Repository
	PharmacyRepository pharmacy.Repository
	ProductRepository  product.Repository
}

func New(conns *connections.Connections) *RepositoryStore {
	return &RepositoryStore{
		UserRepository:     userRepository.NewRepository(conns.Postgres),
		PharmacyRepository: pharmacyRepository.NewRepository(conns.Postgres),
		ProductRepository:  productRepository.NewRepository(conns.Postgres),
	}
}
