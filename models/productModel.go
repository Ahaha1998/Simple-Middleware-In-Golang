package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type ProductModel struct {
	gorm.Model
	Title string `json:"title" form:"title" valid:"required~Product title is required!"`
	Description string `json:"desc" form:"desc" valid:"required~Product description is required!"`
	UserId uint
	User UserModel
}

func (b *ProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(b)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (b *ProductModel) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(b)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}