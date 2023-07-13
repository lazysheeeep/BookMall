package serializer

import (
	"BookMall/config"
	"BookMall/model"
)

type BookVO struct {
	ISBN           string `json:"isbn"`
	Name           string `json:"name" `
	Author         string `json:"author"`
	Publisher      string `json:"publisher"`
	Info           string `json:"info"`
	ImgPath        string `json:"img_path"`
	Price          string `json:"price" `
	DiscountPrice  string `json:"discount_price"`
	OnSale         bool   `json:"on_sale"`
	Num            int    `json:"num"`
	FirstCategory  string `json:"first_category"`
	SecondCategory string `json:"second_category"`
}

func BuildBook(item model.Book) BookVO {
	return BookVO{
		ISBN:           item.ISBN,
		Name:           item.Name,
		Author:         item.Author,
		Publisher:      item.Publisher,
		Info:           item.Info,
		ImgPath:        config.Host + config.HttpPort + config.BookPath + item.ImgPath,
		Price:          item.Price,
		DiscountPrice:  item.DiscountPrice,
		OnSale:         true,
		Num:            item.Num,
		FirstCategory:  item.FirstCategory,
		SecondCategory: item.SecondCategory,
	}
}
