package abstract

import (
	"BusinessServer/common/abstract/models"
	abstract "BusinessServer/common/abstract/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func (c *Controller) GetById(context *gin.Context, service abstract.Service, model any) {
	response := common.NewResponse(context)
	err := service.GetById(context, model)
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithData(model)
}

func (c *Controller) Add(context *gin.Context, service abstract.Service, model any) {
	response := common.NewResponse(context)
	if err := context.ShouldBindJSON(model); err != nil {
		response.ErrorWithMsg(err.Error())
		return
	}

	res, err := service.Add(context, model)
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithData(res)
}
func (c *Controller) Delete(context *gin.Context, service abstract.Service) {
	response := common.NewResponse(context)
	var ids []string
	if err := context.ShouldBindJSON(&ids); err != nil {
		response.ErrorWithMsg(err.Error())
		return
	}
	affects, err := service.DeleteByIds(context, ids)
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithMsg(fmt.Sprintf("成功删除%d条记录", affects))
}

func (c *Controller) Update(context *gin.Context, service abstract.Service, req any) {
	response := common.NewResponse(context)
	if err := context.ShouldBindJSON(req); err != nil {
		response.ErrorWithMsg(err.Error())
	}
	err := service.Update(context, req)
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	response.SuccessWithMsg("更新成功！")
}

func (c *Controller) Page(context *gin.Context, service abstract.Service, req any) {
	response := common.NewResponse(context)
	res, err := service.Page(context, req)
	if err != nil {
		response.ErrorWithMsg(err.GetMsg())
		return
	}
	fmt.Println("共查询得到", res.Total, "条记录")
	response.SuccessWithPageData(res)
}
