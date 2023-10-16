package service

import (
	"context"
	"github.com/CocaineCong/gin-mall/dao"
	"github.com/CocaineCong/gin-mall/serializer"
	"strconv"
)

type ListProductImg struct {
}

func (service *ListProductImg) List(ctx context.Context, pId string) serializer.Response {
	productImgDao := dao.NewProductImgDao(ctx)
	productId, _ := strconv.Atoi(pId)
	productImgs, _ := productImgDao.ListProductImg(uint(productId))
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
