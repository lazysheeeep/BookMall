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

func (dao *BookDao) SearchBookByName(name string) (count int64, books []model.Book, err error) {
	err = dao.Model(&model.Book{}).Where("name=?", name).Find(&books).Count(&count).Error
	return
}

func (dao *BookDao) SearchBookByISBN(isbn string) (count int64, books []model.Book, err error) {
	err = dao.Model(&model.Book{}).Where("isbn=?", isbn).Find(&books).Count(&count).Error
	return
}

func (dao *BookDao) SearchBookByAuthor(author string) (count int64, books []model.Book, err error) {
	err = dao.Model(&model.Book{}).Where("author=?", author).Find(&books).Count(&count).Error
	return
}

func (dao *BookDao) SearchBookByPublisher(publisher string) (count int64, books []model.Book, err error) {
	err = dao.Model(&model.Book{}).Where("publisher=?", publisher).Find(&books).Count(&count).Error
	return
}
