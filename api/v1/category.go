package v1

import (
	"github.com/CocaineCong/gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListCategory 获取商品分类
func ListCategory(c *gin.Context) {
	var listCategory service.CategoryService
	if err := c.ShouldBind(&listCategory); err == nil {
		res := listCategory.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
