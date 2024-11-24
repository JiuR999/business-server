package controller

import (
	"BusinessServer/apps/producer/models"
	"BusinessServer/apps/producer/service"
	controller "BusinessServer/common/abstract/controller"
	as "BusinessServer/common/abstract/service"
	"github.com/gin-gonic/gin"
)

type producerController struct {
	controller.Controller
}

var rc = new(producerController)

func GetProducerController() *producerController {
	return rc
}
func getService() as.Service {
	return service.GetProducerService()
}

func newModel() *models.SwustProducer {
	return &models.SwustProducer{}
}

// @title			根据供销商ID获取供销商信息
// @version		1.0
// @Tags			Producer-供销商管理相关接口
// @description	根据供销商ID获取供销商信息
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Param			id		query		string	false	"资产ID"
// @Success		200		{object}	common.ResponseModel
// @router			/api/producer/getById [Get]
func (api *producerController) GetApi(context *gin.Context) {
	api.GetById(context, getService(), newModel())
}

// @title			增加供销商
// @version		1.0
// @Tags			Producer-供销商管理相关接口
// @description	增加供销商
// @Produce		json
// @Param			token	header		string					false	"用户凭证"
// @Param			body	body		models.SwustProducer	false	"待添加厂商ID"
// @Success		200		{object}	common.ResponseModel
// @router			/api/producer/add [Post]
func (api *producerController) AddApi(context *gin.Context) {
	api.Add(context, getService(), newModel())
}

// @title			根据供销商IDs删除供销商信息
// @version		1.0
// @Tags			Producer-供销商管理相关接口
// @description	根据供销商IDs删除供销商信息
// @accept			json
// @Produce		json
// @Param			token	header		string		false	"用户凭证"
// @Param			ids		body		[]string	false	"待删除供销商IDs"
// @Success		200		{object}	common.ResponseModel
// @router			/api/producer/delete [Post]
func (api *producerController) DeleteApi(context *gin.Context) {
	api.Delete(context, getService())
}

// @title			更新供销商信息
// @version		1.0
// @Tags			Producer-供销商管理相关接口
// @description	更新供销商信息
// @Produce		json
// @Param			token	header		string					false	"用户凭证"
// @Param			body	body		models.SwustProducer	false	"待添加供销商"
// @Success		200		{object}	common.ResponseModel
// @router			/api/producer/update [Post]
func (api *producerController) UpdateApi(context *gin.Context) {
	api.Update(context, getService(), newModel())
}

// @title			分页查询供销商信息
// @version		1.0
// @Tags			Producer-供销商管理相关接口
// @description	分页查询供销商信息
// @accept			json
// @Produce		json
// @Param			token	header		string						false	"用户凭证"
// @Param			req		body		models.ProducerQueryRequest	false	"筛选条件"
// @Success		200		{object}	common.PageResponseModel
// @router			/api/producer/page [Post]
func (api *producerController) PageApi(context *gin.Context) {
	api.Page(context, getService(), &models.ProducerQueryRequest{})
}
