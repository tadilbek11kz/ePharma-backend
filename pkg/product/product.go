package product

import (
	"github.com/tadilbek11kz/ePharma-backend/pkg/util"
)

type Product struct {
	util.BaseModel
	BrandName        string `gorm:"size:255;not null;" validate:"required" json:"brand_name"`
	GenericName      string `gorm:"size:255;not null;" validate:"required" json:"generic_name"`
	Strength         string `gorm:"size:255;not null;" validate:"required" json:"strength"`
	Dosage           string `gorm:"size:255;not null;" validate:"required" json:"dosage"`
	DispenceMode     string `gorm:"size:255;not null;" validate:"required" json:"dispence_mode"`
	InsurancePlan    string `gorm:"size:255;not null;" validate:"required" json:"insurance_plan"`
	PackageSize      string `gorm:"size:255;not null;" validate:"required" json:"package_size"`
	ManufacturerName string `gorm:"size:255;not null;" validate:"required" json:"manufacturer_name"`
	Image            string `gorm:"size:255;not null;" validate:"required" json:"image"`
}
