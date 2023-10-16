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

type FavoritesService struct {
	ProductId  uint `json:"product_id" form:"product_id"`
	BossId     uint `json:"boss_id" form:"boss_id"`
	FavoriteId uint `json:"favorite_id" form:"favorite_id"`
	model.BasePage
}

func (service *FavoritesService) List(ctx context.Context, uId uint) serializer.Response {
	favoriteDao := dao.NewFavoritesDao(ctx)
	code := e.Success
	favorite, err := favoriteDao.ListFavorites(uId)
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorite), uint(len(favorite)))
}

func (service *FavoritesService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	favoriteDao := dao.NewFavoritesDao(ctx)
	exist, _ := favoriteDao.FavoriteExistOrNot(service.ProductId, uId)
	if exist {
		code = e.ErrorFavoriteExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	bossDao := dao.NewUserDao(ctx)
	boss, err := bossDao.GetUserById(service.BossId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(service.ProductId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	favorite := &model.Favorite{
		User:      *user,
		UserId:    uId,
		Product:   *product,
		ProductId: service.ProductId,
		Boss:      *boss,
		BossId:    service.BossId,
	}

	err = favoriteDao.CreateFavorite(favorite)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *FavoritesService) Delete(ctx context.Context, uId uint, fId string) serializer.Response {
	favoriteDao := dao.NewFavoritesDao(ctx)
	code := e.Success
	favoriteId, _ := strconv.Atoi(fId)
	err := favoriteDao.DeleteFavorite(uId, uint(favoriteId))
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
