package v1

import (
	"github.com/CocaineCong/gin-mall/pkg/util"
	"github.com/CocaineCong/gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	/*
		c.MultipartForm()：这部分代码从上下文对象c中获取多部分表单的数据。多部分表单是一种用于上传文件等数据的表单格式。
		form, _ := c.MultipartForm()：这一行代码将多部分表单数据赋值给变量form。使用下划线_来忽略可能返回的错误。
		files := form.File["file"]：这部分代码从多部分表单数据中提取名为"file"的文件字段。假设"file"是表单中的文件上传字段名称。form.File是一个映射，其中键是文件字段的名称，
		值是一个切片，包含了所有上传的文件。在这种情况下，files将是一个切片，其中包含了名为"file"的所有上传文件。
	*/
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	createProductService := service.ProductService{}
	if err := c.ShouldBind(&createProductService); err == nil {
		res := createProductService.Create(c.Request.Context(), claim.ID, files)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ListProduct 获取商品列表
func ListProduct(c *gin.Context) {
	listProductService := service.ProductService{}
	if err := c.ShouldBind(&listProductService); err == nil {
		res := listProductService.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// ShowProduct 展示商品详细信息
func ShowProduct(c *gin.Context) {
	showProductService := service.ProductService{}
	if err := c.ShouldBind(&showProductService); err == nil {
		res := showProductService.Show(c.Request.Context(), c.Param("id")) //c.Param("id")用于获取路径上的id,其返回值为string字符串类型.
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

// SearchProduct 查询商品
func SearchProduct(c *gin.Context) {
	searchProductService := service.ProductService{}
	if err := c.ShouldBind(&searchProductService); err == nil {
		res := searchProductService.Search(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
