package models

import common "BusinessServer/common/abstract/models"

const TableNameSwustRoom = "swust_room"

// SwustRoomModel mapped from table <swust_room>
type SwustRoomModel struct {
	common.DefaultModel
	Floor    *int64  `gorm:"column:floor;type:int" json:"floor"`               // 楼层
	RoomNum  *string `gorm:"column:number;type:varchar(4)" json:"room_num"`    // 房间号
	Location *string `gorm:"column:location;type:varchar(25)" json:"location"` // 房间位置
	Header   *string `gorm:"column:header;type:bigint" json:"header"`          // 房间负责人Id
}

// TableName SwustRoomModel's table name
func (*SwustRoomModel) TableName() string {
	return TableNameSwustRoom
}
