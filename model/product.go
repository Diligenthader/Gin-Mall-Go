package model

import (
	"github.com/CocaineCong/gin-mall/cache"
	"gorm.io/gorm"
	"strconv"
)

type Product struct {
	gorm.Model
	Name          string
	CategoryId    uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"` //表示为默认为未上架状态
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}

func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (product *Product) AddView() {
	//	 增加商品点击数
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))
	cache.RedisClient.ZIncrBy(cache.Rankkey, 1, strconv.Itoa(int(product.ID))) //对这个商品的排名的一个增加.
}
