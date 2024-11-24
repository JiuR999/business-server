package models

import common "BusinessServer/common/abstract/models"

const TableNameSwustSystemUser = "swust_system_user"

// SystemUserModel mapped from table <swust_system_user>
type SystemUserModel struct {
	common.DefaultModel
	Account  *string `gorm:"column:account;type:varchar(255)" json:"account"`   // 用户名
	Password *string `gorm:"column:password;type:varchar(255)" json:"password"` // 密码
	//Role        *string `gorm:"column:role;type:text" json:"role"`                       // 权限
	Avatar      *string `gorm:"column:avatar;type:varchar(255)" json:"avatar"`           // 头像
	Salt        *string `gorm:"column:salt;type:varchar(255)" json:"salt"`               // 盐
	PhoneNumber *string `gorm:"column:phone_number;type:varchar(12)" json:"phoneNumber"` // 电话号码
}

// TableName SystemUserModel's table name
func (*SystemUserModel) TableName() string {
	return TableNameSwustSystemUser
}
