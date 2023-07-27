package dao

import (
	"BookMall/model"
	"context"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDbClient(ctx)}
}

func (dao *OrderDao) Create(order model.Order) error {
	return dao.Model(&model.Order{}).Create(&order).Error
}
