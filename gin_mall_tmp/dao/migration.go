package dao

import (
	"fmt"
	"gin_mall_tmp/model"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model.User{},
		&model.Address{},
		&model.Admin{},
		&model.Category{},
		&model.Cart{},
		&model.Carousel{},
		&model.Notice{},
		&model.Product{},
		&model.ProductImg{},
		&model.Order{},
		&model.Favorite{})
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("register table success")
}
