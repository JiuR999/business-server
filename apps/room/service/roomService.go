package service

import (
	"BusinessServer/apps/room/dao"
	"BusinessServer/apps/room/models"
	"BusinessServer/common"
	common2 "BusinessServer/common/abstract/models"
	"github.com/gin-gonic/gin"
)

type roomService struct {
}

var rService = new(roomService)

func GetRoomService() *roomService {
	return rService
}

func (a *roomService) Add(context *gin.Context, req any) (res any, err common.SwustError) {
	roomModel := req.(*models.SwustRoomModel)
	roomModel.SetNewId()
	res, err = dao.GetRoomDao().Add(roomModel)
	if err != nil {
		return nil, common.NewServiceError(err.GetMsg())
	}
	return res, nil
}

func (a *roomService) Update(context *gin.Context, req any) (err common.SwustError) {
	err = dao.GetRoomDao().Update(req)
	if err != nil {
		return common.NewServiceError(err.GetMsg())
	}
	return nil
}
func (a *roomService) GetById(context *gin.Context, model any) (err common.SwustError) {
	id := context.Query("id")
	if id != "" {
		swustError := dao.GetRoomDao().GetById(id, model)
		if swustError != nil {
			return common.NewServiceError(swustError.GetMsg())
		}
		return nil
	}
	return common.NewServiceError("查询ID不能为空！")
}

func (a *roomService) DeleteByIds(context *gin.Context, ids []string) (res int64, err common.SwustError) {

	if len(ids) < 0 {
		return 0, common.NewServiceError("请输入待删除资产ID")
	}
	affects, err := a.Del(context, ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	return affects, nil

}

func (a *roomService) Del(context *gin.Context, ids []string) (int64, common.SwustError) {
	affects, err := dao.GetRoomDao().Delete(ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	return affects, nil

}

func (a *roomService) Page(context *gin.Context, req any) (res common2.PageResponseModel, err common.SwustError) {
	request := req.(*common2.PageModel)
	request.IfAbsent()
	var record []models.SwustRoomModel
	total, err := dao.GetRoomDao().Page(request, &record)
	if err != nil {
		return res, common.NewServiceError(err.GetMsg())
	}
	res = common2.PageResponseModel{
		Total: int(total),
		Page:  *request,
		Data:  record,
	}
	return res, nil
}

func (a *roomService) ListLocation() (res []string, err common.SwustError) {
	swustError := dao.GetRoomDao().ListLocation(&res)
	if swustError != nil {
		return res, common.NewServiceError(swustError.GetMsg())
	}
	return res, nil
}
