package controller

import (
	"BusinessServer/apps/room/models"
	"BusinessServer/apps/room/service"
	controller "BusinessServer/common/abstract/controller"
	common2 "BusinessServer/common/abstract/models"
	as "BusinessServer/common/abstract/service"
	"github.com/gin-gonic/gin"
)

type roomController struct {
	controller.Controller
}

var rc = new(roomController)

func GetRoomController() *roomController {
	return rc
}
func getService() as.Service {
	return service.GetRoomService()
}

func newModel() *models.SwustRoomModel {
	return &models.SwustRoomModel{}
}

// @title			根据实验室ID获取实验室信息
// @version		1.0
// @Tags			Room-实验室管理相关接口
// @description	根据实验室ID获取实验室信息
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Param			id		query		string	false	"资产ID"
// @Success		200		{object}	models.SwustRoomModel
// @router			/api/room/getById [Get]
func (api *roomController) GetApi(context *gin.Context) {
	api.GetById(context, getService(), newModel())
}

// @title			增加实验室
// @version		1.0
// @Tags			Room-实验室管理相关接口
// @description	增加实验室
// @Produce		json
// @Param			token	header		string					false	"用户凭证"
// @Param			body	body		models.SwustRoomModel	false	"待添加实验室"
// @Success		200		{object}	common.ResponseModel
// @router			/api/room/add [Post]
func (api *roomController) AddApi(context *gin.Context) {
	api.Add(context, getService(), newModel())
}

// @title			根据实验室IDs删除实验室信息
// @version		1.0
// @Tags			Room-实验室管理相关接口
// @description	根据实验室IDs删除实验室信息
// @accept			json
// @Produce		json
// @Param			token	header		string		false	"用户凭证"
// @Param			ids		body		[]string	false	"待删除资产IDs"
// @Success		200		{object}	common.ResponseModel
// @router			/api/room/delete [Post]
func (api *roomController) DeleteApi(context *gin.Context) {
	api.Delete(context, getService())
}

// @title			更新实验室信息
// @version		1.0
// @Tags			Room-实验室管理相关接口
// @description	更新实验室信息
// @Produce		json
// @Param			token	header		string					false	"用户凭证"
// @Param			body	body		models.SwustRoomModel	false	"待添加资产"
// @Success		200		{object}	common.ResponseModel
// @router			/api/room/update [Post]
func (api *roomController) UpdateApi(context *gin.Context) {
	api.Update(context, getService(), newModel())
}

// @title			分页查询实验室信息
// @version		1.0
// @Tags			Room-实验室管理相关接口
// @description	分页查询实验室信息
// @accept			json
// @Produce		json
// @Param			token	header		string						false	"用户凭证"
// @Param			req		body		models.AssetsQueryRequest	false	"待删除资产IDs"
// @Success		200		{object}	common.PageResponseModel
// @router			/api/room/page [Post]
func (api *roomController) PageApi(context *gin.Context) {
	var req common2.PageModel
	if err := context.ShouldBindJSON(&req); err != nil {
		common2.NewResponse(context).ErrorWithMsg(err.Error())
	}
	api.Page(context, getService(), &req)
}

// @title			获取所有实验室楼栋信息
// @version		1.0
// @Tags			Room-实验室管理相关接口
// @description	获取所有实验室楼栋信息
// @accept			json
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Success		200		{object}	models.SwustRoomModel
// @router			/api/room/listLocation [Get]
func (api *roomController) ListLocation(context *gin.Context) {
	response := common2.NewResponse(context)
	res, err := service.GetRoomService().ListLocation()
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithData(res)
}
