package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model        //ToDo: 这里的Model使用了`gorm:"primarykey"`作为主键的设置
	UserID     uint   `gorm:"not null"`
	Name       string `gorm:"type:varchar(20) not null"`
	Phone      string `gorm:"type:varchar(11) not null"`
	Address    string `gorm:"type:varchar(50) not null"`
}
