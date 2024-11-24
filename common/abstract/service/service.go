package abstract

import (
	"BusinessServer/common"
	common2 "BusinessServer/common/abstract/models"
	"github.com/gin-gonic/gin"
)

type Service interface {
	Add(context *gin.Context, req any) (res any, err common.SwustError)
	DeleteByIds(context *gin.Context, ids []string) (affects int64, err common.SwustError)
	Update(context *gin.Context, req any) (err common.SwustError)
	GetById(context *gin.Context, model any) (err common.SwustError)
	Page(context *gin.Context, req any) (res common2.PageResponseModel, err common.SwustError)
}
