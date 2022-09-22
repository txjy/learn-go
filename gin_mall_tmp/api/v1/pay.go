package v1

import (
	"gin_mall_tmp/pkg/e/util"

	"github.com/gin-gonic/gin"
)

func OrderPay(c *gin.Context) {
	orderPay := service.OrderPay{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&orderPay); err == nil {
		res := orderPay.PayDown(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		util.LogrusObj.Infoln(err)
		c.JSON(400, ErrorResponse(err))
	}
}
