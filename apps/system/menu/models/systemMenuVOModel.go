package models

type SwustMenuVO struct {
	SwustMenuModel
	Children []SwustMenuVO `gorm:"-" yaml:"children" json:"children"` //子节点
}
