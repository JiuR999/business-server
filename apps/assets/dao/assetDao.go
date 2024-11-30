package dao

import (
	model "BusinessServer/apps/assets/models"
	"BusinessServer/common"
	abstract "BusinessServer/common/abstract/dao"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type assetDao struct {
	abstract.Dao
}

var assetDaoInstance = new(assetDao)

func init() {
	assetDaoInstance.Init()
	assetDaoInstance.Model = model.AssetsModel{}
}

func GetAssetDao() *assetDao {
	return assetDaoInstance
}

func (d *assetDao) Page(req any, record any) (int64, common.SwustError) {
	tx := d.Gm.Table("swust_asset a").
		Select("a.id,a.name,a.code,a.user_id,a.production_time,a.service_length,a.price,a.status," +
			"a.comment,a.update_time,u.name header,quantity," +
			"t.name type_name,p.name producer_name," +
			"CASE WHEN a.production_time IS NOT NULL THEN DATE_ADD(a.production_time, INTERVAL COALESCE(a.service_length,0) year) ELSE NULL END AS retirement_time")
	conds := req.(*model.AssetsQueryRequest)

	buildBasicCondition(conds, tx)

	tx.Joins("LEFT JOIN swust_system_user u ON u.id = a.user_id").
		Joins("LEFT JOIN swust_asset_type t ON a.type_id = t.id").
		Joins("LEFT JOIN swust_producer p ON p.id = a.producer_id").
		Joins("LEFT JOIN (SELECT a.id,SUM(quantity) quantity " +
			"FROM swust_asset a " +
			"LEFT JOIN swust_warehouse wh " +
			"ON a.id = wh.asset_id " +
			"GROUP BY a.id) wh on a.id = wh.id")
	var total int64
	tx.Count(&total)
	tx.Limit(conds.PageSize).
		Offset((conds.PageNum - 1) * conds.PageSize).
		Find(record)
	if err := tx.Error; err != nil {
		return 0, common.NewDaoError(err.Error())
	}
	return total, nil
}

func (d *assetDao) List(req any, record any) common.SwustError {
	tx := d.Gm.Table("swust_asset a").
		Select("a.id,a.name,a.code,a.user_id,a.production_time,a.service_length,a.price,a.status," +
			"a.comment,a.update_time,u.name header,quantity," +
			"t.name type_name,p.name producer_name," +
			"CASE WHEN a.production_time IS NOT NULL THEN DATE_ADD(a.production_time, INTERVAL COALESCE(a.service_length,0) year) ELSE NULL END AS retirement_time")
	conds := req.(*model.AssetsQueryRequest)

	buildBasicCondition(conds, tx)

	tx.Joins("LEFT JOIN swust_system_user u ON u.id = a.user_id").
		Joins("LEFT JOIN swust_asset_type t ON a.type_id = t.id").
		Joins("LEFT JOIN swust_producer p ON p.id = a.producer_id").
		Joins("LEFT JOIN (SELECT a.id,SUM(quantity) quantity " +
			"FROM swust_asset a " +
			"LEFT JOIN swust_warehouse wh " +
			"ON a.id = wh.asset_id " +
			"GROUP BY a.id) wh on a.id = wh.id")
	var total int64
	tx.Count(&total)
	tx.Find(record)
	if err := tx.Error; err != nil {
		return common.NewDaoError(err.Error())
	}
	return nil
}

func buildBasicCondition(conds *model.AssetsQueryRequest, tx *gorm.DB) {
	if strings.TrimSpace(conds.Name) != "" {
		tx.Where("a.name like ?", "%"+conds.Name+"%")
	}
	if conds.Status != nil {
		tx.Where("a.status = ?", conds.Status)
	}
	if conds.TypeID != nil {
		tx.Where("a.type_id = ?", conds.TypeID)
	}
	if conds.ProducerID != nil {
		tx.Where("a.producer_id = ?", conds.ProducerID)
	}
	if conds.CreateTime != nil && conds.EndTime != nil {
		tx.Where("a.create_time > ? AND a.create_time < ?", conds.CreateTime, conds.EndTime)
	}
	if len(conds.Ids) > 0 {
		tx.Where("a.id IN ?", conds.Ids)
	}
}

func (d *assetDao) AddOrUpdate(record any) common.SwustError {
	tx := d.Gm.Model(d.Model).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "code"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "user_id", "production_time", "service_length", "price",
			"status", "comment", "update_time", "producer_id", "type_id"}),
	}).Create(record)
	if e := tx.Error; e != nil {
		return common.NewDaoError(e.Error())
	}
	return nil
}

func (d *assetDao) Deprecated(ids []string) common.SwustError {
	err := d.Gm.Model(d.Model).
		Where("id IN ?", ids).
		Update("status", 3).
		Error
	if err != nil {
		return common.NewDaoError(err.Error())
	}
	return nil
}
