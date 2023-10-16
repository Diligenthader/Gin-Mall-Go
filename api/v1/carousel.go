package v1

import (
	"github.com/CocaineCong/gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListCarousel(c *gin.Context) {
	var listCarousel service.CarouselService
	if err := c.ShouldBind(&listCarousel); err == nil {
		res := listCarousel.List(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
