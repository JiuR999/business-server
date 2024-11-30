package service

import (
	dao2 "BusinessServer/apps/assets/dao"
	models2 "BusinessServer/apps/assets/models"
	"BusinessServer/apps/order/dao"
	"BusinessServer/apps/order/models"
	"BusinessServer/apps/producer/service"
	"BusinessServer/common"
	"BusinessServer/common/Time"
	common2 "BusinessServer/common/abstract/models"
	publisher "BusinessServer/common/services"
	"BusinessServer/common/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"time"
)

type orderService struct {
}

var oService = new(orderService)

func GetOrderService() *orderService {
	return oService
}

func StartAsyncService() {
	log.Print("启用异步事件服务")
	eventChan := publisher.EB.Subscribe(common.EVENT_ASYNC)
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
			publisher.EB.UnSubscribe(common.EVENT_ASYNC, eventChan)
			StartAsyncService()
		}
	}()
	for {
		select {
		case event := <-eventChan:
			fmt.Println("收到异步事件")
			id := event.Data.(string)
			order := &models.OrderVO{}
			dao.GetOrderDao().GetById(id, order)
			//更新order状态
			dao.GetOrderDao().UpdateStatus(2, id)
			//更新采购时间
			result, err := dao.GetOrderDao().GetOrderDetailById(id)
			if err != nil {
				log.Print(err.GetMsg())
				break
			}
			fmt.Println(result)
			assets := make([]models2.AssetsModel, 0)
			for i, vo := range result {
				for j := 0; j < vo.Quantity; j++ {
					producerList, _ := service.GetProducerService().GetProducerIds()
					pIndex := rand.Intn(len(producerList))
					now := time.Now()
					pt := Time.LocalDay(now)
					sl := rand.Int63n(10)
					price := float32(rand.Int63n(1))
					status := int64(4)
					ou := "采购员甲"
					result[i].OrderTime = &(now)
					result[i].OrderUser = &ou
					result[i].ProducerId = &producerList[pIndex]
					fmt.Println("生产厂商：", producerList[pIndex])
					//TODO 开启事务
					e := dao.GetOrderDao().UpdateOrderDetail(result[i])
					if e != nil {
						fmt.Println(e.GetMsg())
					}
					code := utils.GenerateCode()
					asset := models2.AssetsModel{
						Code:           &code,
						UserID:         order.UserID,
						ProductionTime: &pt,
						ServiceLength:  &sl,
						Price:          &price,
						Status:         &status,
						TypeID:         vo.TypeId,
						ProducerID:     vo.ProducerId,
						OrderID:        &vo.OrderId,
						Comment:        vo.Comment,
					}
					asset.Id = utils.GenerateId()
					asset.Code = &(code)
					asset.Name = *vo.Name + utils.GenerateId()[15:]
					assets = append(assets, asset)
				}
			}

			//插入设备资产表
			_, swustError := dao2.GetAssetDao().Add(&assets)
			if swustError != nil {
				fmt.Println(swustError.GetMsg())
			} else {
				fmt.Println("入库成功!")
				publisher.EB.Publish(common.EVENT_NOTICE, publisher.EventModel{
					Event: common.EVENT_NOTICE,
					Data:  "订单号" + id + "完成采购流程",
				})
			}

		}
	}
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
		order := models.OrderVO{}
		swustError := dao.GetOrderDao().GetById(id, &order)
		if swustError != nil {
			return common.NewServiceError(swustError.GetMsg())
		}
		//获取采购详情表
		assets, e := dao.GetOrderDao().GetOrderDetailById(id)
		if e != nil {
			return e
		}
		order.Assets = assets
		res := model.(*models.OrderVO)
		*res = order
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

func (a *orderService) Approve(req models.ApproveReq) common.SwustError {
	req.ApproveTime = time.Now()

	if err := dao.GetOrderDao().Approve(req); err != nil {
		return common.NewServiceError(err.GetMsg())
	}

	if req.Status == common.APPROVE_AGREE {
		go func() {
			defer func() {
				if e := recover(); e != nil {
					fmt.Println("订单入库失败 请重试!")

				}
			}()
			time.AfterFunc(5*time.Second, func() {
				//推送消息
				publisher.EB.Publish(common.EVENT_NOTICE, publisher.EventModel{
					Event: common.EVENT_ORDER,
					Data:  "订单:" + req.OrderId + "已完成采购！",
				})
				//发送消息入库
				publisher.EB.Publish(common.EVENT_ASYNC, publisher.EventModel{
					Event: common.EVENT_ORDER,
					Data:  req.OrderId,
				})
			})
		}()
	}
	return nil
}
