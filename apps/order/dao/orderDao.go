package dao

import (
	"BusinessServer/apps/order/models"
	"BusinessServer/common"
	abstract "BusinessServer/common/abstract/dao"
)

type orderDao struct {
	abstract.Dao
}

var orderDaoInstance = new(orderDao)

func init() {
	orderDaoInstance.Init()
	orderDaoInstance.Model = models.SwustOrder{}
}

func GetOrderDao() *orderDao {
	return orderDaoInstance
}

func (d *orderDao) Page(req any, record any) (int64, common.SwustError) {
	tx := d.Gm.Model(d.Model).
		Table("swust_order o").
		Select("o.*,u.name apply_user").
		Joins("LEFT JOIN swust_system_user u ON u.id = o.user_id")
	conds := req.(*models.OrderQueryRequest)
	var total int64
	if conds.StartTime != nil &&
		!conds.StartTime.IsZero() &&
		conds.EndTime != nil &&
		!conds.EndTime.IsZero() {
		tx.Where("o.apply_time between ? and ?", conds.StartTime, conds.EndTime)
	}
	if conds.Status != nil {
		tx.Where("o.status  = ?", conds.Status)
	}
	if conds.ApplyUser != "" {
		tx.Where("u.name Like ?", "%"+conds.ApplyUser+"%")
	}

	tx.Count(&total)

	tx.Limit(conds.PageSize).
		Offset((conds.PageNum - 1) * conds.PageSize).
		Find(record)
	if err := tx.Error; err != nil {
		return 0, common.NewDaoError(err.Error())
	}
	return total, nil
}

func (d *orderDao) AddOrderDetailBatch(req []models.OrderAssetRequest) common.SwustError {
	err := d.Gm.Table("swust_order_assets").Create(req).Error
	if err != nil {
		return common.NewDaoError(err.Error())
	}
	return nil
}
