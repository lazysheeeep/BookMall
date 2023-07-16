package serializer

import (
	"BookMall/config"
	"BookMall/model"
	"BookMall/pkg/e"
)

type BookVO struct {
	ISBN          string `json:"isbn"`
	Name          string `json:"name" `
	Author        string `json:"author"`
	Publisher     string `json:"publisher"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price" `
	DiscountPrice string `json:"discount_price"`
	OnSale        bool   `json:"on_sale"`
	Num           int    `json:"num"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
	Category      string `json:"category"`
}

func BuildBook(item model.Book) BookVO {
	return BookVO{
		ISBN:          item.ISBN,
		Name:          item.Name,
		Author:        item.Author,
		Publisher:     item.Publisher,
		Info:          item.Info,
		ImgPath:       config.Host + config.HttpPort + config.BookPath + item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		OnSale:        true,
		Num:           item.Num,
		BossId:        item.BossId,
		BossName:      item.BossName,
		BossAvatar:    config.Host + config.HttpPort + config.BookPath + item.BossAvatar,
		Category:      item.Category,
	}
}

func BuildSearchResponse(items []model.Book, total uint) Response {
	return Response{
		Status: 200,
		Msg:    e.GetMsg(200),
		Data: ListData{
			List:  BuildBooks(items),
			Total: total,
		},
	}
}

func BuildBooks(books []model.Book) (items []BookVO) {
	for _, item := range books {
		tmp := BuildBook(item)
		items = append(items, tmp)
	}
	return
}
