package models

import (
	"BusinessServer/common/Time"
	common "BusinessServer/common/abstract/models"
	"time"
)

const TableNameSwustAssets = "swust_asset"

// AssetsModel mapped from table <swust_assets>
type AssetsModel struct {
	common.DefaultModel
	Code           *string        `gorm:"column:code;type:varchar(50)" json:"code"`                              // 资产编码
	UserID         *string        `gorm:"column:user_id;type:varchar(19)" json:"userID"`                         // 责任人
	ProductionTime *Time.LocalDay `gorm:"column:production_time;type:date" json:"productionTime"`                // 生产日期
	ServiceLength  *int64         `gorm:"column:service_length;type:int" json:"serviceLength"`                   // 服役年限
	Price          *float32       `gorm:"column:price;type:float" json:"price"`                                  // 价格
	Status         *int64         `gorm:"column:status;type:int" json:"status"`                                  // 0 在用 1 故障 2 维修 3 报废
	TypeID         *string        `gorm:"column:type_id;type:varchar(19)" json:"typeID"`                         // 所在房间
	ProducerID     *string        `gorm:"column:producer_id;type:varchar(19)" json:"producerID"`                 // 厂商ID
	OrderID        *string        `gorm:"column:order_id;type:varchar(19)" json:"orderID"`                       // 采购编号
	Comment        *string        `gorm:"column:comment;type:varchar(255)" json:"comment"`                       // 备注
	UpdateTime     *time.Time     `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"updateTime"` // 更新时间
	CreateTime     *time.Time     `gorm:"column:create_time;type:int unsigned;autoCreateTime" json:"createTime"` // 创建时间
}

// TableName AssetsModel's table name
func (*AssetsModel) TableName() string {
	return TableNameSwustAssets
}

type AssetsVO struct {
	common.DefaultModel
	Code           *string       `gorm:"column:code;type:varchar(50)" json:"code"`                              // 资产编码
	UserID         *string       `gorm:"column:user_id;type:varchar(19)" json:"userID"`                         // 责任人
	ProductionTime Time.LocalDay `gorm:"column:production_time;type:date" json:"productionTime"`                // 生产日期
	ServiceLength  *int64        `gorm:"column:service_length;type:int" json:"serviceLength"`                   // 服役年限
	Price          *float32      `gorm:"column:price;type:float" json:"price"`                                  // 价格
	Status         *int64        `gorm:"column:status;type:int" json:"status"`                                  // 0 在用 1 故障 2 维修 3 报废
	Comment        *string       `gorm:"column:comment;type:varchar(255)" json:"comment"`                       // 备注
	UpdateTime     *time.Time    `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"updateTime"` // 更新时间
	Header         string        `gorm:"column:header" json:"header"`                                           //负责人
	//RoomName       *string       `gorm:"column:room_name;type:varchar(25)" json:"roomName"`                     // 房间名称
	//Floor          *int64        `gorm:"column:floor;type:int" json:"floor"`                                    // 楼层
	//RoomNum        *string       `gorm:"column:number;type:varchar(4)" json:"room_num"`                         // 房间号
	//Address       *string       `gorm:"column:location;type:varchar(25)" json:"location"`                      // 房间位置
	TypeName       *string    `gorm:"column:type_name" json:"typeName"`                //类型名称
	RetireMentTime *time.Time `gorm:"column:retirement_time" json:"retireMentTime"`    //报废退役年限
	ProducerName   *string    `gorm:"column:producer_name" json:"producerName"`        //供应商名字
	Quantity       int        `gorm:"quantity" json:"quantity,string"`                 //数量
	OrderID        *string    `gorm:"column:order_id;type:varchar(19)" json:"orderID"` // 采购编号
	//Unit           string        `gorm:"column:unit" json:"unit"`                                               //单位
}
