package v1

import (
	"github.com/CocaineCong/gin-mall/pkg/util"
	"github.com/CocaineCong/gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateAddress 创建商品
func CreateAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createAddressService := service.AddressService{}
	if err := c.ShouldBind(&createAddressService); err == nil {
		res := createAddressService.Create(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ListAddress 获取商品列表
func ListAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	listAddressService := service.AddressService{}
	if err := c.ShouldBind(&listAddressService); err == nil {
		res := listAddressService.List(c.Request.Context(), claim.ID)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// DeleteAddress 删除商品详细信息
func DeleteAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	showAddressService := service.AddressService{}
	if err := c.ShouldBind(&showAddressService); err == nil {
		res := showAddressService.Delete(c.Request.Context(), claim.ID, c.Param("id")) //c.Param("id")用于获取路径上的id,其返回值为string字符串类型.
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func ShowAddress(c *gin.Context) {
	showAddressService := service.AddressService{}
	if err := c.ShouldBind(&showAddressService); err == nil {
		res := showAddressService.Show(c.Request.Context(), c.Param("id")) //c.Param("id")用于获取路径上的id,其返回值为string字符串类型.
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func UpdateAddress(c *gin.Context) {
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	updateAddressService := service.AddressService{}
	if err := c.ShouldBind(&updateAddressService); err == nil {
		res := updateAddressService.UpDate(c.Request.Context(), claim.ID, c.Param("id")) //c.Param("id")用于获取路径上的id,其返回值为string字符串类型.
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}

}
