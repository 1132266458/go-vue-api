package v1

import (
	"github.com/gin-gonic/gin"
	"go-vue-api/thirdcall"
	"go-vue-api/utils/errmsg"
)

/**
	发送短信
 */
func SendMessage(c *gin.Context)  {
	phone := c.Query("phone")
	code := thirdcall.Send(phone)

	c.JSON(
		200,gin.H{
			"status":"200",
			"message": errmsg.GetErrMsg(code),
		},
	)

}