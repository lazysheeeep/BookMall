package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Category string `gorm:"unique"`
	Text     string
}
