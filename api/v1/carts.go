package v1

import (
	"github.com/CocaineCong/gin-mall/pkg/util"
	"github.com/CocaineCong/gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateCart 创建购物车
func CreateCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createCartService := service.CartService{}
	if err := c.ShouldBind(&createCartService); err == nil {
		res := createCartService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// DeleteCart 删除购物车
func DeleteCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showCartService := service.CartService{}
	if err := c.ShouldBind(&showCartService); err == nil {
		res := showCartService.Delete(c.Request.Context(), claim.ID, c.Param("id")) //c.Param("id")用于获取路径上的id,其返回值为string字符串类型.
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ListCart 展示购物车
func ListCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showCartService := service.CartService{}
	if err := c.ShouldBind(&showCartService); err == nil {
		res := showCartService.List(c.Request.Context(), claim.ID) //c.Param("id")用于获取路径上的id,其返回值为string字符串类型.
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// UpdateCart 更新购物车商品数量
func UpdateCart(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	updateCartService := service.CartService{}
	if err := c.ShouldBind(&updateCartService); err == nil {
		res := updateCartService.UpDate(c.Request.Context(), claim.ID, c.Param("id")) //c.Param("id")用于获取路径上的id,其返回值为string字符串类型.
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}
