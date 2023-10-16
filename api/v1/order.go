package v1

import (
	"github.com/CocaineCong/gin-mall/pkg/util"
	"github.com/CocaineCong/gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateOrder 创建购物车
func CreateOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createOrderService := service.OrderService{}
	if err := c.ShouldBind(&createOrderService); err == nil {
		res := createOrderService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// DeleteOrder 删除购物车
func DeleteOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	deleteOrderService := service.OrderService{}
	if err := c.ShouldBind(&deleteOrderService); err == nil {
		res := deleteOrderService.Delete(c.Request.Context(), claim.ID, c.Param("id")) //c.Param("id")用于获取路径上的id,其返回值为string字符串类型.
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ListOrder 展示购物车
func ListOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	listOrderService := service.OrderService{}
	if err := c.ShouldBind(&listOrderService); err == nil {
		res := listOrderService.List(c.Request.Context(), claim.ID) //c.Param("id")用于获取路径上的id,其返回值为string字符串类型.
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ShowOrder 更新购物车商品数量
func ShowOrder(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showOrderService := service.OrderService{}
	if err := c.ShouldBind(&showOrderService); err == nil {
		res := showOrderService.Show(c.Request.Context(), claim.ID, c.Param("id")) //c.Param("id")用于获取路径上的id,其返回值为string字符串类型.
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}
