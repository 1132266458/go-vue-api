package routes

import (
	"github.com/gin-gonic/gin"
	v1 "go-vue-api/api/v1"
	"go-vue-api/middleware"
	"go-vue-api/utils"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Log())

	router := r.Group("api")
	{
		//用户相关
		router.POST("/user/add",v1.AddUser)

		//第三方调用
		router.GET("/message/send",v1.SendMessage)
	}

	_ = r.Run(utils.HttpPort)
}
