package models

import (
	common "BusinessServer/common/abstract/models"
	"time"
)

const TableNameSwustOrder = "swust_order"

// SwustOrder mapped from table <swust_order>
type SwustOrder struct {
	common.DefaultModel
	UserID      *string    `gorm:"column:user_id;type:varchar(19)" json:"user_id"`                      // 采购申请人
	Status      int64      `gorm:"column:status;type:int" json:"status"`                                // 进度：0 待审批 1 采购中 2 采购完成 -1拒绝
	ApplyTime   *time.Time `gorm:"column:apply_time;type:int unsigned;autoCreateTime" json:"applyTime"` // 申请时间
	AllowUserID *string    `gorm:"column:allow_user_id;type:varchar(19)" json:"allow_user_id"`          // 审批人
}

type SwustOrderVO struct {
	SwustOrder
	ApplyUser     string `gorm:"column:apply_user" json:"applyUser"`
	AllowUserName string `gorm:"column:allow_user" json:"allowUserName"`
}

// TableName SwustOrder's table name
func (*SwustOrder) TableName() string {
	return TableNameSwustOrder
}

type OrderRequest struct {
	UserID    *string             `gorm:"column:user_id;type:varchar(19)" json:"user_id"` // 采购申请人
	OrderName string              `gorm:"column:name" json:"orderName"`                   //采购事由
	Assets    []OrderAssetRequest `json:"assets"`                                         //采购明细信息
}

type OrderAssetRequest struct {
	Name    *string `gorm:"column:name" json:"name"`
	OrderId string  `gorm:"column:order_id" json:"orderId"` //采购订单id
	//UserID     *string    `gorm:"column:user_id;type:varchar(19)" json:"userID"`                       // 责任人
	ApplyTime  *time.Time `gorm:"column:apply_time;type:int unsigned;autoCreateTime" json:"applyTime"` // 申请时间
	TypeId     *string    `gorm:"column:type_id" json:"typeId"`                                        //类型名称
	OrderTime  *time.Time `gorm:"column:order_time;type:int unsigned;autoUpdateTime" json:"orderTime"` // 采购时间
	ProducerId *string    `gorm:"column:producer_id" json:"producerId"`                                //供应商名字
	Quantity   int        `gorm:"quantity" json:"quantity,string"`                                     //数量
	Comment    *string    `gorm:"column:comment" json:"comment"`                                       // 备注
}

type OrderQueryRequest struct {
	common.PageModel
	StartTime *time.Time `json:"startTime"` // 申请起始时间
	EndTime   *time.Time `json:"endTime"`   // 申请结束时间
	Status    *int       `json:"status"`    //采购单状态
	ApplyUser string     `json:"applyUser"` //申请人名称
}
