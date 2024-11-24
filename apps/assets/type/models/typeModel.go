package models

import common "BusinessServer/common/abstract/models"

const TableNameSwustAssetType = "swust_asset_type"

// SwustAssetType mapped from table <swust_asset_type>
type SwustAssetType struct {
	common.DefaultModel
}

type TypeQueryRequest struct {
	common.PageModel
	Name *string `json:"name"` //类型名称
}

// TableName SwustAssetType's table name
func (*SwustAssetType) TableName() string {
	return TableNameSwustAssetType
}
