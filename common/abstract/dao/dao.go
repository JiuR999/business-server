package abstract

import (
	"BusinessServer/common"
	common2 "BusinessServer/common/abstract/models"
	"BusinessServer/common/db"
	"gorm.io/gorm"
)

type Dao struct {
	Gm    *gorm.DB
	Model any
}

func (d *Dao) Init() {
	d.Gm = db.Orm.DB()
}

func (d *Dao) Add(req any) (any, common.SwustError) {
	tx := d.Gm.Create(req)
	if err := tx.Error; err != nil {
		return req, common.NewDaoError(err.Error())
	}
	return req, nil
}

func (d *Dao) GetById(id string, record any) common.SwustError {
	tx := d.Gm.Where("id = ?", id).
		Select("id,name,account,avatar,phone_number").
		Take(record)
	if err := tx.Error; err != nil {
		return common.NewDaoError(err.Error())
	}
	return nil
}

func (d *Dao) Update(req any) common.SwustError {
	tx := d.Gm.Updates(req)
	if err := tx.Error; err != nil {
		return common.NewDaoError(err.Error())
	}
	return nil
}

func (d *Dao) Delete(ids []string) (int64, common.SwustError) {
	tx := d.Gm.Delete(d.Model, "id in ?", ids)
	if err := tx.Error; err != nil {
		return 0, common.NewDaoError(err.Error())
	}
	return tx.RowsAffected, nil
}

func (d *Dao) Page(req any, record any) (int64, common.SwustError) {
	request := req.(*common2.PageModel)
	var total int64
	tx := d.Gm.Model(d.Model)
	tx.Count(&total)
	tx.Limit(request.PageSize).
		Offset((request.PageNum - 1) * request.PageSize).
		Find(record)
	if err := tx.Error; err != nil {
		return 0, common.NewDaoError(err.Error())
	}
	return total, nil
}
