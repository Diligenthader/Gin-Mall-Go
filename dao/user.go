package dao

import (
	"context"
	"fmt"
	"github.com/CocaineCong/gin-mall/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 根据username 判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("user_name=?", userName).Find(&user).Count(&count).Error
	fmt.Println(user, err)
	if count == 0 {
		return nil, false, err
	}
	return user, true, nil
}

func (dao *UserDao) CreateUser(user *model.User) error {
	return dao.DB.Model(&model.User{}).Create(&user).Error //表示为创建一个user类型的对象并返回err
}

// GetUserById 根据id获取user
func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	//对于不同的业务，其返回值也会有所不同 {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).Find(&user).Error
	//然后，调用 .Find(&user) 方法执行查询操作，并将结果存储在 user 变量中。
	//注意，这里传递的是指向 model.User 对象的指针，以便在查询成功后将结果填充到该对象中。
	return
}

// UpdateUserById 根据id更新user信息
// 在此方法中，dao 是 UserDao 类型的接收器，表示该方法是 UserDao 结构体的一个成员方法。uid 是要更新的用户的ID，user 是包含更新后的用户信息的 model.User 类型的指针。
func (dao *UserDao) UpdateUserById(uid uint, user *model.User) error {
	return dao.DB.Model(&model.User{}).Where("id=?", uid).Updates(&user).Error
}
