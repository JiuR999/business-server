package controller

import (
	service2 "BusinessServer/apps/system/log/service"
	"BusinessServer/apps/system/role/service"
	common2 "BusinessServer/common"
	ac "BusinessServer/common/abstract/controller"
	common "BusinessServer/common/abstract/models"
	as "BusinessServer/common/abstract/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type roleController struct {
	ac.Controller
}

var lc = new(roleController)

//goland:noinspection ALL
func GetRoleController() *roleController {
	return lc
}
func getService() as.Service {
	return service.GetRoleService()
}

// @title			分页查询权限信息
// @version		1.0
// @Tags			SystemRole-系统权限管理相关接口
// @description	分页查询权限信息
// @accept			json
// @Produce		json
//
// @Param			token	header		string	false	"权限凭证"
// @Success		200		{object}	common.PageResponseModel
// @router			/api/system/role/page [Post]
func (api *roleController) PageApi(context *gin.Context) {
	api.Page(context, getService(), &common.PageModel{})
}

// @title			分页查询权限信息
// @version		1.0
// @Tags			SystemRole-系统权限管理相关接口
// @description	分页查询权限信息
// @accept			json
// @Produce		json
// @Param			userId	path		int		true	"用户ID"
// @Param			token	header		string	false	"权限凭证"
// @Success		200		{object}	[]models.SwustSystemRole
// @router			/api/system/role/getRolesByUserId/{userId} [Get]
func (api *roleController) GetRolesByUserId(context *gin.Context) {
	response := common.NewResponse(context)
	roles, swustError := service.GetRoleService().GetRolesByUserId(context.Param("userId"))
	if swustError != nil {
		response.ErrorWithMsg(swustError.GetMsg())
		return
	}
	response.SuccessWithData(roles)
}

// @title			分配权限
// @version		1.0
// @Tags			SystemRole-系统权限管理相关接口
// @description	分配权限
// @accept			json
// @Produce		json
// @Param			userId	path		int			true	"用户Id"
// @Param			token	header		string		false	"权限凭证"
// @Param			roles	body		[]string	false	"权限IDs"
// @Success		200		{object}	common.ResponseModel
// @router			/api/system/role/modify/{userId} [Post]
func (api *roleController) ModifyRole(context *gin.Context) {
	response := common.NewResponse(context)
	userId := context.Param("userId")

	if userId == "" {
		response.ErrorWithMsg("用户ID为空!")
		return
	}
	var roles []string
	if err := context.ShouldBindJSON(&roles); err != nil {
		response.ErrorWithMsg("参数错误!")
		return
	}
	err := service.GetRoleService().ModifyRole(userId, roles)
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	service2.WriteLog(context, common2.LOG_EVENT_UPDATE, fmt.Sprintf("修改用户%s的权限为%s", userId, roles))
	response.Success()
}
