package controller

import (
	"BusinessServer/apps/statistic/service"
	controller "BusinessServer/common/abstract/controller"
	"BusinessServer/common/abstract/models"
	"github.com/gin-gonic/gin"
)

type statisticController struct {
	controller.Controller
}

var rc = new(statisticController)

func GetStatisticController() *statisticController {
	return rc
}

// @title			根据资产类型统计
// @version		1.0
// @Tags			Statistic-统计相关接口
// @description	根据资产类型统计
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Success		200		{object}	models.StatisticModel
// @router			/api/statistic/countAssetsByType [Get]
func (api *statisticController) CountAssetsByType(context *gin.Context) {
	response := common.NewResponse(context)
	res, err := service.GetStatisticService().CountAssetsByType()
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithData(res)
}

// @title			根据资产状态统计
// @version		1.0
// @Tags			Statistic-统计相关接口
// @description	根据资产状态统计
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Success		200		{object}	models.StatisticModel
// @router			/api/statistic/countAssetsByStatus [Get]
func (api *statisticController) CountAssetsByStatus(context *gin.Context) {
	response := common.NewResponse(context)
	res, err := service.GetStatisticService().CountAssetsByStatus()
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithData(res)
}

// @title			获取资产申请趋势统计
// @version		1.0
// @Tags			Statistic-统计相关接口
// @description	获取资产申请趋势统计
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Success		200		{object}	models.StatisticModel
// @router			/api/statistic/countAssetsApplyTrend [Get]
func (api *statisticController) CountAssetsApplyTrend(context *gin.Context) {
	response := common.NewResponse(context)
	res, err := service.GetStatisticService().CountAssetsApplyTrend()
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithData(res)
}

// @title			获取资产报废趋势统计
// @version		1.0
// @Tags			Statistic-统计相关接口
// @description	获取资产报废趋势统计
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Success		200		{object}	models.StatisticModel
// @router			/api/statistic/countAssetsDepTrend [Get]
func (api *statisticController) CountAssetsDepTrend(context *gin.Context) {
	response := common.NewResponse(context)
	res, err := service.GetStatisticService().CountAssetsDepTrend()
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithData(res)
}

// @title			获取资产报废趋势统计
// @version		1.0
// @Tags			Statistic-统计相关接口
// @description	获取资产报废趋势统计
// @Produce		json
// @Param			token	header		string	false	"用户凭证"
// @Success		200		{object}	models.StatisticModel
// @router			/api/statistic/countOrderDetail [Get]
func (api *statisticController) CountOrderDetail(context *gin.Context) {
	response := common.NewResponse(context)
	res, err := service.GetStatisticService().CountOrderDeTail()
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithData(res)
}
