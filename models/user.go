package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nickname    string `json:"nickname" gorm:"unique" validate:"required,min=2,max=32"`
	Description string `json:"description" validate:"max=120"`
	Password    string `json:"-" validate:"required,min=8,max=32"`
}
