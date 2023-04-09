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
	User *UserModel
}

func (p *ProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *ProductModel) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}