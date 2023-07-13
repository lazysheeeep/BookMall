package dao

import (
	"BookMall/model"
	"context"
	"gorm.io/gorm"
)

type BookDao struct {
	*gorm.DB
}

func NewBookDao(ctx context.Context) *BookDao {
	return &BookDao{NewDbClient(ctx)}
}

func (dao *BookDao) CreateBook(book model.Book) error {
	return dao.Model(&model.Book{}).Create(&book).Error
}
