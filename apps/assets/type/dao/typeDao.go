package dao

import (
	"BusinessServer/apps/assets/type/models"
	"BusinessServer/common"
	abstract "BusinessServer/common/abstract/dao"
	"BusinessServer/common/utils"
)

type typeDao struct {
	abstract.Dao
}

var typeDaoInstance = new(typeDao)

func init() {
	typeDaoInstance.Init()
	typeDaoInstance.Model = models.SwustAssetType{}
}

func GetTypeDao() *typeDao {
	return typeDaoInstance
}

func (d *typeDao) Page(req any, record any) (int64, common.SwustError) {
	tx := d.Gm.Model(d.Model)
	conds := req.(*models.TypeQueryRequest)
	var total int64
	if !utils.IsBlank(conds.Name) {
		tx.Where("name Like ?", "%"+*conds.Name+"%")
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

func (d *typeDao) GetTypeList() (res []string, err common.SwustError) {
	tx := d.Gm.Model(d.Model).Select("name").Find(&res)
	if e := tx.Error; e != nil {
		return res, common.NewDaoError(e.Error())
	}
	return res, nil
}

func (d *typeDao) GetIdByName(typeName string) (id string, err common.SwustError) {
	tx := d.Gm.Model(d.Model).
		Select("id").
		Where("name = ?", typeName).
		Take(&id)
	if e := tx.Error; e != nil {
		return "", common.NewDaoError(e.Error())
	}
	return id, nil
}
