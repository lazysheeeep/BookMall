package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ISBN          string `gorm:"unique"`
	Name          string
	Author        string
	Publisher     string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool
	Num           int
	Category      string
	BossId        uint
	BossName      string
	BossAvatar    string
}
