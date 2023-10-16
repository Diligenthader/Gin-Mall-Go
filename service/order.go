package service

import (
	"context"
	"fmt"
	"github.com/CocaineCong/gin-mall/dao"
	"github.com/CocaineCong/gin-mall/model"
	"github.com/CocaineCong/gin-mall/pkg/e"
	"github.com/CocaineCong/gin-mall/pkg/util"
	"github.com/CocaineCong/gin-mall/serializer"
	"math/rand"
	"strconv"
	"time"
)

// ToDo :增加复用，提高效率，便于开发.

type OrderService struct {
	ProductId uint    `json:"product_id" form:"product_id"`
	Num       int     `json:"num" form:"num"`
	AddressId uint    `json:"address_id" form:"address_id"`
	Money     float64 `json:"money" form:"money"`
	BossId    uint    `json:"boss_id" form:"boss_id"`
	UserId    uint    `json:"user_id" form:"user_id"`
	OrderNum  int     `json:"order_num" form:"order_num"`
	Type      int     `json:"type" form:"type"` //用于判断所进行业务的类型.
	model.BasePage
}

func (service *OrderService) Create(ctx context.Context, uId uint) serializer.Response {
	var order *model.Order
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	order = &model.Order{ //表示为Order这个指针指向model.Order这个结构体.
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       service.Num,
		Money:     service.Money,
		Type:      1, //未支付 默认值为1
	}
	// 检验地址是否存在
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(service.AddressId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	order.AddressId = address.ID
	// 生成订单号,自动生成的随机number+唯一标识的product_id+用户的id
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000000))
	/*
		rand.NewSource(time.Now().UnixNano()) 创建了一个以当前时间作为种子的随机数生成器源。
		rand.New(...) 使用上述生成器源创建了一个新的随机数生成器。
		Int31n(10000000) 生成一个介于 0 到 9999999 之间的随机数。
		最后，使用 fmt.Sprintf 函数将该随机数格式化为字符串。格式化字符串 %09v 指示将该随机数转换为字符串，并在前面填充 0，使其总长度为 9 位。
		总之，该行代码的目的是生成一个 9 位长度的随机数字符串，不足 9 位的部分在前面填充 0。
	*/
	productNum := strconv.Itoa(int(service.ProductId))
	userNum := strconv.Itoa(int(service.UserId))
	number = number + productNum + userNum
	orderNum, _ := strconv.ParseUint(number, 10, 64)
	order.OrderNum = orderNum
	err = orderDao.CreateOrder(order)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *OrderService) Show(ctx context.Context, uId uint, oId string) serializer.Response {
	orderId, _ := strconv.Atoi(oId)
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	// 1.首先先拿到订单号
	order, err := orderDao.GetOrderById(uint(orderId), uId) //用于查询当前用户的订单号.
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 2.解析该订单的地址
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(order.AddressId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 3.解析该订单的货品号
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildOrder(order, product, address),
	}
}

func (service *OrderService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	//若当要使用列表对数据进行展示时，便可以进行分页查询的设置.
	if service.PageSize == 0 {
		service.PageSize = 15
	}

	orderDao := dao.NewOrderDao(ctx)
	condition := make(map[string]interface{})
	if service.Type != 0 {
		condition["type"] = service.Type
	}
	condition["user_id"] = uId
	orderList, err := orderDao.ListOrderByCondition(condition, service.BasePage)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildOrders(ctx, orderList), uint(len(orderList)))
}

func (service *OrderService) Delete(ctx context.Context, uId uint, aId string) serializer.Response {
	code := e.Success
	orderDao := dao.NewOrderDao(ctx)
	orderId, _ := strconv.Atoi(aId)
	err := orderDao.DeleteOrderByOrderId(uint(orderId), uId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
