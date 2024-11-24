package controller

import (
	model "BusinessServer/apps/assets/models"
	"BusinessServer/apps/assets/service"
	logWriter "BusinessServer/apps/system/log/service"
	common2 "BusinessServer/common"
	controller "BusinessServer/common/abstract/controller"
	"BusinessServer/common/abstract/models"
	as "BusinessServer/common/abstract/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
	"strings"
	"time"
)

type assetsController struct {
	controller.Controller
}

var ac = new(assetsController)

func GetAssetController() *assetsController {
	return ac
}
func getService() as.Service {
	return service.GetAssetsService()
}

func newModel() *model.AssetsModel {
	return &model.AssetsModel{}
}

// @title			根据资产ID获取资产信息
// @version		1.0
// @Tags			Asset-资产设备管理雄相关接口
// @description	根据资产ID获取资产信息
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Param			id		query		string	false	"资产ID"
// @Success		200		{object}	common.ResponseModel
// @router			/api/assets/getById [Get]
func (api *assetsController) GetApi(context *gin.Context) {
	api.GetById(context, getService(), newModel())
}

// @title			增加资产
// @version		1.0
// @Tags			Asset-资产设备管理雄相关接口
// @description	增加资产
// @Produce		json
// @Param			token	header		string				false	"用户凭证"
// @Param			body	body		models.AssetsModel	false	"待添加资产D"
// @Success		200		{object}	common.ResponseModel
// @router			/api/assets/add [Post]
func (api *assetsController) AddApi(context *gin.Context) {
	api.Add(context, getService(), newModel())
}

// @title			根据资产IDs删除资产信息
// @version		1.0
// @Tags			Asset-资产设备管理雄相关接口
// @description	根据资产IDs删除资产信息
// @accept			json
// @Produce		json
// @Param			token	header		string		false	"用户凭证"
// @Param			ids		body		[]string	false	"待删除资产IDs"
// @Success		200		{object}	common.ResponseModel
// @router			/api/assets/delete [Post]
func (api *assetsController) DeleteApi(context *gin.Context) {
	api.Delete(context, getService())
}

// @title			更新资产信息
// @version		1.0
// @Tags			Asset-资产设备管理雄相关接口
// @description	更新资产信息
// @Produce		json
// @Param			token	header		string				false	"用户凭证"
// @Param			body	body		models.AssetsModel	false	"待添加资产"
// @Success		200		{object}	common.ResponseModel
// @router			/api/assets/update [Post]
func (api *assetsController) UpdateApi(context *gin.Context) {
	api.Update(context, getService(), newModel())
}

// @title			分页查询资产信息
// @version		1.0
// @Tags			Asset-资产设备管理雄相关接口
// @description	分页查询资产信息
// @accept			json
// @Produce		json
// @Param			token	header		string						false	"用户凭证"
// @Param			req		body		models.AssetsQueryRequest	false	"待删除资产IDs"
// @Success		200		{object}	models.AssetsVO
// @router			/api/assets/page [Post]
func (api *assetsController) PageApi(context *gin.Context) {
	logWriter.WriteLog(context, common2.LOG_EVENT_READ, "查询设备列表")
	api.Page(context, getService(), &model.AssetsQueryRequest{})
}

// @title			获取导入资产信息模板
// @version		1.0
// @Tags			Asset-资产设备管理雄相关接口
// @description	获取导入资产信息模板
// @accept			json
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Success		200		{object}	string
// @router			/api/assets/template [Get]
func (api *assetsController) Template(context *gin.Context) {
	logWriter.WriteLog(context, common2.LOG_EVENT_READ, "获取导入资产设备模板")
	context.Header("Content-Type", "application/octet-stream")
	disposition := fmt.Sprintf("attachment; filename*=utf-8''%s-%s.xlsx", url.QueryEscape("资产设备导入模板"), time.Now().Format(time.DateTime))
	context.Header("Content-Disposition", disposition)
	file, err := service.GetAssetsService().Template()
	file.SaveAs("asset.xlsx")
	if err != nil {
		common.NewResponse(context).ErrorWithMsg(err.GetMsg())
		return
	}
	context.File("asset.xlsx")
}

// @title			导入资产信息
// @version		1.0
// @Tags			Asset-资产设备管理雄相关接口
// @description	导入资产信息
// @accept			json
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Param			file	formData	file	true	"文件上传"
// @Success		200		{object}	string
// @router			/api/assets/import [Post]
func (api *assetsController) Import(context *gin.Context) {
	logWriter.WriteLog(context, common2.LOG_EVENT_ADD, "导入资产设备")
	response := common.NewResponse(context)
	file, err := context.FormFile("file")
	if err != nil {
		response.ErrorWithMsg(err.Error())
		return
	}
	err2 := service.GetAssetsService().Import(file)
	if err2 != nil {
		response.ErrorWithMsg(err2.GetMsg())
		return
	}
	common.NewResponse(context).SuccessWithMsg("文件上传成功，等待数据导入！")
}

// @title			导出资产信息
// @version		1.0
// @Tags			Asset-资产设备管理雄相关接口
// @description	导出资产信息
// @accept			json
// @Produce		json
// @Param			token	header		string		false	"用户凭证"
// @Param			ids		body		[]string	true	"需要导出的ids"
// @Success		200		{object}	string
// @router			/api/assets/export [Post]
func (api *assetsController) Export(context *gin.Context) {
	logWriter.WriteLog(context, common2.LOG_EVENT_READ, "导出资产设备")
	var (
		ids []string
	)
	response := common.NewResponse(context)
	if err := context.ShouldBindJSON(&ids); err != nil {
		response.ErrorWithMsg(err.Error())
		return
	}

	file, msgs := service.GetAssetsService().Export(ids)
	if len(msgs) > 0 {
		common.NewResponse(context).ErrorWithMsg(strings.Join(msgs, "||"))
	}
	file.SaveAs("asset.xlsx")
	context.Header("Content-Type", "application/octet-stream")
	disposition := fmt.Sprintf("attachment; filename*=utf-8''%s-%s.xlsx", url.QueryEscape("资产设备列表"), time.Now().Format(time.DateTime))
	context.Header("Content-Disposition", disposition)
	context.File("asset.xlsx")
	common.NewResponse(context).Success()
}
