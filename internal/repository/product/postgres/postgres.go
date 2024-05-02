package postgres

import (
	model "github.com/tadilbek11kz/ePharma-backend/pkg/product"
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

func (p *Repository) CreateProduct(data model.CreateProductRequest) (product model.Product, err error) {
	product = model.Product{
		BrandName:        data.BrandName,
		GenericName:      data.GenericName,
		Strength:         data.Strength,
		Dosage:           data.Dosage,
		DispenceMode:     data.DispenceMode,
		InsurancePlan:    data.InsurancePlan,
		PackageSize:      data.PackageSize,
		ManufacturerName: data.ManufacturerName,
	}
	err = p.db.Create(&product).Error
	return
}

func (p *Repository) GetAllProducts() (products []model.Product, err error) {
	err = p.db.Find(&products).Error
	return
}

func (p *Repository) GetProduct(id string) (product model.Product, err error) {
	err = p.db.First(&product, id).Error
	return
}

func (p *Repository) UpdateProduct(id string, data model.UpdateProductRequest) (product model.Product, err error) {
	err = p.db.Model(&product).Where("id = ?", id).Updates(data).Error
	return
}

func (p *Repository) DeleteProduct(id string) (err error) {
	err = p.db.Delete(&model.Product{}, "id = ?", id).Error
	return
}
