package dao

import (
	"BusinessServer/apps/system/role/models"
	"BusinessServer/common"
	abstract "BusinessServer/common/abstract/dao"
	"fmt"
)

type roleDao struct {
	abstract.Dao
}

var roleDaoInstance = new(roleDao)

func init() {
	roleDaoInstance.Init()
	roleDaoInstance.Model = models.SwustSystemRole{}
}
func GetRoleDao() *roleDao {
	return roleDaoInstance
}

func (d *roleDao) Page(req any, record any) (int64, common.SwustError) {
	//request := req.(*common2.PageModel)
	var total int64
	tx := d.Gm.Model(d.Model).
		Table("swust_system_role r").
		Select("r.id,r.name,r.resource,r.parent_id").
		Where("parent_id != ?", "0").
		Find(record)
	if err := tx.Error; err != nil {
		return 0, common.NewDaoError(err.Error())
	}
	return total, nil
}

func (d *roleDao) GetRolesByUserId(userId string) ([]models.SystemRoleVO, common.SwustError) {
	var res []models.SystemRoleVO
	err := d.Gm.Raw(`WITH RECURSIVE parent_menu AS (
	SELECT
		r.* 
	FROM
		swust_system_role r
		JOIN swust_user_role ur ON JSON_CONTAINS(
			ur.roles,
		CONCAT( '"', r.id, '"' )) 
	AND
		ur.user_id = ? UNION ALL
	SELECT
		r2.* 
	FROM
		swust_system_role r2
		INNER JOIN parent_menu p ON r2.id = p.parent_id 
	) SELECT DISTINCT
	* 
FROM
	parent_menu
Order By id
`, userId).Find(&res).Error
	if err != nil {
		return res, common.NewDaoError(err.Error())
	}
	return res, nil
}

func (d *roleDao) UpdateByUserId(userId string, roles []string) common.SwustError {
	//查询用Raw 必须搭配Scan 增删改用Exec
	tx := d.Gm.Exec(`
UPDATE swust_user_role
SET roles = JSON_SET(roles, '$', JSON_ARRAY ?)
WHERE user_id = ?`, roles, userId)
	fmt.Println(tx.RowsAffected)
	if tx.RowsAffected <= 0 {
		return common.NewDaoError("更新失败！")
	}
	return nil
}

func (d *roleDao) Add2UserRole(req *models.UserRoleModel) common.SwustError {
	tx := d.Gm.Exec(`
		INSERT INTO swust_user_role 
		VALUES(?,JSON_ARRAY ?)`, req.UserId, req.Roles)
	if err := tx.Error; err != nil {
		return common.NewDaoError(err.Error())
	}
	return nil
}
