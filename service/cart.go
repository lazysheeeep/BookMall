package service

import (
	"BookMall/dao"
	"BookMall/model"
	"BookMall/pkg/e"
	"BookMall/serializer"
	"context"
	"strconv"
)

type CartService struct {
	BossId uint `json:"boss_id" form:"boss_id"`
	BookId uint `json:"book_id" form:"book_id"`
	Num    uint `json:"num" form:"num"`
}

type DeleteCartService struct {
	Id uint `json:"id" form:"id"`
}

func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	var cart model.Cart
	var err error

	code := e.Success

	cartDao := dao.NewCartDao(ctx)
	cart, _ = cartDao.GetCart(uId, service.BookId, service.BossId)

	bookDao := dao.NewBookDao(ctx)
	book, _ := bookDao.GetBookById(service.BookId)

	if cart == (model.Cart{}) {
		cart = model.Cart{
			UserId: uId,
			BookId: service.BookId,
			BossId: service.BossId,
			Num:    service.Num,
			MaxNum: 9,
			Check:  false,
		}
		err = cartDao.Create(cart)
		if err != nil {
			code = e.ErrorDao
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Err:    err.Error(),
			}
		}

		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   serializer.BuildCart(cart, book),
		}
	} else if (cart.Num + service.Num) <= cart.MaxNum {
		cart.Num = cart.Num + service.Num
		err = cartDao.Update(cart.ID, cart)
		if err != nil {
			code = e.ErrorDao
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Err:    err.Error(),
			}
		}

		return serializer.Response{
			Status: 201,
			Msg:    "商品已经在购物车里了，商品数量+" + strconv.Itoa(int(service.Num)),
			Data:   serializer.BuildCart(cart, book),
		}
	}

	return serializer.Response{
		Status: 202,
		Msg:    "超过商品数量最大上限",
	}
}

func (service *CartService) Show(ctx context.Context, uId uint) serializer.Response {
	var carts []model.Cart
	var err error

	cartDao := dao.NewCartDao(ctx)
	carts, err = cartDao.Show(uId)
	if err != nil {
		code := e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildCarts(carts, ctx), uint(len(carts)))
}

func (service *CartService) Change(ctx context.Context, uId uint) serializer.Response {
	var cart model.Cart
	var err error

	code := e.Success

	cartDao := dao.NewCartDao(ctx)
	cart, err = cartDao.GetCart(uId, service.BookId, service.BookId)
	if err != nil {
		code = e.ErrorDao
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	if service.Num <= 0 || service.Num > 9 {
		code = e.ErrorChangeCart
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	cart.Num = service.Num
	err = cartDao.Update(cart.ID, cart)

	if err != nil {
		code = e.ErrorDao
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

func (service *DeleteCartService) Delete(ctx context.Context) serializer.Response {
	var err error

	code := e.Success

	cartDao := dao.NewCartDao(ctx)
	cart, _ := cartDao.GetCartById(service.Id)
	err = cartDao.Delete(cart)

	if err != nil {
		code = e.ErrorDao
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
