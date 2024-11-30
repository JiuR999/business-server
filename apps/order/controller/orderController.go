package controller

import (
	"BusinessServer/apps/order/models"
	"BusinessServer/apps/order/service"
	logWriter "BusinessServer/apps/system/log/service"
	"BusinessServer/common"
	controller "BusinessServer/common/abstract/controller"
	common2 "BusinessServer/common/abstract/models"
	as "BusinessServer/common/abstract/service"
	"github.com/gin-gonic/gin"
)

type orderController struct {
	controller.Controller
}

var oc = new(orderController)

func GetOrderController() *orderController {
	return oc
}
func getService() as.Service {
	return service.GetOrderService()
}

func newModel() *models.SwustOrder {
	return &models.SwustOrder{}
}

// @title			根据订单ID获取采购信息
// @version		1.0
// @Tags			Order-采购信息管理相关接口
// @description	根据订单ID获取采购信息
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Param			id		query		string	false	"订单ID"
// @Success		200		{object}	common.ResponseModel
// @router			/api/order/getById [Get]
func (api *orderController) GetApi(context *gin.Context) {
	api.GetById(context, getService(), &models.OrderVO{})
}

// @title			增加采购信息
// @version		1.0
// @Tags			Order-采购信息管理相关接口
// @description	增加采购信息
// @Produce		json
// @Param			token	header		string				false	"用户凭证"
// @Param			body	body		models.OrderRequest	false	"采购单信息"
// @Success		200		{object}	common.ResponseModel
// @router			/api/order/add [Post]
func (api *orderController) AddApi(context *gin.Context) {
	api.Add(context, getService(), &models.OrderRequest{})
}

// @title			根据采购IDs删除采购信息
// @version		1.0
// @Tags			Order-采购信息管理相关接口
// @description	根据采购IDs删除采购信息
// @accept			json
// @Produce		json
// @Param			token	header		string		false	"用户凭证"
// @Param			ids		body		[]string	false	"待删除采购信息IDs"
// @Success		200		{object}	common.ResponseModel
// @router			/api/order/delete [Post]
func (api *orderController) DeleteApi(context *gin.Context) {
	api.Delete(context, getService())
}

// @title			更新采购信息
// @version		1.0
// @Tags			Order-采购信息管理相关接口
// @description	更新资产信息
// @Produce		json
// @Param			token	header		string				false	"用户凭证"
// @Param			body	body		models.OrderRequest	false	"待更新采购单信息"
// @Success		200		{object}	common.ResponseModel
// @router			/api/order/update [Post]
func (api *orderController) UpdateApi(context *gin.Context) {
	api.Update(context, getService(), newModel())
}

// @title			分页查询采购信息
// @version		1.0
// @Tags			Order-采购信息管理相关接口
// @description	分页查询资产信息
// @accept			json
// @Produce		json
// @Param			token	header		string						false	"用户凭证"
// @Param			req		body		models.OrderQueryRequest	false	"待删除资产IDs"
// @Success		200		{object}	common.PageResponseModel
// @router			/api/order/page [Post]
func (api *orderController) PageApi(context *gin.Context) {
	api.Page(context, getService(), &models.OrderQueryRequest{})
}

// @title			审批采购信息
// @version		1.0
// @Tags			Order-采购信息管理相关接口
// @description	审批采购信息
// @accept			json
// @Produce		json
// @Param			token	header		string						false	"用户凭证"
// @Param			req		body		models.ApproveReq	false	"待审批订单ID"
// @Success		200		{object}	common.ResponseModel
// @router			/api/order/approve [Post]
func (api *orderController) Approve(context *gin.Context) {
	var req models.ApproveReq
	response := common2.NewResponse(context)
	if err := context.ShouldBindJSON(&req); err != nil {
		response.ErrorWithMsg(err.Error())
		return
	}
	service.GetOrderService().Approve(req)
	//写入日志
	status := "同意"
	if req.Status == common.APPROVE_REFUSE {
		status = "拒绝"
	}
	logWriter.WriteLog(context, common.LOG_EVENT_OA, status+"采购申请")

	response.Success()
}
