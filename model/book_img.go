package model

import "gorm.io/gorm"

type BookImg struct {
	gorm.Model
	BookId  uint `gorm:"not null"`
	ImgPath string
}
