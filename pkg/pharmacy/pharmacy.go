package pharmacy

import (
	"github.com/tadilbek11kz/ePharma-backend/pkg/util"
)

type Pharmacy struct {
	util.BaseModel
	Name        string `gorm:"size:255;not null" validate:"required" json:"name"`
	Description string `gorm:"size:255;not null" json:"descritpion"`
	Location    string `gorm:"size:255;not null;" validate:"required" json:"location"`
}
