package router

import (
	userController "BusinessServer/apps/system/user/controller"
	"BusinessServer/common/Websocket"
	"BusinessServer/common/middleware"
	_ "BusinessServer/docs"
	"BusinessServer/env"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	r := gin.Default()
	r.Use(middleware.Recovery())
	api := r.Group("/api")
	api.POST("/system/user/login", userController.GetUserController().Login)
	if env.GetConfig().EnableAuth {
		api.Use(middleware.JwtHandler())
	}
	assetsApi(api.Group("/assets"))
	menuApi(api.Group("/system/menu"))
	userApi(api.Group("/system/user"))
	statisticApi(api.Group("/statistic"))
	roomApi(api.Group("/room"))
	producerApi(api.Group("/producer"))
	logApi(api.Group("/system/log"))
	roleApi(api.Group("/system/role"))
	orderApi(api.Group("/order"))
	//Websocket
	r.GET("/ws/common", Websocket.HandFunc)
	r.GET("/ws/notice", Websocket.HandNoticeFunc)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":" + env.GetConfig().Port)
	if err != nil {
		fmt.Println("服务器端失败:", err.Error())
	}
}
