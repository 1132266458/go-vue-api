package v1

import (
	"github.com/gin-gonic/gin"
	"go-vue-api/model"
	"go-vue-api/utils/errmsg"
	"go-vue-api/utils/validator"
	"net/http"
)

/**
 * 添加用户
 */
func AddUser(c *gin.Context)  {

	var data model.User
	var msg string
	var validCode int
	_ = c.ShouldBindJSON(&data)

	msg, validCode = validator.Validate(&data)
	if validCode != errmsg.SUCCSE {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
			},
		)
		c.Abort() //阻止调用后续的处理函数
		return
	}

	//判断用户是否已存在
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}

	c.JSON(
		200,gin.H{
				"status":"200",
				"message": errmsg.GetErrMsg(code),
			},
		)
}