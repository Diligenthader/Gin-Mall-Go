package dao

import (
	"context"
	"github.com/CocaineCong/gin-mall/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// CreateProduct 创建商品
func (dao *ProductDao) CreateProduct(product *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Create(&product).Error
}

// CountProductByCondition 计算商品数量
func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

// ListProductByCondition 商品列表展示
func (dao *ProductDao) ListProductByCondition(condition map[string]interface{}, page model.BasePage) (products []*model.Product, err error) {
	err = dao.DB.Where(condition).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&products).Error
	return
}

// SearchProduct 查询商品
func (dao *ProductDao) SearchProduct(info string, page model.BasePage) (products []*model.Product, count int64, err error) {
	err = dao.DB.Model(&model.Product{}).
		Where("title Like ? OR info Like ?", "%"+info+"%", "%"+info+"%").Count(&count).Error
	if err != nil {
		return
	}
	err = dao.DB.Model(&model.Product{}).
		Where("title Like ? OR info Like ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * (page.PageSize)).
		Limit(page.PageSize).Find(&products).Error
	// .Find(&products).Error是一个查询操作.并将查询结果存储在products变量中
	return
}

// GetProductById 根据id查询对应的商品信息
func (dao *ProductDao) GetProductById(id uint) (product *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where("id=?", id).First(&product).Error
	return
}

func (dao *ProductDao) UpdateProduct(id uint, product *model.Product) error {
	return dao.DB.Model(&model.Product{}).Where("id=?", id).Updates(product).Error
}
