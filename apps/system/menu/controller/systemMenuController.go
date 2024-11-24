package controller

import (
	"BusinessServer/apps/system/menu/service"
	ac "BusinessServer/common/abstract/controller"
	"BusinessServer/common/abstract/models"
	as "BusinessServer/common/abstract/service"
	"github.com/gin-gonic/gin"
)

type systemMenuController struct {
	ac.Controller
}

var sc = new(systemMenuController)

func GetSystemMenuController() *systemMenuController {
	return sc
}

func getService() as.Service {
	return service.GetSystemMenuService()
}

// @title			获取菜单树
// @version		1.0
// @Tags			SystemMenu-系统菜单管理相关接口
// @description	获取菜单树
// @accept			json
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Success		200		{object}	common.ResponseModel
// @router			/api/system/menu/list [Get]
func (api *systemMenuController) ListApi(context *gin.Context) {
	response := common.NewResponse(context)
	list, swustError := service.GetSystemMenuService().List()
	if swustError != nil {
		response.ErrorWithMsg(swustError.GetMsg())
		return
	}
	response.SuccessWithData(list)
}
