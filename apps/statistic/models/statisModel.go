package models

type AssetTypeModel struct {
	//Name     *string `gorm:"column:name;type:varchar(25)" json:"name"` // 房间名称
	TypeName *string `gorm:"column:type_name" json:"typeName"` //类型名称
	Quantity int     `gorm:"column:quantity" json:"quantity"`  //数量
}
