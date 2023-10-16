package dao

import (
	"context"
	"github.com/CocaineCong/gin-mall/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

func (dao *OrderDao) CreateOrder(in *model.Order) error {
	return dao.DB.Model(&model.Order{}).Create(&in).Error
}

func (dao *OrderDao) GetOrderById(id, userId uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", id, userId).First(&order).Error
	return
}

func (dao *OrderDao) ListOrderByUserId(uId uint) (orders []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("user_id=?", uId).Find(&orders).Error
	return
}

func (dao *OrderDao) UpdateOrderByUserId(aId uint, order *model.Order) error {
	return dao.DB.Model(&model.Order{}).Where("id=?", aId).Updates(&order).Error

}

func (dao *OrderDao) DeleteOrderByOrderId(aId, uId uint) error {
	return dao.DB.Model(&model.Order{}).Where("id=? AND user_id=?", aId, uId).Delete(&model.Order{}).Error
	//Todo 注意在进行数据库删除操作时的语句 这是一个软删除操作
}

func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (order []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).
		Where(condition).
		Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).
		Find(&order).Error
	return
}
