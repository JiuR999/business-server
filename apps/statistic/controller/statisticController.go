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
// @Success		200		{object}	models.AssetTypeModel
// @router			/api/statistic/countassetsbytype [Get]
func (api *statisticController) CountAssetsByType(context *gin.Context) {
	response := common.NewResponse(context)
	res, err := service.GetStatisticService().CountAssetsByType()
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithData(res)
}
