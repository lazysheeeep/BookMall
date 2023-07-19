package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User   User `gorm:"ForeignKey:UserId"`
	UserId uint `gorm:"not null"`
	Book   Book `gorm:"ForeignKey:BookId"`
	BookId uint `gorm:"not null"`
	Boss   User `gorm:"ForeignKey:BossId"`
	BossId uint `gorm:"not null"`
}
