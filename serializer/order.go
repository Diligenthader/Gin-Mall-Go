package serializer

import (
	"context"
	"github.com/CocaineCong/gin-mall/conf"
	"github.com/CocaineCong/gin-mall/dao"
	"github.com/CocaineCong/gin-mall/model"
)

type Order struct {
	Id           uint    `json:"id"`
	OrderNum     uint64  `json:"order_num"`
	CreatedAt    int64   `json:"created_at"`
	UpdatedAt    int64   `json:"updated_at"`
	UserId       uint    `json:"user_id"`
	ProductId    uint    `json:"product_id"`
	BossId       uint    `json:"boss_id"`
	Num          int     `json:"num"`
	AddressName  string  `json:"address_name"`
	AddressPhone string  `json:"address_phone"`
	Type         uint    `json:"type"`
	ProductName  string  `json:"product_name"`
	ImgPath      string  `json:"img_path"`
	Money        float64 `json:"discount_price"`
}

//ToDo :serializer用于返回给前端数据看，所以要提前定义好返回的格式.

func BuildOrder(order *model.Order, product *model.Product, address *model.Address) Order {
	return Order{
		UserId:       order.UserId,
		OrderNum:     order.OrderNum,
		BossId:       order.BossId, //ToDo 可以理解为这是订单的商家，即卖货的人
		ProductId:    order.ProductId,
		CreatedAt:    order.CreatedAt.Unix(),
		Type:         order.Type,
		ProductName:  product.Name,
		Money:        order.Money,
		Num:          product.Num,
		ImgPath:      conf.Host + conf.ProductPath + product.ImgPath,
		AddressName:  address.Name,
		AddressPhone: address.Phone,
	}
}

func BuildOrders(ctx context.Context, items []*model.Order) (orders []Order) {
	productDao := dao.NewProductDao(ctx)
	addressDao := dao.NewAddressDao(ctx)

	for _, item := range items {
		product, err := productDao.GetProductById(item.ProductId)
		if err != nil {
			continue
		}
		address, err := addressDao.GetAddressByAid(item.AddressId)
		if err != nil {
			continue
		}
		order := BuildOrder(item, product, address)
		orders = append(orders, order)
	}
	return orders
}
