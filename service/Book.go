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
	ISBN           string `json:"isbn" form:"isbn"`
	Name           string `json:"name" form:"name"`
	Author         string `json:"author" form:"author"`
	Publisher      string `json:"publisher" form:"publisher"`
	Info           string `json:"info" form:"info"`
	ImgPath        string `json:"img_path" form:"img_path"`
	Price          string `json:"price" form:"price"`
	DiscountPrice  string `json:"discount_price" form:"discount_price"`
	OnSale         bool   `json:"on_sale" form:"on_sale"`
	Num            int    `json:"num" form:"num"`
	FirstCategory  string `json:"first_category" form:"first_category"`
	SecondCategory string `json:"second_category" form:"second_category"`
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
		ISBN:           service.ISBN,
		Name:           service.Name,
		Author:         service.Author,
		Publisher:      service.Publisher,
		Info:           service.Info,
		ImgPath:        path,
		Price:          service.Price,
		DiscountPrice:  service.DiscountPrice,
		OnSale:         true,
		Num:            service.Num,
		FirstCategory:  service.FirstCategory,
		SecondCategory: service.SecondCategory,
		BossId:         uID,
		BossName:       boss.UserName,
		BossAvatar:     boss.Avatar,
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
