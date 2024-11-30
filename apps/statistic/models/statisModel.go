package models

type StatisticModel struct {
	Name  string `gorm:"column:name" json:"name"`   //名称
	Value int    `gorm:"column:value" json:"value"` //数量
}
