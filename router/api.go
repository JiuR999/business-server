package router

import (
	assetController "BusinessServer/apps/assets/controller"
	assetTypeController "BusinessServer/apps/assets/type/controller"
	orderController "BusinessServer/apps/order/controller"
	producerController "BusinessServer/apps/producer/controller"
	roomController "BusinessServer/apps/room/controller"
	"BusinessServer/apps/statistic/controller"
	logController "BusinessServer/apps/system/log/controller"
	menuController "BusinessServer/apps/system/menu/controller"
	roleController "BusinessServer/apps/system/role/controller"
	userController "BusinessServer/apps/system/user/controller"
	"BusinessServer/common/middleware"
	"BusinessServer/env"
	"github.com/gin-gonic/gin"
)

/*
*
设备资产APi
*/
func assetsApi(group *gin.RouterGroup) {
	if env.GetConfig().EnableAuth == true {
		group.Use(middleware.JwtHandler())
	}
	group.GET("/getById", assetController.GetAssetController().GetApi)
	group.POST("/add", assetController.GetAssetController().AddApi)
	group.POST("/page", assetController.GetAssetController().PageApi)
	group.POST("/delete", assetController.GetAssetController().DeleteApi)
	group.POST("/update", assetController.GetAssetController().UpdateApi)
	group.GET("/template", assetController.GetAssetController().Template)
	group.POST("/import", assetController.GetAssetController().Import)
	group.POST("/export", assetController.GetAssetController().Export)
	group.POST("/deprecate", assetController.GetAssetController().Deprecate)
	{
		group.POST("/type/page", assetTypeController.GetTypeController().PageApi)
		group.POST("/type/add", assetTypeController.GetTypeController().AddApi)
		group.GET("/type/getById", assetTypeController.GetTypeController().GetApi)
		group.POST("/type/delete", assetTypeController.GetTypeController().DeleteApi)
		group.POST("/type/update", assetTypeController.GetTypeController().UpdateApi)
	}
}

/*
*
系统菜单APi
*/
func menuApi(group *gin.RouterGroup) {
	if env.GetConfig().EnableAuth == true {
		group.Use(middleware.JwtHandler())
	}
	group.GET("/list", menuController.GetSystemMenuController().ListApi)
}

/*
*
用户APi
*/
func userApi(group *gin.RouterGroup) {
	group.GET("/getById", userController.GetUserController().GetApi)
	group.GET("/getCurrentUser", userController.GetUserController().GetCurrentUser)
	group.POST("/add", userController.GetUserController().AddApi)
	group.POST("/logout", userController.GetUserController().Logout)
	group.POST("/page", userController.GetUserController().PageApi)
}

/*
*
统计相关APi
*/
func statisticApi(group *gin.RouterGroup) {
	group.GET("/countAssetsByType", controller.GetStatisticController().CountAssetsByType)
	group.GET("/countAssetsByStatus", controller.GetStatisticController().CountAssetsByStatus)
	group.GET("/countAssetsApplyTrend", controller.GetStatisticController().CountAssetsApplyTrend)
	group.GET("/countAssetsDepTrend", controller.GetStatisticController().CountAssetsDepTrend)
	group.GET("/countOrderDetail", controller.GetStatisticController().CountOrderDetail)
}

// 实验室相关API
func roomApi(group *gin.RouterGroup) {
	group.POST("/page", roomController.GetRoomController().PageApi)
	group.GET("/listLocation", roomController.GetRoomController().ListLocation)
	group.GET("/getById", roomController.GetRoomController().GetApi)
	group.POST("/add", roomController.GetRoomController().AddApi)
	group.POST("/delete", roomController.GetRoomController().DeleteApi)
	group.POST("/update", roomController.GetRoomController().UpdateApi)
}

// 生产厂商相关API
func producerApi(group *gin.RouterGroup) {
	group.POST("/page", producerController.GetProducerController().PageApi)
	group.GET("/getById", producerController.GetProducerController().GetApi)
	group.POST("/add", producerController.GetProducerController().AddApi)
	group.POST("/delete", producerController.GetProducerController().DeleteApi)
	group.POST("/update", producerController.GetProducerController().UpdateApi)
}

// websocket
func wsApi(group *gin.RouterGroup) {
	group.GET("/websocket/common")
}

// 日志相关API
func logApi(group *gin.RouterGroup) {
	group.POST("/page", logController.GetLogController().PageApi)
	group.POST("/delete", logController.GetLogController().DeleteApi)
	group.GET("/getById", logController.GetLogController().GetApi)
}

// 权限相关API
func roleApi(group *gin.RouterGroup) {
	group.POST("/page", roleController.GetRoleController().PageApi)
	group.GET("/getRolesByUserId/:userId", roleController.GetRoleController().GetRolesByUserId)
	group.POST("/modify/:userId", roleController.GetRoleController().ModifyRole)
}

// 权限相关API
func orderApi(group *gin.RouterGroup) {
	group.POST("/page", orderController.GetOrderController().PageApi)
	//group.GET("/getRolesByUserId/:userId", roleController.GetRoleController().GetRolesByUserId)
	group.POST("/add", orderController.GetOrderController().AddApi)
	group.GET("/getById", orderController.GetOrderController().GetApi)
	group.POST("/approve", orderController.GetOrderController().Approve)
}
