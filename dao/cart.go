package dao

import (
	"context"
	"github.com/CocaineCong/gin-mall/model"
	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{NewDBClient(ctx)}
}

func (dao *CartDao) CreateCart(in *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Create(&in).Error
}

func (dao *CartDao) GetCartByAid(cId uint) (cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("id=?", cId).First(&cart).Error
	return
}

func (dao *CartDao) ListCartByUserId(uId uint) (carts []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where("user_id=?", uId).Find(&carts).Error
	return
}

func (dao *CartDao) UpdateCartById(cId uint, cart *model.Cart) error {
	return dao.DB.Model(&model.Cart{}).Where("id=?", cId).Updates(&cart).Error

}

func (dao *CartDao) DeleteCartByCartId(cId, uId uint) error {
	return dao.DB.Model(&model.Cart{}).
		Where("id=? AND user_id=?", cId, uId).
		Delete(&model.Cart{}).
		Error
	//Todo 注意在进行数据库删除操作时的语句
}

func (dao *CartDao) UpdateCartNumById(cId uint, num int, uId uint) error {
	return dao.DB.Model(&model.Cart{}).Where("id=? And user_id=?", cId, uId).
		Update("num", num).
		Error
}
