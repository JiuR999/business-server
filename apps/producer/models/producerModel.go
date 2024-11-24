package models

import common "BusinessServer/common/abstract/models"

const TableNameSwustProducer = "swust_producer"

// SwustProducer mapped from table <swust_producer>
type SwustProducer struct {
	common.DefaultModel
	ContactUser  *string `gorm:"column:contact_user;type:varchar(255)" json:"contact_user"`  // 联络人
	ContactPhone *string `gorm:"column:contact_phone;type:varchar(11)" json:"contact_phone"` // 联系电话号码
	Address      *string `gorm:"column:address;type:varchar(255)" json:"address"`            // 联络地址
}

type ProducerQueryRequest struct {
	common.PageModel
	Name        *string `json:"name"`         //名称
	ContactUser *string `json:"contact_user"` // 联络人
	PhoneNumber *string `json:"phoneNumber"`  // 手机号
	Address     *string `json:"address"`      //厂商位置
}

// TableName SwustProducer's table name
func (*SwustProducer) TableName() string {
	return TableNameSwustProducer
}
