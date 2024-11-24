package models

import (
	common "BusinessServer/common/abstract/models"
	"time"
)

const TableNameSwustMenu = "swust_system_menu"

// SwustMenuModel mapped from table <swust_menu>
type SwustMenuModel struct {
	common.DefaultModel
	Router     *string    `gorm:"column:router;type:varchar(255)" json:"router"`                          // 路由
	Level      *string    `gorm:"column:level;type:varchar(255)" json:"level"`                            // 等级
	ParentID   *string    `gorm:"column:parent_id;type:int" json:"parent_id,string"`                      // 父节点id
	Role       *string    `gorm:"column:role;type:varchar(255)" json:"role"`                              // 权限名称
	CreateTime *time.Time `gorm:"column:create_time;type:int unsigned;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime *time.Time `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"update_time"` // 更新时间
}

// TableName SwustMenuModel's table name
func (*SwustMenuModel) TableName() string {
	return TableNameSwustMenu
}
