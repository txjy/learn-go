package v1

import (
	"gin_mall_tmp/pkg/e/util"

	service2 "gin_mall_tmp/service"

	"github.com/gin-gonic/gin"
)

//创建收藏
func CreateFavorite(c *gin.Context) {
	service := service2.FavoritesService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

//收藏夹详情接口
func ShowFavorites(c *gin.Context) {
	service := service2.FavoritesService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}

func DeleteFavorite(c *gin.Context) {
	service := service2.FavoritesService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
