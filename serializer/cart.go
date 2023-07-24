package serializer

import (
	"BookMall/config"
	"BookMall/dao"
	"BookMall/model"
	"context"
)

type CartVO struct {
	Id            uint   `json:"id"`
	CreatedAt     uint64 `json:"create_at"`
	UserId        uint   `json:"user_id"`
	BookId        uint   `json:"book_id"`
	BookName      string `json:"book_name"`
	Num           uint   `json:"num"`
	MaxNum        uint   `json:"max_num"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	Check         bool
}

func BuildCart(cart model.Cart, book model.Book) CartVO {
	return CartVO{
		Id:            cart.ID,
		CreatedAt:     uint64(cart.CreatedAt.Unix()),
		UserId:        cart.UserId,
		BookId:        cart.BookId,
		BookName:      book.Name,
		Num:           cart.Num,
		MaxNum:        cart.MaxNum,
		ImgPath:       config.Host + config.HttpPort + config.BookPath + book.ImgPath,
		DiscountPrice: book.DiscountPrice,
		BossId:        book.BossId,
		BossName:      book.BossName,
		Check:         false,
	}
}

func BuildCarts(items []model.Cart, ctx context.Context) (carts []CartVO) {
	bookDao := dao.NewBookDao(ctx)
	for _, item := range items {
		book, _ := bookDao.GetBookById(item.BookId)
		tmp := BuildCart(item, book)
		carts = append(carts, tmp)
	}
	return
}
