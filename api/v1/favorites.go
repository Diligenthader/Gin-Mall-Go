package v1

import (
	"github.com/CocaineCong/gin-mall/pkg/util"
	"github.com/CocaineCong/gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateFavorites 创建商品
func CreateFavorites(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createFavoritesService := service.FavoritesService{}
	if err := c.ShouldBind(&createFavoritesService); err == nil {
		res := createFavoritesService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ListFavorites 获取商品列表
func ListFavorites(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	listFavoritesService := service.FavoritesService{}
	if err := c.ShouldBind(&listFavoritesService); err == nil {
		res := listFavoritesService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// DeleteFavorites 删除商品详细信息
func DeleteFavorites(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showFavoritesService := service.FavoritesService{}
	if err := c.ShouldBind(&showFavoritesService); err == nil {
		res := showFavoritesService.Delete(c.Request.Context(), claim.ID, c.Param("id")) //c.Param("id")用于获取路径上的id,其返回值为string字符串类型.
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
