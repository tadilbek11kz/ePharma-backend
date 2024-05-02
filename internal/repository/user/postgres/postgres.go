package postgres

import (
	"fmt"

	model "github.com/tadilbek11kz/ePharma-backend/pkg/user"
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

func (p *Repository) CreateUser(data model.CreateUserRequest) (user model.User, err error) {
	user = model.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	err = p.db.Create(&user).Error
	return
}

func (p *Repository) GetUser(key string, value interface{}) (user model.User, err error) {
	err = p.db.Model(model.User{}).Where(fmt.Sprintf("%v = ?", key), value).Take(&user).Error
	return

}
