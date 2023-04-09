package models

import (
	"challenge-12/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Username string `gorm:"not null" json:"username" form:"username" valid:"required~Username is required!"`
	Email string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required!,email~Invalid format email!"`
	Password string `gorm:"not null" json:"-" form:"password" valid:"required~Password is required!,minstringlength(6)~Password minimum length must be 6 characters!"`
	IsAdmin bool `gorm:"default:false" json:"isadmin" form:"isadmin"`
	Products []ProductModel `gorm:"foreignKey:UserId" json:"products"`
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}