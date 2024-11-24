package dao

import (
	"BusinessServer/apps/system/menu/models"
	"BusinessServer/common"
	abstract "BusinessServer/common/abstract/dao"
)

type systemMenuDao struct {
	abstract.Dao
}

var menuDao = new(systemMenuDao)

func init() {
	menuDao.Init()
	menuDao.Model = models.SwustMenuModel{}
}

func GetSystemMenuDao() *systemMenuDao {
	return menuDao
}

func (d systemMenuDao) List() (res []models.SwustMenuVO, err common.SwustError) {
	tx := d.Gm.Raw(`WITH RECURSIVE sub_menu AS (
	SELECT
		id,
		name,
		router,
		level,
		parent_id,
		role 
	FROM
		swust_system_menu 
	WHERE
		parent_id = 0 UNION ALL
	SELECT
		m.id,
		m.name,
		m.router,
		m.level,
		m.parent_id,
		m.role 
	FROM
		swust_system_menu m
		INNER JOIN sub_menu ON sub_menu.id = m.parent_id 
	) SELECT
	id,
	name,
	router,
	level,
	parent_id,
	role 
FROM
	sub_menu`)
	if err := tx.Error; err != nil {
		return nil, common.NewDaoError(err.Error())
	}
	tx.Find(&res)
	return res, nil
}
