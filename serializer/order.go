package serializer

import (
	"BookMall/dao"
	"BookMall/model"
	"context"
)

type OrderVO struct {
	ID            uint    `json:"id"`
	OrderNum      uint64  `json:"order_num"`
	CreatedAt     int64   `json:"created_at"`
	UpdatedAt     int64   `json:"updated_at"`
	UserId        uint    `json:"user_id"`
	BossId        uint    `json:"boss_id"`
	BossName      string  `json:"boss_name"`
	BookId        uint    `json:"book_id"`
	BookName      string  `json:"book_name"`
	ImgPath       string  `json:"img_path"`
	Num           uint    `json:"num"`
	AddressId     uint    `json:"address_id"`
	AddressName   string  `json:"address_name"`
	AddressPhone  string  `json:"address_phone"`
	Province      string  `json:"address"`
	City          string  `json:"city"`
	Area          string  `json:"area"`
	Street        string  `json:"street"`
	Detail        string  `json:"detail"`
	State         uint    `json:"type"`
	DiscountPrice string  `json:"discount_price"`
	ExpiredTime   float64 `json:"expired_time"`
}

func BuildOrder(book model.Book, address model.Address, order model.Order) OrderVO {
	return OrderVO{
		ID:            order.ID,
		OrderNum:      order.OrderNum,
		CreatedAt:     order.CreatedAt.Unix(),
		UpdatedAt:     order.UpdatedAt.Unix(),
		UserId:        order.UserId,
		BossId:        order.BossId,
		BossName:      book.BossName,
		BookId:        book.ID,
		BookName:      book.Name,
		ImgPath:       book.ImgPath,
		Num:           order.Num,
		AddressId:     address.ID,
		AddressName:   address.Name,
		AddressPhone:  address.Phone,
		Province:      address.Province,
		City:          address.City,
		Area:          address.Area,
		Street:        address.Street,
		Detail:        address.Detail,
		State:         order.State,
		DiscountPrice: book.DiscountPrice,
		ExpiredTime:   order.ExpiredTime,
	}
}

func BuildOrders(ctx context.Context, orders []model.Order) (items []OrderVO) {
	bookDao := dao.NewBookDao(ctx)
	addressDao := dao.NewAddressDao(ctx)

	for _, item := range orders {
		book, _ := bookDao.GetBookById(item.BookId)
		address, _ := addressDao.GetAddressById(item.AddressId)
		tmp := BuildOrder(book, address, item)
		items = append(items, tmp)
	}

	return
}
