package dao

import (
	"context"
	"github.com/CocaineCong/gin-mall/model"
	"gorm.io/gorm"
)

type ProductImgDao struct {
	*gorm.DB
}

func NewProductImgDao(ctx context.Context) *ProductImgDao { //返回一个指向ProductImgDao的指针对其进行操作
	return &ProductImgDao{NewDBClient(ctx)}
}

func NewProductImgDaoByDB(db *gorm.DB) *ProductImgDao { //返回一个指向ProductImgDao的指针对其进行操作
	return &ProductImgDao{db}
}

func (dao *ProductImgDao) CreateProductImg(productImg *model.ProductImg) (err error) {
	return dao.DB.Model(&model.ProductImg{}).Create(&productImg).Error //在函数体内在对其数据进行操作，则需要再抽取出一个指针.
}

func (dao *ProductImgDao) ListProductImg(id uint) (productImg []*model.ProductImg, err error) {
	err = dao.DB.Model(&model.ProductImg{}).Where("product_id=?", id).Find(&productImg).Error
	return
}
