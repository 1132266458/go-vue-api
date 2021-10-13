package v1

import "github.com/gin-gonic/gin"

/**
 * 添加用户
 */
func AddUser(c *gin.Context)  {
	c.JSON(
		200,gin.H{
				"status":"200",
			},
		)
}