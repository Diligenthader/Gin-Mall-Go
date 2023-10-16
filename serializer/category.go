package serializer

import "github.com/CocaineCong/gin-mall/model"

type Category struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt    int64  `json:"created_at"`
}

func BuildCategory(item model.Category) Category {
	return Category{
		Id:           item.ID,
		CategoryName: item.CategoryName,
		CreatedAt:    item.CreatedAt.Unix(),
	}
}
func BuildCategories(items []model.Category) (categories []Category) {
	for _, item := range items {
		category := BuildCategory(item)
		categories = append(categories, category)
	}
	return categories
}
