package service

import (
	"BusinessServer/apps/assets/type/dao"
	"BusinessServer/apps/assets/type/models"
	"BusinessServer/apps/system/log/service"
	"BusinessServer/common"
	common2 "BusinessServer/common/abstract/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

type typeService struct {
}

var rService = new(typeService)

func GetTypeService() *typeService {
	return rService
}

func (a *typeService) Add(context *gin.Context, req any) (res any, err common.SwustError) {
	assetsModel := req.(*models.SwustAssetType)
	assetsModel.SetNewId()
	res, err = dao.GetTypeDao().Add(assetsModel)
	service.WriteLog(context, common.LOG_EVENT_ADD, fmt.Sprintf("增加资产类型-%s", assetsModel.Name))
	if err != nil {
		return nil, common.NewServiceError(err.GetMsg())
	}
	return res, nil
}

func (a *typeService) Update(context *gin.Context, req any) (err common.SwustError) {
	err = dao.GetTypeDao().Update(req)
	if err != nil {
		return common.NewServiceError(err.GetMsg())
	}
	r := req.(*models.SwustAssetType)
	service.WriteLog(context, common.LOG_EVENT_ADD, fmt.Sprintf("更新资产类型%s", r.Name))
	return nil
}
func (a *typeService) GetById(context *gin.Context, model any) (err common.SwustError) {
	id := context.Query("id")
	if id != "" {
		swustError := dao.GetTypeDao().GetById(id, model)
		if swustError != nil {
			return common.NewServiceError(swustError.GetMsg())
		}
		return nil
	}
	return common.NewServiceError("查询ID不能为空！")
}

func (a *typeService) DeleteByIds(context *gin.Context, ids []string) (res int64, err common.SwustError) {

	if len(ids) < 0 {
		return 0, common.NewServiceError("请输入待删除资产ID")
	}
	service.WriteLog(context, common.LOG_EVENT_ADD, fmt.Sprintf("删除资产类型%s", ids))
	affects, err := a.Del(context, ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	return affects, nil

}

func (a *typeService) Del(context *gin.Context, ids []string) (int64, common.SwustError) {
	affects, err := dao.GetTypeDao().Delete(ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	return affects, nil

}

func (a *typeService) Page(context *gin.Context, model any) (res common2.PageResponseModel, err common.SwustError) {
	req := model.(*models.TypeQueryRequest)
	req.IfAbsent()
	var record []models.SwustAssetType
	total, err := dao.GetTypeDao().Page(req, &record)
	if err != nil {
		return res, common.NewServiceError(err.GetMsg())
	}
	res = common2.PageResponseModel{
		Total: int(total),
		Page: common2.PageModel{
			PageNum:  req.PageNum,
			PageSize: req.PageSize,
		},
		Data: record,
	}
	return res, nil
}

func (a *typeService) GetTypeList() ([]string, common.SwustError) {
	return dao.GetTypeDao().GetTypeList()
}
