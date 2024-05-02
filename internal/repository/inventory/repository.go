package inventory

import "github.com/tadilbek11kz/ePharma-backend/pkg/inventory"

type Repository interface {
	CreateInventory(data inventory.CreateInventoryRequest) (inventory.Inventory, error)
	GetAllInventories() ([]inventory.Inventory, error)
	GetInventory(id string) (inventory.Inventory, error)
	UpdateInventory(id string, data inventory.UpdateInventoryRequest) (inventory.Inventory, error)
	DeleteInventory(id string) error
}
