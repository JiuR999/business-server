package service

import (
	"BusinessServer/apps/producer/dao"
	"BusinessServer/apps/producer/models"
	"BusinessServer/common"
	common2 "BusinessServer/common/abstract/models"
	"github.com/gin-gonic/gin"
)

type producerService struct {
}

var rService = new(producerService)

func GetProducerService() *producerService {
	return rService
}

func (a *producerService) Add(context *gin.Context, req any) (res any, err common.SwustError) {
	assetsModel := req.(*models.SwustProducer)
	assetsModel.SetNewId()
	res, err = dao.GetProducerDao().Add(assetsModel)
	if err != nil {
		return nil, common.NewServiceError(err.GetMsg())
	}
	return res, nil
}

func (a *producerService) Update(context *gin.Context, req any) (err common.SwustError) {
	err = dao.GetProducerDao().Update(req)
	if err != nil {
		return common.NewServiceError(err.GetMsg())
	}
	return nil
}
func (a *producerService) GetById(context *gin.Context, model any) (err common.SwustError) {
	id := context.Query("id")
	if id != "" {
		swustError := dao.GetProducerDao().GetById(id, model)
		if swustError != nil {
			return common.NewServiceError(swustError.GetMsg())
		}
		return nil
	}
	return common.NewServiceError("查询ID不能为空！")
}

func (a *producerService) DeleteByIds(context *gin.Context, ids []string) (res int64, err common.SwustError) {

	if len(ids) < 0 {
		return 0, common.NewServiceError("请输入待删除资产ID")
	}
	affects, err := a.Del(context, ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	return affects, nil

}

func (a *producerService) Del(context *gin.Context, ids []string) (int64, common.SwustError) {
	affects, err := dao.GetProducerDao().Delete(ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	return affects, nil

}

func (a *producerService) Page(context *gin.Context, req any) (res common2.PageResponseModel, err common.SwustError) {
	request := req.(*models.ProducerQueryRequest)
	request.IfAbsent()
	var record []models.SwustProducer
	total, err := dao.GetProducerDao().Page(request, &record)
	if err != nil {
		return res, common.NewServiceError(err.GetMsg())
	}
	res = common2.PageResponseModel{
		Total: int(total),
		Page: common2.PageModel{
			PageNum:  request.PageNum,
			PageSize: request.PageSize,
		},
		Data: record,
	}
	return res, nil
}

func (a *producerService) GetProducerList() ([]string, common.SwustError) {
	return dao.GetProducerDao().GetProducerList()
}

func (a *producerService) GetProducerIds() ([]string, common.SwustError) {
	return dao.GetProducerDao().GetProducerIds()
}
