package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserId"`
	UserId    uint    `gorm:"not null"`
	Product   Product `gorm:"ForeignKey:ProductId"`
	ProductId uint    `gorm:"not null"`
	Boss      User    `gorm:"ForeignKey:BossId"`
	BossId    uint    `gorm:"not null"`
}

//通过数据库表结构关系图可知	Boss是User的一个对象，其通过外键关联到Favorite

/*
gorm.Model：这个字段是一个内嵌字段（Embedded Field），它可以为模型提供一些公共字段，如 ID、CreatedAt、UpdatedAt 和 DeletedAt，用于跟踪记录的创建、更新和软删除。通过嵌入 gorm.Model，Favorite 结构体也会继承这些字段。
User：这个字段表示一个用户（User）对象，并使用 ForeignKey:UserId 标签指定了外键关系。这意味着 Favorite 模型将与 User 模型建立关联，并使用 UserId 作为外键。
UserId：这个字段是一个无符号整数（uint），它存储与 Favorite 相关联的用户（User）的ID。gorm:"not null" 标签指定了该字段不能为空。
Product：这个字段表示一个产品（Product）对象，并使用 ForeignKey:ProductId 标签指定了外键关系。这意味着 Favorite 模型将与 Product 模型建立关联，并使用 ProductId 作为外键。
ProductId：这个字段是一个无符号整数（uint），它存储与 Favorite 相关联的产品（Product）的ID。gorm:"not null" 标签指定了该字段不能为空。
Boss：这个字段表示一个老板（Boss）对象，并使用 ForeignKey:BossId 标签指定了外键关系。这意味着 Favorite 模型将与 User 模型建立关联，并使用 BossId 作为外键。
BossId：这个字段是一个无符号整数（uint），它存储与 Favorite 相关联的老板（Boss）的ID。gorm:"not null" 标签指定了该字段不能为空。
通过以上字段定义，Favorite 模型将与 User 模型和 Product 模型建立关联，并且使用 UserId、ProductId 和 BossId 作为外键。这样的关联关系可以让你在进行数据库操作时，方便地检索、查询和操作相关的用户、产品和老板信息。
*/
