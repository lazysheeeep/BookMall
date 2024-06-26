package serializer

import (
	"BookMall/model"
)

type CarouselVO struct {
	ID       uint   `json:"id"`
	ImgPath  string `json:"img"`
	BookID   uint   `json:"book_id"`
	CreateAt int64  `json:"create_at"`
}

func BuildCarousel(carousel model.Carousel) CarouselVO {
	return CarouselVO{
		ID:       carousel.ID,
		ImgPath:  carousel.ImgPath,
		BookID:   carousel.BookId,
		CreateAt: carousel.CreatedAt.Unix(),
	}
}

func BuildCarousels(items []model.Carousel) (carousels []CarouselVO) {
	for _, item := range items {
		carousel := BuildCarousel(item)
		carousels = append(carousels, carousel)
	}
	return
}
