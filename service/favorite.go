package service

import (
	"BookMall/dao"
	"BookMall/model"
	"BookMall/pkg/e"
	"BookMall/serializer"
	"context"
)

type FavoriteService struct {
	BookId uint `json:"book_id" form:"book_id"`
	model.BasePage
}

func (service *FavoriteService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	var favorite model.Favorite
	var user model.User
	var boss model.User
	var book model.Book
	var err error

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserByUserId(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	bookDao := dao.NewBookDao(ctx)
	book, err = bookDao.GetBookById(service.BookId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}
	boss, err = userDao.GetUserByUserId(book.BossId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	favorite = model.Favorite{
		User:   user,
		UserId: uId,
		Book:   book,
		BookId: service.BookId,
		Boss:   boss,
		BossId: book.BossId,
	}

	daoFavorite := dao.NewFavoriteDao(ctx)
	err = daoFavorite.Create(favorite)
	if err != nil {
		code = e.ErrorCreateDaoFavorite
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *FavoriteService) Show(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	var favorites []model.Favorite
	var err error
	var count int64

	if service.PageSize == 0 {
		service.PageSize = 15
	}

	favoritesDao := dao.NewFavoriteDao(ctx)
	favorites, count, err = favoritesDao.Show(service.BasePage, uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	if count == 0 {
		code = e.ErrorNoneFavorite
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	return serializer.BuildListResponse(serializer.BuildFavorites(favorites, ctx), uint(len(favorites)))
}

func (service *FavoriteService) Delete(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	var favorite model.Favorite
	var err error

	favoriteDao := dao.NewFavoriteDao(ctx)
	favorite, err = favoriteDao.GetFavorite(uId, service.BookId)

	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	err = favoriteDao.Delete(favorite)
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
	}
}
