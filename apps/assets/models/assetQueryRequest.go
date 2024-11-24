package models

import (
	"BusinessServer/common/abstract/models"
	"time"
)

type AssetsQueryRequest struct {
	common.PageModel
	AssetsModel
	EndTime *time.Time `json:"endTime"` //结束时间
	Ids     []string   `json:"ids"`     //根据传入ids查询
}
