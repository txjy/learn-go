package v1

import (
	"gin_mall_tmp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProductImg(c *gin.Context) {
	ListCarousel := service.CarouselService{}
	if err := c.ShouldBind(&ListCarousel); err == nil {
		res := ListCarousel.List()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
