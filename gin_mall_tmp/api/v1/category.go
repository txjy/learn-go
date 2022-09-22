package v1

import (
	"gin_mall_tmp/pkg/e/util"
	"gin_mall_tmp/service"

	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	var listCategory service.CategoryService{}
	if err := c.ShouldBind(&listCategory); err == nil {
		res := listCategory.List(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
