package postgres

import (
	model "github.com/tadilbek11kz/ePharma-backend/pkg/pharmacy"
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

func (p *Repository) CreatePharmacy(data model.CreatePharmacyRequest) (pharmacy model.Pharmacy, err error) {
	pharmacy = model.Pharmacy{
		Name:        data.Name,
		Description: data.Description,
		Location:    data.Location,
	}
	err = p.db.Create(&pharmacy).Error
	return
}

func (p *Repository) GetAllPharmacies() (pharmacies []model.Pharmacy, err error) {
	err = p.db.Find(&pharmacies).Error
	return
}

func (p *Repository) GetPharmacy(id string) (pharmacy model.Pharmacy, err error) {
	err = p.db.First(&pharmacy, "id = ?", id).Error
	return
}

func (p *Repository) UpdatePharmacy(id string, data model.UpdatePharmacyRequest) (pharmacy model.Pharmacy, err error) {
	err = p.db.Model(&pharmacy).Where("id = ?", id).Updates(data).Error
	return
}

func (p *Repository) DeletePharmacy(id string) (err error) {
	err = p.db.Delete(&model.Pharmacy{}, "id = ?", id).Error
	return
}
