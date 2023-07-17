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

func (dao *FavoriteDao) Show(page model.BasePage, uId uint) (favorites []model.Favorite, count int64, err error) {
	err = dao.Model(&model.Favorite{}).Where("user_id=?", uId).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&favorites).Count(&count).Error
	return
}
