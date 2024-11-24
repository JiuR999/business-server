package models

import common "BusinessServer/common/abstract/models"

const TableNameSwustSystemRole = "swust_system_role"

// SwustSystemRole mapped from table <swust_system_role>
type SwustSystemRole struct {
	common.DefaultModel
	Resource *string `gorm:"column:resource;type:varchar(50)" json:"resource"`    // 权限关联资源
	ParentID *string `gorm:"column:parent_id;type:varchar(255)" json:"parent_id"` // 父节点id
}

type UserRoleModel struct {
	UserId string   `gorm:"user_id" json:"userId"`    //用户ID
	Roles  []string `gorm:"column:role" json:"roles"` //权限列表
}

type SystemRoleVO struct {
	common.DefaultModel
	Path *string `gorm:"column:resource;type:varchar(50)" json:"path"` // 权限关联资源
	Name string  `gorm:"column:name" json:"name"`
	Meta struct {
		Title string `gorm:"column:title" json:"title"`
		Icon  string `gorm:"column:icon" json:"icon"`
	} `gorm:"embedded" json:"meta"`
	ParentID *string        `gorm:"column:parent_id;type:varchar(255)" json:"parent_id"` // 父节点id
	Children []SystemRoleVO `gorm:"-" json:"children"`                                   //子菜单
}

// TableName SwustSystemRole's table name
func (*SwustSystemRole) TableName() string {
	return TableNameSwustSystemRole
}
