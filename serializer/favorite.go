package serializer

import (
	"BookMall/dao"
	"BookMall/model"
	"context"
)

type FavoriteVO struct {
	UserId        uint   `json:"user_id"`
	BookId        uint   `json:"book_id"`
	CreateAt      int64  `json:"create_at"`
	Name          string `json:"name"`
	Category      string `json:"category"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

func BuildFavorite(favorite model.Favorite, book model.Book, boss model.User) FavoriteVO {
	return FavoriteVO{
		UserId:        favorite.UserId,
		BookId:        favorite.BookId,
		CreateAt:      favorite.CreatedAt.Unix(),
		Name:          book.Name,
		Category:      book.Category,
		Info:          book.Category,
		ImgPath:       book.ImgPath,
		Price:         book.Price,
		DiscountPrice: book.DiscountPrice,
		Num:           book.Num,
		BossId:        boss.ID,
		OnSale:        true,
	}
}

func BuildFavorites(favourites []model.Favorite, ctx context.Context) (items []FavoriteVO) {
	bookDao := dao.NewBookDao(ctx)
	bossDao := dao.NewUserDao(ctx)
	for _, item := range favourites {
		book := model.Book{}
		boss := model.User{}
		book, _ = bookDao.GetBookById(item.BookId)
		boss, _ = bossDao.GetUserByUserId(item.BossId)
		tmp := BuildFavorite(item, book, boss)
		items = append(items, tmp)
	}
	return
}
