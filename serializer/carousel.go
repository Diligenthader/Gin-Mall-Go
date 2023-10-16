package serializer

import (
	"github.com/CocaineCong/gin-mall/model"
)

type Carousel struct {
	Id        uint   `json:"id"`
	ImgPath   string `json:"img_path"`
	ProductId uint   `json:"product_id"`
	CreatedAt int64  `json:"create_at"`
}

func BuildCarousel(item *model.Carousel) Carousel {
	return Carousel{
		Id:        item.ID,
		ImgPath:   item.ImgPath,
		ProductId: item.ProductId,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

func BuildCarousels(items []model.Carousel) (carousels []Carousel) {
	for _, item := range items {
		carousel := BuildCarousel(&item)
		carousels = append(carousels, carousel)
	}
	return carousels
}
