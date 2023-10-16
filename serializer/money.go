package serializer

import (
	"github.com/CocaineCong/gin-mall/model"
	"github.com/CocaineCong/gin-mall/pkg/util"
)

type Money struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"` //form表单数据
}

/*
每个字段都使用 json 标签和 form 标签进行了注释，
这些标签可用于在序列化和反序列化 JSON 数据或在处理表单数据时进行字段映射。
*/

func BuildMoney(item *model.User, key string) Money {
	util.Encrypt.SetKey(key)
	return Money{
		UserId:    item.ID,
		UserName:  item.UserName,
		UserMoney: util.Encrypt.AssDecoding(item.Money),
	}
}
