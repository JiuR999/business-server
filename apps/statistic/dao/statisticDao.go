package dao

import (
	"BusinessServer/apps/statistic/models"
	"BusinessServer/common"
	abstract "BusinessServer/common/abstract/dao"
)

type statisticDao struct {
	abstract.Dao
}

func (d statisticDao) CountAssetsByType() (record []models.AssetTypeModel, err common.SwustError) {
	tx := d.Gm.Raw(`
		SELECT SUM(COALESCE(number,0)) quantity,t.name type_name 
		FROM swust_asset a 
		LEFT JOIN swust_asset_type t 
		ON a.type_id = t.id
		LEFT JOIN swust_warehouse wh ON wh.asset_id = a.id 
		GROUP BY type_name
	`)
	if e := tx.Error; e != nil {
		return record, common.NewDaoError(e.Error())
	}
	tx.Scan(&record)
	return record, nil
}

var statisticDaoInstance = new(statisticDao)

func init() {
	statisticDaoInstance.Init()
	statisticDaoInstance.Model = models.AssetTypeModel{}
}

func GetStatisticDao() *statisticDao {
	return statisticDaoInstance
}
