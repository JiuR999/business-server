package controller

import (
	"BusinessServer/apps/system/log/models"
	"BusinessServer/apps/system/log/service"
	ac "BusinessServer/common/abstract/controller"
	as "BusinessServer/common/abstract/service"
	"github.com/gin-gonic/gin"
)

type logController struct {
	ac.Controller
}

var lc = new(logController)

//goland:noinspection ALL
func GetLogController() *logController {
	return lc
}
func getService() as.Service {
	return service.GetLogService()
}

func newModel() *models.SwustSystemLog {
	return &models.SwustSystemLog{}
}

// @title			根据日志ID获取日志信息
// @version		1.0
// @Tags			SystemLog-系统日志管理相关接口
// @description	根据日志ID获取日志信息
// @Produce		json
//
// @Param			token	header		string	false	"日志凭证"
//
// @Param			id		query		string	false	"资产ID"
// @Success		200		{object}	common.ResponseModel
// @router			/api/system/log/getById [Get]
func (api *logController) GetApi(context *gin.Context) {
	api.GetById(context, getService(), newModel())
}

/*// @title			增加日志
// @version		1.0
// @Tags			SystemLog-系统日志管理相关接口
// @description	增加日志
// @Produce		json
//
// @Param			token	header		string					false	"日志凭证"
//
// @Param			body	body		models.SwustSystemLog	false	"待添加日志"
// @Success		200		{object}	common.ResponseModel
// @router			/api/system/log/add [Post]
func (api *logController) AddApi(context *gin.Context) {
	api.Add(context, getService(), newModel())
}*/

// @title			根据日志IDs删除日志信息
// @version		1.0
// @Tags			SystemLog-系统日志管理相关接口
// @description	根据资产IDs删除资产信息
// @accept			json
// @Produce		json
// @Param			token	header		string		false	"日志凭证"
// @Param			ids		body		[]string	false	"待删除日志IDs"
// @Success		200		{object}	common.ResponseModel
// @router			/api/system/log/delete [Post]
func (api *logController) DeleteApi(context *gin.Context) {
	api.Delete(context, getService())
}

// @title			分页查询日志信息
// @version		1.0
// @Tags			SystemLog-系统日志管理相关接口
// @description	分页查询日志信息
// @accept			json
// @Produce		json
//
// @Param			token	header		string							false	"日志凭证"
//
// @Param			body	body		models.SystemLogQueryRequest	false	"待删除资产IDs"
// @Success		200		{object}	common.PageResponseModel
// @router			/api/system/log/page [Post]
func (api *logController) PageApi(context *gin.Context) {
	api.Page(context, getService(), &models.SystemLogQueryRequest{})
}
