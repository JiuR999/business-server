package dao

import (
	"BusinessServer/apps/producer/models"
	"BusinessServer/common"
	abstract "BusinessServer/common/abstract/dao"
	"BusinessServer/common/utils"
)

type producerDao struct {
	abstract.Dao
}

var producerDaoInstance = new(producerDao)

func init() {
	producerDaoInstance.Init()
	producerDaoInstance.Model = models.SwustProducer{}
}

func GetProducerDao() *producerDao {
	return producerDaoInstance
}

func (d *producerDao) Page(req any, record any) (int64, common.SwustError) {
	tx := d.Gm.Model(d.Model)
	conditions := req.(*models.ProducerQueryRequest)
	var total int64
	if !utils.IsBlank(conditions.Name) {
		tx.Where("name like ?", "%"+*conditions.Name+"%")
	}
	if !utils.IsBlank(conditions.ContactUser) {
		tx.Where("contact_user like ?", "%"+*conditions.ContactUser+"%")
	}
	if !utils.IsBlank(conditions.PhoneNumber) {
		tx.Where("contact_phone like ?", *conditions.PhoneNumber+"%")
	}
	if !utils.IsBlank(conditions.Address) {
		tx.Where("address like ?", "%"+*conditions.Address+"%")
	}
	tx.Count(&total)
	tx.Limit(conditions.PageSize).
		Offset((conditions.PageNum - 1) * conditions.PageSize).
		Find(record)
	if err := tx.Error; err != nil {
		return 0, common.NewDaoError(err.Error())
	}
	return total, nil
}

func (d *producerDao) ListLocation(record any) common.SwustError {
	tx := d.Gm.Model(d.Model)
	tx.Select("location").
		Group("location").
		Find(record)
	if err := tx.Error; err != nil {
		return common.NewDaoError(err.Error())
	}
	return nil
}

func (d *producerDao) GetProducerList() (res []string, err common.SwustError) {
	rows, e := d.Gm.Model(d.Model).Select("name,address").Rows()
	if e != nil {
		return res, common.NewDaoError(e.Error())
	}
	for rows.Next() {
		var name, address string
		e := rows.Scan(&name, &address)
		if e != nil {
			return res, common.NewDaoError(e.Error())
		}
		res = append(res, name+":"+address)
	}
	return res, nil
}

func (d *producerDao) GetIdByNameAndAddr(name string, addr string) (id string, err common.SwustError) {
	tx := d.Gm.Model(d.Model).
		Select("id").
		Where("name = ? AND address = ? ", name, addr).
		Take(&id)
	if e := tx.Error; e != nil {
		return "", nil
	}
	return id, nil
}

func (d *producerDao) GetProducerIds() (res []string, err common.SwustError) {
	e := d.Gm.Model(d.Model).Select("id").Find(&res).Error
	if e != nil {
		return res, common.NewDaoError(e.Error())
	}
	return res, nil
}
