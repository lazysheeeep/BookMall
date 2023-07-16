package service

import (
	"BookMall/dao"
	"BookMall/model"
	"BookMall/pkg/e"
	"BookMall/serializer"
	"context"
)

type CategoryService struct {
}

func (service *CategoryService) ListCategory(ctx context.Context) serializer.Response {
	code := e.Success
	var categories []model.Category
	var err error
	categoryDao := dao.NewCategoryDao(ctx)

	categories, err = categoryDao.GetCategory()
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Err:    err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildCategories(categories), uint(len(categories)))
}
