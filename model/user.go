package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string
	Money          string
}

/*
总之，这段代码定义了一个用户模型，并使用GORM库将其映射到数据库中的表。
通过GORM，你可以使用这个模型来进行数据库操作，如创建用户、查询用户、更新用户信息等。
主键字段（ID）：gorm.Model中的ID字段会自动映射到数据库表的主键字段，并具有自动生成的唯一标识符。这样可以方便地对数据进行唯一标识和索引。
创建时间字段（CreatedAt）：gorm.Model中的CreatedAt字段会自动记录每条数据的创建时间。这在跟踪数据创建和排序时非常有用。
更新时间字段（UpdatedAt）：gorm.Model中的UpdatedAt字段会自动记录每条数据的更新时间。这可以方便地跟踪数据的最后更新时间。
软删除字段（DeletedAt）：gorm.Model中的DeletedAt字段启用了软删除功能。当删除数据时，GORM会将DeletedAt字段设置为非空值，而不是直接从数据库中删除数据。这样可以保留数据的历史记录，并且可以很容易地恢复或永久删除数据
*/

const (
	PasswordCost        = 12       //密码加密等级
	Active       string = "active" //激活用户
)

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return err
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
