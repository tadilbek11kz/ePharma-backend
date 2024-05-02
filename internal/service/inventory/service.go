package inventory

import (
	"context"

	"github.com/tadilbek11kz/ePharma-backend/internal/app/store"
	model "github.com/tadilbek11kz/ePharma-backend/pkg/inventory"
)

//go:generate mockery --name Service
type Service interface {
	CreateInventory(ctx context.Context, req model.CreateInventoryRequest) (pharmacy model.Inventory, err error)
	GetAllInventories(ctx context.Context) (pharmacies []model.Inventory, err error)
	GetInventory(ctx context.Context, id string) (pharmacy model.Inventory, err error)
	UpdateInventory(ctx context.Context, id string, req model.UpdateInventoryRequest) (pharmacy model.Inventory, err error)
	DeleteInventory(ctx context.Context, id string) (err error)
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

func (s *service) CreateInventory(ctx context.Context, req model.CreateInventoryRequest) (inventory model.Inventory, err error) {
	return s.st.InventoryRepository.CreateInventory(req)
}

func (s *service) GetAllInventories(ctx context.Context) (pharmacies []model.Inventory, err error) {
	return s.st.InventoryRepository.GetAllInventories()
}

func (s *service) GetInventory(ctx context.Context, id string) (pharmacy model.Inventory, err error) {
	return s.st.InventoryRepository.GetInventory(id)
}

func (s *service) UpdateInventory(ctx context.Context, id string, req model.UpdateInventoryRequest) (pharmacy model.Inventory, err error) {
	return s.st.InventoryRepository.UpdateInventory(id, req)
}

func (s *service) DeleteInventory(ctx context.Context, id string) (err error) {
	return s.st.InventoryRepository.DeleteInventory(id)
}
