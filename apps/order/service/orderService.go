package service

import (
	"BusinessServer/apps/order/dao"
	"BusinessServer/apps/order/models"
	"BusinessServer/common"
	common2 "BusinessServer/common/abstract/models"
	"github.com/gin-gonic/gin"
)

type orderService struct {
}

var oService = new(orderService)

func GetOrderService() *orderService {
	return oService
}

func (a *orderService) Add(context *gin.Context, req any) (res any, err common.SwustError) {
	orderModel := &models.SwustOrder{}
	orderReq := req.(*models.OrderRequest)

	orderModel.SetNewId()
	orderModel.UserID = orderReq.UserID
	orderModel.Status = int64(common.STATUS_PREPARE)
	if orderReq.OrderName != "" {
		orderModel.Name = orderReq.OrderName
	}
	for i := range orderReq.Assets {
		orderReq.Assets[i].OrderId = orderModel.Id
	}
	err = dao.GetOrderDao().AddOrderDetailBatch(orderReq.Assets)
	if err != nil {
		return nil, common.NewServiceError(err.GetMsg())
	}
	res, err = dao.GetOrderDao().Add(orderModel)
	if err != nil {
		return nil, common.NewServiceError(err.GetMsg())
	}
	return res, nil
}

func (a *orderService) Update(context *gin.Context, req any) (err common.SwustError) {
	err = dao.GetOrderDao().Update(req)
	if err != nil {
		return common.NewServiceError(err.GetMsg())
	}
	return nil
}
func (a *orderService) GetById(context *gin.Context, model any) (err common.SwustError) {
	id := context.Query("id")
	if id != "" {
		swustError := dao.GetOrderDao().GetById(id, model)
		if swustError != nil {
			return common.NewServiceError(swustError.GetMsg())
		}
		return nil
	}
	return common.NewServiceError("查询ID不能为空！")
}

func (a *orderService) DeleteByIds(context *gin.Context, ids []string) (res int64, err common.SwustError) {

	if len(ids) < 0 {
		return 0, common.NewServiceError("请输入待删除资产ID")
	}
	affects, err := a.Del(context, ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	return affects, nil

}

func (a *orderService) Del(context *gin.Context, ids []string) (int64, common.SwustError) {
	affects, err := dao.GetOrderDao().Delete(ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	return affects, nil

}

func (a *orderService) Page(context *gin.Context, req any) (res common2.PageResponseModel, err common.SwustError) {
	request := req.(*models.OrderQueryRequest)
	request.IfAbsent()
	var record []models.SwustOrderVO
	total, err := dao.GetOrderDao().Page(request, &record)
	if err != nil {
		return res, common.NewServiceError(err.GetMsg())
	}
	res = common2.PageResponseModel{
		Total: int(total),
		Page:  request.PageModel,
		Data:  record,
	}
	return res, nil
}
