package dao

import (
	"BusinessServer/apps/room/models"
	"BusinessServer/common"
	abstract "BusinessServer/common/abstract/dao"
	common2 "BusinessServer/common/abstract/models"
)

type roomDao struct {
	abstract.Dao
}

var roomDaoInstance = new(roomDao)

func init() {
	roomDaoInstance.Init()
	roomDaoInstance.Model = models.SwustRoomModel{}
}

func GetRoomDao() *roomDao {
	return roomDaoInstance
}

func (d *roomDao) Page(req any, record any) (int64, common.SwustError) {
	tx := d.Gm.Model(d.Model)
	conds := req.(*common2.PageModel)

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

func (d *roomDao) ListLocation(record any) common.SwustError {
	tx := d.Gm.Model(d.Model)
	//conds := req.(*abstract2.PageModel)
	tx.Select("location").
		Group("location").
		Find(record)
	if err := tx.Error; err != nil {
		return common.NewDaoError(err.Error())
	}
	return nil
}
