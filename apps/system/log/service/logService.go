package service

import (
	"BusinessServer/apps/system/log/dao"
	"BusinessServer/apps/system/log/models"
	dao2 "BusinessServer/apps/system/user/dao"
	models2 "BusinessServer/apps/system/user/models"
	"BusinessServer/common"
	common2 "BusinessServer/common/abstract/models"
	publisher "BusinessServer/common/services"
	"BusinessServer/common/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type logService struct {
}

var ls = new(logService)

func GetLogService() *logService {
	return ls
}

func StartLogService() {
	log.Print("启用日志服务")
	eventChan := publisher.EB.Subscribe(common.EVENT_LOG)
	for {
		select {
		case event := <-eventChan:
			fmt.Println("收到日志写入事件")
			data := event.Data.(*models.SwustSystemLog)
			//fmt.Println(data.Content)
			dao.GetLogDao().Add(data)
		}
	}
}

func (a *logService) Add(context *gin.Context, req any) (res any, err common.SwustError) {
	logModel := req.(*models.SwustSystemLog)
	logModel.SetNewId()

	return res, nil
}

func (a *logService) Update(context *gin.Context, req any) (err common.SwustError) {
	err = dao.GetLogDao().Update(req)
	if err != nil {
		return common.NewServiceError(err.GetMsg())
	}
	return nil
}
func (a *logService) GetById(context *gin.Context, model any) (err common.SwustError) {
	id := context.Query("id")
	if id != "" {
		swustError := dao.GetLogDao().GetById(id, model)
		if swustError != nil {
			return common.NewServiceError(swustError.GetMsg())
		}
		return nil
	}
	return common.NewServiceError("查询ID不能为空！")
}

func (a *logService) DeleteByIds(context *gin.Context, ids []string) (res int64, err common.SwustError) {

	if len(ids) < 0 {
		return 0, common.NewServiceError("请输入待删除资产ID")
	}
	affects, err := a.Del(context, ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	return affects, nil

}

func (a *logService) Del(context *gin.Context, ids []string) (int64, common.SwustError) {
	affects, err := dao.GetLogDao().Delete(ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	return affects, nil

}

func (a *logService) Page(context *gin.Context, req any) (res common2.PageResponseModel, err common.SwustError) {
	if err := context.ShouldBindJSON(req); err != nil {
		common2.NewResponse(context).ErrorWithMsg(err.Error())
	}
	request := req.(*models.SystemLogQueryRequest)
	request.IfAbsent()

	id, swustError := dao2.GetUserDao().GetIdByNameAndPhone(models2.SystemUserQueryRequest{Name: request.OperateUser})
	if swustError != nil {
		return res, common.NewServiceError("该用户不存在")
	}
	request.OperateUser = id
	var record []models.SwustSystemLogVO
	total, err := dao.GetLogDao().Page(request, &record)
	if err != nil {
		return res, common.NewServiceError(err.GetMsg())
	}
	res = common2.PageResponseModel{
		Total: int(total),
		Page: common2.PageModel{
			PageSize: request.PageSize,
			PageNum:  request.PageNum,
		},
		Data: record,
	}
	return res, nil
}

// WriteLog 写入日志
func WriteLog(ctx *gin.Context, event int64, content string) {
	claims, exists := ctx.Get(common.COMMON_AUTH_CURRENT)
	user := claims.(*utils.MyClaims)
	if exists {
		log := &models.SwustSystemLog{
			Content: content,
			Event:   event,
			UserID:  user.UserId,
		}
		log.SetNewId()
		publisher.EB.Publish(common.EVENT_LOG, publisher.EventModel{
			Event: "写入日志",
			Data:  log,
		})
		//dao.GetLogDao().Add(&log)
	}
}
