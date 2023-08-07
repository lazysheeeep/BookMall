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

func (dao *BookDao) ListBooks(page model.BasePage) (books []model.Book, err error) {
	err = dao.Model(&model.Book{}).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&books).Error
	return
}

func (dao *BookDao) GetBookById(id uint) (book model.Book, err error) {
	err = dao.Model(&model.Book{}).Where("id=?", id).Find(&book).Error
	return
}

func (dao *BookDao) SearchBookByName(name string, page model.BasePage) (count int64, books []model.Book, err error) {
	err = dao.Model(&model.Book{}).Where("name=?", name).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&books).Count(&count).Error
	return
}

func (dao *BookDao) SearchBookByISBN(isbn string, page model.BasePage) (count int64, books []model.Book, err error) {
	err = dao.Model(&model.Book{}).Where("isbn=?", isbn).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&books).Count(&count).Error
	return
}

func (dao *BookDao) SearchBookByAuthor(author string, page model.BasePage) (count int64, books []model.Book, err error) {
	err = dao.Model(&model.Book{}).Where("author=?", author).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&books).Count(&count).Error
	return
}

func (dao *BookDao) SearchBookByPublisher(publisher string, page model.BasePage) (count int64, books []model.Book, err error) {
	err = dao.Model(&model.Book{}).Where("publisher=?", publisher).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&books).Count(&count).Error
	return
}

func (dao *BookDao) UpdateBook(bId uint, book model.Book) error {
	return dao.Model(&model.Book{}).Where("id=?", bId).Updates(&book).Error
}
