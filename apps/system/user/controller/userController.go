package controller

import (
	"BusinessServer/apps/system/user/models"
	"BusinessServer/apps/system/user/service"
	"BusinessServer/common"
	controller "BusinessServer/common/abstract/controller"
	common2 "BusinessServer/common/abstract/models"
	as "BusinessServer/common/abstract/service"
	"github.com/gin-gonic/gin"
	"strings"
)

type userController struct {
	controller.Controller
}

var uc = new(userController)

func GetUserController() *userController {
	return uc
}
func getService() as.Service {
	return service.GetUserService()
}

func newModel() *models.SystemUserModel {
	return &models.SystemUserModel{}
}

// @title			根据用户ID获取用户信息
// @version		1.0
// @Tags			SystemUser-系统用户管理相关接口
// @description	根据用户ID获取用户信息
// @Produce		json
//
// @Param			token	header		string	false	"用户凭证"
//
// @Param			id		query		string	false	"资产ID"
// @Success		200		{object}	common.ResponseModel
// @router			/api/system/user/getById [Get]
func (api *userController) GetApi(context *gin.Context) {
	api.GetById(context, getService(), newModel())
}

// @title			增加用户
// @version		1.0
// @Tags			SystemUser-系统用户管理相关接口
// @description	增加用户
// @Produce		json
//
// @Param			token	header		string					false	"用户凭证"
//
// @Param			body	body		models.SystemUserModel	false	"待添加用户"
// @Success		200		{object}	common.ResponseModel
// @router			/api/system/user/add [Post]
func (api *userController) AddApi(context *gin.Context) {
	api.Add(context, getService(), newModel())
}

// @title			根据用户IDs删除用户信息
// @version		1.0
// @Tags			SystemUser-系统用户管理相关接口
// @description	根据资产IDs删除资产信息
// @accept			json
// @Produce		json
//
// @Param			token	header		string		false	"用户凭证"
//
// @Param			ids		body		[]string	false	"待删除用户IDs"
// @Success		200		{object}	common.ResponseModel
// @router			/api/system/user/delete [Post]
func (api *userController) DeleteApi(context *gin.Context) {
	api.Delete(context, getService())
}

// @title			更新用户信息
// @version		1.0
// @Tags			SystemUser-系统用户管理相关接口
// @description	更新用户信息
// @Produce		json
//
// @Param			token	header		string					false	"用户凭证"
//
// @Param			body	body		models.SystemUserModel	false	"待添加资产"
// @Success		200		{object}	common.ResponseModel
// @router			/api/system/user/update [Post]
func (api *userController) UpdateApi(context *gin.Context) {
	api.Update(context, getService(), newModel())
}

// @title			分页查询用户信息
// @version		1.0
// @Tags			SystemUser-系统用户管理相关接口
// @description	分页查询用户信息
// @accept			json
// @Produce		json
//
// @Param			token	header		string							false	"用户凭证"
//
// @Param			body	body		models.SystemUserQueryRequest	false	"待删除资产IDs"
// @Success		200		{object}	common.PageResponseModel
// @router			/api/system/user/page [Post]
func (api *userController) PageApi(context *gin.Context) {
	api.Page(context, getService(), &common2.PageModel{})
}

// @title			登录
// @version		1.0
// @Tags			SystemUser-系统用户管理相关接口
// @description	登录
// @Produce		json
// @Param			body	body		models.LoginRequest	false	"账号密码"
// @Success		200		{object}	common.ResponseModel
// @router			/api/system/user/login [Post]
func (api *userController) Login(context *gin.Context) {
	response := common2.NewResponse(context)
	var req models.LoginRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(err.Error())
		return
	}
	if req.Account == "" || req.Password == "" {
		response.ErrorWithMsg("账号密码不能为空")
	}
	res, err := service.GetUserService().Login(req)
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithData(res)
}

// @title			获取当前登录用户信息
// @version		1.0
// @Tags			SystemUser-系统用户管理相关接口
// @description	获取当前登录用户信息
// @accept			json
// @Produce		json
//
// @Param			token	header		string	false	"用户凭证"
//
// @Success		200		{object}	models.SystemUserModel
// @router			/api/system/user/getCurrentUser [Get]
func (api *userController) GetCurrentUser(context *gin.Context) {
	response := common2.NewResponse(context)
	currentUser, err := service.GetUserService().GetCurrentUser(context)
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithData(currentUser)
}

// @title			退出登录
// @version		1.0
// @Tags			SystemUser-系统用户管理相关接口
// @description	退出登录
// @Produce		json
//
// @Param			token	header		string	false	"用户凭证"
//
// @Success		200		{object}	common.ResponseModel
// @router			/api/system/user/logout [Post]
func (api *userController) Logout(context *gin.Context) {
	response := common2.NewResponse(context)
	token := context.GetHeader("token")
	if strings.TrimSpace(token) == "" {
		response.ErrorWithMsg("用户未登录!")
		return
	}
	common.TokenMap[token] = token
	response.Success()
}
