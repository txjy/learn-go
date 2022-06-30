package v1

import (
	"go_bin_blog/model"
	"go_bin_blog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int

//添加用户
func AddUser(c *gin.Context) {
	//
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个用户

//查询用户列表
func GetUsers(c *gin.Context) {
	//
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	username :=c.Query("username")

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data,total := model.GetUsers(username,pageSize, pageNum)

	code = errmsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

//编辑用户
func UpdateUser(c *gin.Context) {
	//
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code = model.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		model.UpdateUser(id, &data)
	}

	
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

//删除用户
func DeleteUser(c *gin.Context) {
	//
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
