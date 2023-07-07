package dao

import (
	"BookMall/model"
	"context"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDbClient(ctx)}
}

func (dao *CarouselDao) ListCarousel() (items []model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&items).Error
	return
}
