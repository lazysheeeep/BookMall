package dao

import (
	"BookMall/model"
	"context"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDbClient(ctx)}
}

func (dao *CategoryDao) GetCategory() (categories []model.Category, err error) {
	err = dao.Model(&model.Category{}).Find(&categories).Error
	return
}
