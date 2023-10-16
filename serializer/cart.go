package serializer

import (
	"context"
	"github.com/CocaineCong/gin-mall/dao"
	"github.com/CocaineCong/gin-mall/model"
)

// 返回给前端的数据格式.

type Cart struct {
	Id            uint   `json:"id"`
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	CreatedAt     int64  `json:"created_at"`
	Name          string `json:"name"`
	Num           int    `json:"num"`
	MaxNum        int    `json:"max_num"`
	ImgPath       string `json:"img_path"`
	Check         bool   `json:"check"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) Cart {
	return Cart{
		Id:            cart.ID,
		UserId:        cart.UserId,
		ProductId:     cart.ProductId,
		CreatedAt:     cart.CreatedAt.Unix(),
		Num:           int(cart.Num),
		MaxNum:        int(cart.MaxNum),
		Check:         cart.Check,
		Name:          product.Name,
		ImgPath:       product.ImgPath,
		DiscountPrice: product.DiscountPrice,
		BossId:        boss.ID,
		BossName:      boss.UserName,
	}
}

func BuildCarts(ctx context.Context, items []*model.Cart) (carts []Cart) {
	productDao := dao.NewProductDao(ctx)
	bossDao := dao.NewUserDao(ctx)

	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		boss, err := bossDao.GetUserById(item.UserId)
		if err != nil {
			continue
		}
		cart := BuildCart(item, product, boss)
		carts = append(carts, cart)
	}
	return carts
}
