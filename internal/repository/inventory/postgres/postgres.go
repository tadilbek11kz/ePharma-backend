package postgres

import (
	model "github.com/tadilbek11kz/ePharma-backend/pkg/inventory"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (p *Repository) CreateInventory(data model.CreateInventoryRequest) (model.Inventory, error) {
	inventory := model.Inventory{
		ProductID:  data.ProductID,
		PharmacyID: data.PharmacyID,
		Price:      data.Price,
		Quantity:   data.Quantity,
	}
	err := p.db.Create(&inventory).Error
	return inventory, err
}

func (p *Repository) GetAllInventories() ([]model.Inventory, error) {
	var inventories []model.Inventory
	err := p.db.Find(&inventories).Error
	return inventories, err
}

func (p *Repository) GetInventory(id string) (model.Inventory, error) {
	var inventory model.Inventory
	err := p.db.Where("id = ?", id).First(&inventory).Error
	return inventory, err
}

func (p *Repository) UpdateInventory(id string, data model.UpdateInventoryRequest) (model.Inventory, error) {
	inventory := model.Inventory{
		Price:    data.Price,
		Quantity: data.Quantity,
	}
	err := p.db.Model(&model.Inventory{}).Where("id = ?", id).Updates(&inventory).Error
	return inventory, err
}

func (p *Repository) DeleteInventory(id string) error {
	err := p.db.Where("id = ?", id).Delete(&model.Inventory{}).Error
	return err
}
