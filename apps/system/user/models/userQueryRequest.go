package models

import (
	"BusinessServer/common/abstract/models"
)

type SystemUserQueryRequest struct {
	common.PageModel
	Name        string  `gorm:"column:name;type:varchar(255)" json:"name"`               //名称
	PhoneNumber *string `gorm:"column:phone_number;type:varchar(12)" json:"phoneNumber"` // 电话号码
}

type LoginRequest struct {
	Account  string `json:"account"`  //用户名
	Password string `json:"password"` //密码
}
