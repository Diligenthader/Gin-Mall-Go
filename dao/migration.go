package dao

import (
	"fmt"
	"github.com/CocaineCong/gin-mall/model"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.User{},
			&model.Address{},
			&model.Admin{},
			&model.Category{},
			&model.Carousel{},
			&model.Cart{},
			&model.Notice{},
			&model.Product{},
			&model.ProductImg{},
			&model.Order{},
			&model.Favorite{})
	if err != nil {
		fmt.Println("err", err)
	}
	/*
		根据提供的代码，migration() 函数使用 GORM 进行数据库迁移操作。在迁移过程中，它自动创建或更新数据库表结构以匹配指定的模型。
		在迁移过程中，通过使用 _db.Set("gorm:table_options", "charset=utf8mb4") 设置了 gorm:table_options 选项，将字符集设置为 UTF-8MB4，以支持更广泛的字符编码。
	*/
	return
}
