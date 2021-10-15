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
		router.POST("/user/add",v1.AddUser)
	}

	_ = r.Run(utils.HttpPort)
}
