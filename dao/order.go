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

func (dao *OrderDao) GetOrders(uId uint, pageNum int, pageSize int) (orders []model.Order, err error) {
	err = dao.Model(&model.Order{}).Where("user_id=?", uId).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&orders).Error
	return
}

func (dao *OrderDao) GetOrderId(id uint) (order model.Order, err error) {
	err = dao.Model(&model.Order{}).Where("id=?", id).Find(&order).Error
	return
}
