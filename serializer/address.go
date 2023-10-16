package serializer

import "github.com/CocaineCong/gin-mall/model"

//ToDo 标准化形式的序列化构造，涉及单个以及切片构造等.

type Address struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"user_id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
}

func BuildAddress(item *model.Address) Address {
	return Address{
		Id:        item.ID,
		UserId:    item.UserID,
		Name:      item.Name,
		Phone:     item.Phone,
		Address:   item.Address,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

func BuildAddresses(items []*model.Address) (addresses []Address) {
	for _, item := range items {
		address := BuildAddress(item)
		addresses = append(addresses, address)
	}
	return
}
