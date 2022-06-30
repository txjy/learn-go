package v1

import (
	"go_bin_blog/middleware"
	"go_bin_blog/model"
	"go_bin_blog/utils/errmsg"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var data model.User
	var code int
	var token string
	c.ShouldBindJSON(&data)

	data, code = model.CheckLogin(data.UserName, data.PassWord)

	if code == errmsg.SUCCESS {
		setToken(c, data)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data.UserName,
			"id":      data.ID,
			"message": errmsg.GetErrMsg(code),
			"token":   token,
		})
	}
}

// token生成函数
func setToken(c *gin.Context, user model.User) {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: user.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "GinBlog",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    user.UserName,
		"id":      user.ID,
		"message": errmsg.GetErrMsg(200),
		"token":   token,
	})
	return
}
