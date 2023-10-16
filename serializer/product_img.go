package serializer

import (
	"github.com/CocaineCong/gin-mall/conf"
	"github.com/CocaineCong/gin-mall/model"
)

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

// BuildProductImg 返回单条查询数据
func BuildProductImg(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductId: item.ProductId,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
	}
}

// BuildProductImgs 返回多条查询数据
func BuildProductImgs(items []*model.ProductImg) (productImg []ProductImg) {
	for _, item := range items {
		//使用 range 迭代切片时，它会返回两个值：索引和元素值。
		product := BuildProductImg(item)
		productImg = append(productImg, product)
	}
	return
}
