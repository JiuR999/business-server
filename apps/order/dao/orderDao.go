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

func (d *orderDao) GetById(id string, record any) common.SwustError {
	tx := d.Gm.Model(record).Table("swust_order o").
		Select("o.*,u1.name apply_user,u2.name approve_user").
		Where("o.id = ?", id).
		Joins("LEFT JOIN swust_system_user u1 ON u1.id = user_id").
		Joins("LEFT JOIN swust_system_user u2 ON u2.id = approve_id").
		Take(record)
	if err := tx.Error; err != nil {
		return common.NewDaoError(err.Error())
	}
	return nil
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

func (d *orderDao) GetOrderDetailById(id string) (result []models.OrderAssetVO, err common.SwustError) {
	e := d.Gm.Table("swust_order_assets oa").
		Select("oa.name,oa.order_id,oa.producer_id,oa.type_id,oa.apply_time,t.name type_name,oa.order_time,p.name producer_name,oa.quantity,oa.comment,oa.order_user").
		Joins("LEFT JOIN swust_asset_type t ON t.id = oa.type_id").
		Joins("LEFT JOIN swust_producer p ON p.id = oa.producer_id").
		Where("order_id = ?", id).
		Find(&result).Error
	if e != nil {
		return result, common.NewDaoError(e.Error())
	}
	return result, nil
}

func (d *orderDao) Approve(req models.ApproveReq) common.SwustError {
	err := d.Gm.Table("swust_order").
		Omit("id").
		Updates(req).Error
	if err != nil {
		return common.NewDaoError(err.Error())
	}
	return nil
}

func (d *orderDao) UpdateOrderDetail(req models.OrderAssetVO) common.SwustError {
	err := d.Gm.Table("swust_order_assets").
		Select("order_time,order_user,producer_id").
		Where("order_id = ?", req.OrderId).
		Updates(req).
		Error
	if err != nil {
		return common.NewDaoError(err.Error())
	}
	return nil
}

func (d *orderDao) UpdateStatus(status int, id string) common.SwustError {
	err := d.Gm.Table("swust_order").
		Where("id = ?", id).
		Update("status", status).
		Error
	if err != nil {
		return common.NewDaoError(err.Error())
	}
	return nil
}

func (d *orderDao) CountByStatus(status int) (total int64, e common.SwustError) {
	err := d.Gm.Table("swust_order").
		Where("status = ?", status).
		Count(&total).
		Error
	if err != nil {
		return 0, common.NewDaoError(err.Error())
	}
	return total, nil
}
