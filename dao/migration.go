package dao

import (
	"BookMall/model"
	"fmt"
)

func migration() {
	err := Db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.Address{},
		&model.Admin{},
		&model.BasePage{},
		&model.Book{},
		&model.BookImg{},
		&model.Carousel{},
		&model.Cart{},
		&model.Favorite{},
		&model.FirstCategory{},
		&model.SecondCategory{},
		&model.Notice{},
		&model.Order{},
		&model.User{},
	)
	if err != nil {
		fmt.Println("err", err)
	}
	return
}
