package dao

import (
	"BusinessServer/apps/statistic/models"
	"BusinessServer/common"
	abstract "BusinessServer/common/abstract/dao"
)

type statisticDao struct {
	abstract.Dao
}

func (d *statisticDao) CountAssetsByType() (record []models.StatisticModel, err common.SwustError) {
	tx := d.Gm.Raw(`SELECT COUNT(a.id) value,COALESCE(t.name,"未知") name
			FROM swust_asset a 
			LEFT JOIN swust_asset_type t 
			ON a.type_id = t.id
			GROUP BY name
		`)
	if e := tx.Error; e != nil {
		return record, common.NewDaoError(e.Error())
	}
	tx.Scan(&record)
	return record, nil
}

func (d *statisticDao) CountAssetsByStatus() (record []models.StatisticModel, err common.SwustError) {
	tx := d.Gm.Raw(`SELECT
	COUNT( a.id ) value,
CASE
	STATUS 
		WHEN 0 THEN
		'在用' 
		WHEN 1 THEN
		'故障' 
		WHEN 2 THEN
		'维修' 
		WHEN 3 THEN
		'报废' ELSE '闲置' 
	END AS name 
FROM
	swust_asset a
	LEFT JOIN swust_asset_type t ON a.type_id = t.id 
GROUP BY
STATUS`).
		Scan(&record)
	if e := tx.Error; e != nil {
		return record, common.NewDaoError(e.Error())
	}
	return record, nil
}

func (d *statisticDao) CountAssetsApplyTrend() (record []models.StatisticModel, err common.SwustError) {
	tx := d.Gm.Raw(`WITH RECURSIVE date_series AS (
SELECT CURRENT_DATE AS date
UNION ALL 
SELECT DATE_SUB(date,INTERVAL 1 day)
FROM date_series
WHERE date > DATE_SUB(CURRENT_DATE, INTERVAL 6 DAY)
)

SELECT date name,COUNT(a.name) value
FROM date_series
LEFT JOIN swust_asset a
ON DATE(a.create_time) = date
GROUP BY date
ORDER BY date`).Find(&record)
	if e := tx.Error; e != nil {
		return record, common.NewDaoError(e.Error())
	}
	return record, nil
}

func (d *statisticDao) CountAssetsDepTrend() (record []models.StatisticModel, err common.SwustError) {
	tx := d.Gm.Raw(`WITH RECURSIVE month_series AS (
SELECT 1 AS date
UNION ALL 
SELECT date+1
FROM month_series
WHERE date+1 <= MONTH(CURRENT_DATE)
)

SELECT date name,COUNT(a.name) value
FROM month_series
LEFT JOIN swust_asset a
ON MONTH(DATE(a.update_time)) = date
AND a.status = 3
GROUP BY date
ORDER BY date`).
		Find(&record)
	if e := tx.Error; e != nil {
		return record, common.NewDaoError(e.Error())
	}
	return record, nil
}

var statisticDaoInstance = new(statisticDao)

func init() {
	statisticDaoInstance.Init()
	statisticDaoInstance.Model = models.StatisticModel{}
}

func GetStatisticDao() *statisticDao {
	return statisticDaoInstance
}
