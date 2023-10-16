package dao

import (
	"context"
	"github.com/CocaineCong/gin-mall/model"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

// ListCategory 用于展示
func (dao *CategoryDao) ListCategory() (Category []model.Category, err error) { //todo 切片和普通指针类型的差异区别.
	err = dao.DB.Model(&model.Category{}).Find(&Category).Error
	return
}
