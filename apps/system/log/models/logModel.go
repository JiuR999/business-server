package models

import (
	common "BusinessServer/common/abstract/models"
	"time"
)

const TableNameSwustSystemLog = "swust_system_log"

// SwustSystemLog mapped from table <swust_system_log>
type SwustSystemLog struct {
	common.IdModel
	Content     string    `gorm:"column:content;type:varchar(255)" json:"content"`                         // 日志内容
	OperateTime time.Time `gorm:"column:operate_time;type:int unsigned;autoCreateTime" json:"operateTime"` // 操作时间
	UserID      string    `gorm:"column:user_id" json:"user_id"`                                           // 操作人
	Event       int64     `gorm:"column:event" json:"event"`                                               //日志事件 1 增加 2 删除 3 更新 4 查询 5 审批 6 登录
}

type SwustSystemLogVO struct {
	common.IdModel
	Content      string    `gorm:"column:content;type:varchar(255)" json:"content"`                         // 日志内容
	OperateTime  time.Time `gorm:"column:operate_time;type:int unsigned;autoCreateTime" json:"operateTime"` // 操作时间
	OperatorName string    `gorm:"column:operator_name" json:"operateName"`                                 // 操作人
	Event        int64     `gorm:"column:event" json:"event"`                                               //日志事件 1 增加 2 删除 3 更新 4 查询 5 审批 6 登录
}

func NewSystemLog(event int64, userId, content string) *SwustSystemLog {
	log := &SwustSystemLog{
		Event:   event,
		UserID:  userId,
		Content: content,
	}
	log.SetNewId()
	return log
}

// TableName SwustSystemLog's table name
func (*SwustSystemLog) TableName() string {
	return TableNameSwustSystemLog
}

type SystemLogQueryRequest struct {
	common.PageModel
	Event       int64      `json:"event"`       //事件类型
	StartTime   *time.Time `json:"startTime"`   //操作时间
	EndTime     *time.Time `json:"endTime"`     //结束时间
	OperateUser string     `json:"operateUser"` //操作人
}
