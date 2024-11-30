package models

import (
	common "BusinessServer/common/abstract/models"
	"time"
)

const TableNameSwustOrder = "swust_order"

// SwustOrder mapped from table <swust_order>
type SwustOrder struct {
	common.DefaultModel
	UserID      *string    `gorm:"column:user_id;type:varchar(19)" json:"userID"`                       // 采购申请人
	Status      int64      `gorm:"column:status;type:int" json:"status"`                                // 进度：0 待审批 1 采购中 2 采购完成 -1拒绝
	ApplyTime   *time.Time `gorm:"column:apply_time;type:int unsigned;autoCreateTime" json:"applyTime"` // 申请时间
	AllowUserID *string    `gorm:"column:approve_id;type:varchar(19)" json:"allowUserID"`               // 审批人
	Reason      string     `gorm:"column:reason" json:"reason"`                                         //拒绝理由
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

type OrderVO struct {
	UserID      *string        `gorm:"column:user_id;type:varchar(19)" json:"userID"`       // 采购申请人ID
	ApplyUser   *string        `gorm:"column:apply_user;type:varchar(19)" json:"applyUser"` // 采购申请人
	ApproveUser *string        `gorm:"column:approve_user" json:"approveUser"`              // 审批人
	ApplyTime   *time.Time     `gorm:"column:apply_time;type:varchar(19)" json:"applyTime"` // 采购申请人
	ApproveTime *time.Time     `gorm:"column:approve_time" json:"approveTime"`              // 审批人
	OrderName   string         `gorm:"column:name" json:"orderName"`                        //采购事由
	Status      int            `gorm:"column:status" json:"status"`                         //当前进度
	Reason      *string        `gorm:"column:reason" json:"reason"`                         //驳回理由
	Assets      []OrderAssetVO `gorm:"-" json:"assets"`                                     //采购明细信息
}

type OrderAssetVO struct {
	Name    *string `gorm:"column:name" json:"name"`
	OrderId string  `gorm:"column:order_id" json:"orderId"` //采购订单id
	//UserID     *string    `gorm:"column:user_id;type:varchar(19)" json:"userID"`                       // 责任人
	ApplyTime    *time.Time `gorm:"column:apply_time;type:int unsigned;autoCreateTime" json:"applyTime"` // 申请时间
	TypeName     *string    `gorm:"column:type_name" json:"typeName"`                                    //类型名称
	OrderTime    *time.Time `gorm:"column:order_time;type:int unsigned;autoUpdateTime" json:"orderTime"` // 采购时间
	OrderUser    *string    `gorm:"column:order_user" json:"orderUser"`                                  //采购人
	ProducerName *string    `gorm:"column:producer_name" json:"producerName"`                            //供应商名字
	Quantity     int        `gorm:"quantity" json:"quantity,string"`                                     //数量
	Comment      *string    `gorm:"column:comment" json:"comment"`                                       // 备注
	TypeId       *string    `gorm:"column:type_id" json:"typeId"`                                        //类型名称
	ProducerId   *string    `gorm:"column:producer_id" json:"producerId"`                                //供应商Id
}

type ApproveReq struct {
	UserId      string    `gorm:"column:approve_id" json:"userId"`        //审批人ID
	OrderId     string    `gorm:"column:id" json:"orderId"`               //采购单ID
	Status      int       `gorm:"column:status" json:"status"`            //审核意见 -1 拒绝 1 同意
	Reason      *string   `gorm:"column:reason" json:"reason"`            //驳回理由
	ApproveTime time.Time `gorm:"column:approve_time" json:"approveTime"` //审批时间
}
