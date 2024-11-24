package controller

import (
	"BusinessServer/apps/assets/type/models"
	"BusinessServer/apps/assets/type/service"
	controller "BusinessServer/common/abstract/controller"
	as "BusinessServer/common/abstract/service"
	"github.com/gin-gonic/gin"
)

type typeController struct {
	controller.Controller
}

var rc = new(typeController)

func GetTypeController() *typeController {
	return rc
}
func getService() as.Service {
	return service.GetTypeService()
}

func newModel() *models.SwustAssetType {
	return &models.SwustAssetType{}
}

// @title			根据资产类型ID获取资产信息
// @version		1.0
// @Tags			AssetType-资产类型管理相关接口
// @description	根据资产类型ID获取供资产信息
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Param			id		query		string	false	"资产ID"
// @Success		200		{object}	models.SwustAssetType
// @router			/api/assets/type/getById [Get]
func (api *typeController) GetApi(context *gin.Context) {
	api.GetById(context, getService(), newModel())
}

// @title			增加资产类型
// @version		1.0
// @Tags			AssetType-资产类型管理相关接口
// @description	增加资产类型
// @Produce		json
// @Param			token	header		string					false	"用户凭证"
// @Param			body	body		models.SwustAssetType	false	"待添加资产类型"
// @Success		200		{object}	common.ResponseModel
// @router			/api/assets/type/add [Post]
func (api *typeController) AddApi(context *gin.Context) {
	api.Add(context, getService(), newModel())
}

// @title			根据资产IDs删除实资产类型信息
// @version		1.0
// @Tags			AssetType-资产类型管理相关接口
// @description	根据资产IDs删除资产类型信息
// @accept			json
// @Produce		json
// @Param			token	header		string		false	"用户凭证"
// @Param			ids		body		[]string	false	"待删除资产IDs"
// @Success		200		{object}	common.ResponseModel
// @router			/api/assets/type/delete [Post]
func (api *typeController) DeleteApi(context *gin.Context) {
	api.Delete(context, getService())
}

// @title			更新资产类型信息
// @version		1.0
// @Tags			AssetType-资产类型管理相关接口
// @description	更新资产类型信息
// @Produce		json
// @Param			token	header		string					false	"用户凭证"
// @Param			body	body		models.SwustAssetType	false	"待添加资产类型"
// @Success		200		{object}	common.ResponseModel
// @router			/api/assets/type/update [Post]
func (api *typeController) UpdateApi(context *gin.Context) {
	api.Update(context, getService(), newModel())
}

// @title			分页查询资产类型信息
// @version		1.0
// @Tags			AssetType-资产类型管理相关接口
// @description	分页查询资产类型信息
// @accept			json
// @Produce		json
// @Param			token	header		string					false	"用户凭证"
// @Param			req		body		models.TypeQueryRequest	false	"筛选条件"
// @Success		200		{object}	models.SwustAssetType
// @router			/api/assets/type/page [Post]
func (api *typeController) PageApi(context *gin.Context) {
	api.Page(context, getService(), &models.TypeQueryRequest{})
}
