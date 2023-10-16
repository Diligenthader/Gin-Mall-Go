package service

import (
	"context"
	"github.com/CocaineCong/gin-mall/dao"
	"github.com/CocaineCong/gin-mall/model"
	"github.com/CocaineCong/gin-mall/pkg/e"
	"github.com/CocaineCong/gin-mall/pkg/util"
	"github.com/CocaineCong/gin-mall/serializer"
	"strconv"
)

type CartService struct {
	Id        uint `json:"id" form:"id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	Num       int  `json:"num" form:"num"`
}

func (service *CartService) Create(ctx context.Context, uId uint) serializer.Response {
	var cart *model.Cart
	code := e.Success
	// 判断有没有这个商品
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	cartDao := dao.NewCartDao(ctx)
	cart = &model.Cart{ //表示为Cart这个指针指向model.Cart这个结构体.
		UserId:    uId,
		ProductId: service.ProductId,
		BossId:    service.BossId,
		Num:       uint(service.Num),
	}

	err = cartDao.CreateCart(cart)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetUserById(service.BossId)
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
		Data:   serializer.BuildCart(cart, product, boss),
	}
}

func (service *CartService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	carts, err := cartDao.ListCartByUserId(uId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	/*	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarts(ctx, carts),
	}*/
	return serializer.BuildListResponse(serializer.BuildCarts(ctx, carts), uint(len(carts)))
}

func (service *CartService) UpDate(ctx context.Context, uId uint, cId string) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	cartId, _ := strconv.Atoi(cId)
	err := cartDao.UpdateCartNumById(uint(cartId), service.Num, uId)
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

func (service *CartService) Delete(ctx context.Context, uId uint, aId string) serializer.Response {
	code := e.Success
	cartDao := dao.NewCartDao(ctx)
	cartId, _ := strconv.Atoi(aId)
	err := cartDao.DeleteCartByCartId(uint(cartId), uId) //表示为Delete操作根据购物车id和用户id进行删除操作.
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
