package pharmacy

import (
	"github.com/tadilbek11kz/ePharma-backend/pkg/pharmacy"
)

type Repository interface {
	CreatePharmacy(data pharmacy.CreatePharmacyRequest) (pharmacy.Pharmacy, error)
	GetAllPharmacies() ([]pharmacy.Pharmacy, error)
	GetPharmacy(id string) (pharmacy.Pharmacy, error)
	UpdatePharmacy(id string, data pharmacy.UpdatePharmacyRequest) (pharmacy.Pharmacy, error)
	DeletePharmacy(id string) error
}
