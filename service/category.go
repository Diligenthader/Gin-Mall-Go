package service

import (
	"context"
	"github.com/CocaineCong/gin-mall/dao"
	"github.com/CocaineCong/gin-mall/pkg/e"
	"github.com/CocaineCong/gin-mall/pkg/util"
	"github.com/CocaineCong/gin-mall/serializer"
)

type CategoryService struct {
}

func (service *CategoryService) List(ctx context.Context) serializer.Response {
	categorydao := dao.NewCategoryDao(ctx)
	code := e.Success
	categories, err := categorydao.ListCategory()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCategories(categories), uint(len(categories)))
}
