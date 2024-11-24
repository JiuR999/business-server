package common

import (
	"BusinessServer/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	context *gin.Context
}
type PageResponseModel struct {
	Total int `json:"total"` //总记录数
	Page  PageModel
	Data  any `json:"data"`
}
type ResponseModel struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func NewResponse(context *gin.Context) *Response {
	return &Response{
		context: context,
	}
}

func (r *Response) Success() {
	r.context.JSON(http.StatusOK, ResponseModel{
		Code: http.StatusOK,
		Data: common.STATUS_SUCCESS,
	})
}

func (r *Response) SuccessWithMsg(msg string) {
	r.context.JSON(http.StatusOK, ResponseModel{
		Code: http.StatusOK,
		Msg:  msg,
	})
}

func (r *Response) SuccessWithData(data any) {
	r.context.JSON(http.StatusOK, ResponseModel{
		Code: http.StatusOK,
		Msg:  common.STATUS_SUCCESS,
		Data: data,
	})
}

func (r *Response) SuccessWithPageData(model PageResponseModel) {
	responseModel := PageResponseModel{
		Total: model.Total,
		Page: PageModel{
			PageNum:  model.Page.PageNum,
			PageSize: model.Page.PageSize,
		},
		Data: model.Data,
	}
	result := &ResponseModel{
		Code: http.StatusOK,
		Msg:  common.STATUS_SUCCESS,
		Data: responseModel,
	}
	r.context.JSON(http.StatusOK, result)
}

func (r *Response) ErrorWithMsg(msg string) {
	r.context.JSON(common.RETURN_FAILED, ResponseModel{
		Code: common.RETURN_FAILED,
		Data: msg,
	})
}

func (r *Response) ErrorWithCode(code int, msg string) {
	r.context.JSON(common.RETURN_FAILED, ResponseModel{
		Code: code,
		Data: msg,
	})
}
