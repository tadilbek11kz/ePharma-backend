package store

import (
	"github.com/tadilbek11kz/ePharma-backend/internal/connections"
	"github.com/tadilbek11kz/ePharma-backend/internal/repository/inventory"
	inventoryRepository "github.com/tadilbek11kz/ePharma-backend/internal/repository/inventory/postgres"
	"github.com/tadilbek11kz/ePharma-backend/internal/repository/notification"
	NotificationRepository "github.com/tadilbek11kz/ePharma-backend/internal/repository/notification/postgres"
	"github.com/tadilbek11kz/ePharma-backend/internal/repository/pharmacy"
	pharmacyRepository "github.com/tadilbek11kz/ePharma-backend/internal/repository/pharmacy/postgres"
	"github.com/tadilbek11kz/ePharma-backend/internal/repository/product"
	productRepository "github.com/tadilbek11kz/ePharma-backend/internal/repository/product/postgres"
	"github.com/tadilbek11kz/ePharma-backend/internal/repository/user"
	userRepository "github.com/tadilbek11kz/ePharma-backend/internal/repository/user/postgres"
)

type RepositoryStore struct {
	UserRepository         user.Repository
	PharmacyRepository     pharmacy.Repository
	ProductRepository      product.Repository
	InventoryRepository    inventory.Repository
	NotificationRepository notification.Repository
}

func New(conns *connections.Connections) *RepositoryStore {
	return &RepositoryStore{
		UserRepository:         userRepository.NewRepository(conns.Postgres),
		PharmacyRepository:     pharmacyRepository.NewRepository(conns.Postgres),
		ProductRepository:      productRepository.NewRepository(conns.Postgres),
		InventoryRepository:    inventoryRepository.NewRepository(conns.Postgres),
		NotificationRepository: NotificationRepository.NewRepository(conns.Postgres),
	}
}
