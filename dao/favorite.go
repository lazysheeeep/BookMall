package dao

import (
	"BookMall/model"
	"context"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDbClient(ctx)}
}

func (dao *FavoriteDao) Create(favorite model.Favorite) error {
	err := dao.Model(&model.Favorite{}).Create(&favorite).Error
	return err
}
