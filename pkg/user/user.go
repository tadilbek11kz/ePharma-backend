package user

import (
	"html"
	"strings"

	"github.com/tadilbek11kz/ePharma-backend/pkg/util"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	util.BaseModel
	Email    string `gorm:"unique;index;not null" validate:"required,email" json:"email"`
	Password string `gorm:"->;<-;not null" json:"password"`
	Name     string `gorm:"size:255;not null;" validate:"required" json:"name"`
}

func (user *User) BeforeSave(tx *gorm.DB) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	//remove spaces in IIN
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))

	return nil

}
