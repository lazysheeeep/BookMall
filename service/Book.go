package service

import (
	"BookMall/dao"
	"BookMall/model"
	"BookMall/pkg/e"
	"BookMall/serializer"
	"context"
	"mime/multipart"
)

type BookService struct {
	ISBN          string `json:"isbn" form:"isbn"`
	Name          string `json:"name" form:"name"`
	Author        string `json:"author" form:"author"`
	Publisher     string `json:"publisher" form:"publisher"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	Category      string `json:"category" form:"category"`
	model.BasePage
}

type SearchBookService struct {
	OperationType uint   `json:"operation_type" form:"operation_type"`
	Name          string `json:"name" form:"name"`
	ISBN          string `json:"isbn" form:"isbn"`
	Author        string `json:"author" form:"author"`
	Publisher     string `json:"publisher" form:"publisher"`
	model.BasePage
}

func (service *BookService) Create(ctx context.Context, file []*multipart.FileHeader, uID uint) serializer.Response {
	code := e.Success
	var book model.Book
	userDao := dao.NewUserDao(ctx)
	boss, _ := userDao.GetUserByUserId(uID)
	//选择图书封面
	tmp, _ := file[0].Open()
	path, err := UploadBookToLocalStatic(tmp, uID, service.Name)
	if err != nil {
		code = e.ErrorUploadBookToStatic
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	book = model.Book{
		ISBN:          service.ISBN,
		Name:          service.Name,
		Author:        service.Author,
		Publisher:     service.Publisher,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		Category:      service.Category,
		BossId:        uID,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}

	bookDao := dao.NewBookDao(ctx)
	err = bookDao.CreateBook(book)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildBook(book),
	}
}

func (service *SearchBookService) Search(ctx context.Context) serializer.Response {
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	typeId := service.OperationType
	bookDao := dao.NewBookDao(ctx)
	var books []model.Book
	var err error
	var count int64
	flag := false
	code := e.Success
	switch typeId {
	case 1: //名字查找
		count, books, err = bookDao.SearchBookByName(service.Name, service.BasePage)
	case 2: //ISBN查找
		count, books, err = bookDao.SearchBookByISBN(service.ISBN, service.BasePage)
	case 3: //作者查找
		count, books, err = bookDao.SearchBookByAuthor(service.Author, service.BasePage)
	case 4: //出版社查找
		count, books, err = bookDao.SearchBookByPublisher(service.Publisher, service.BasePage)
	default:
		flag = true
	}
	if flag {
		code = e.ErrorOperationType
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	if count == 0 {
		code = e.ErrorBookNotExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildBooks(books), uint(len(books)))
}

func (service *BookService) List(ctx context.Context) serializer.Response {
	code := e.Success
	var books []model.Book
	var err error

	if service.PageSize == 0 {
		service.PageSize = 15
	}

	bookDao := dao.NewBookDao(ctx)
	books, err = bookDao.ListBooks(service.BasePage)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildBooks(books), uint(len(books)))
}
