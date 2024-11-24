package dao

import (
	"BusinessServer/apps/system/user/models"
	"BusinessServer/common"
	abstract "BusinessServer/common/abstract/dao"
	"BusinessServer/common/utils"
	"gorm.io/gorm"
)

type userDao struct {
	abstract.Dao
}

var userDaoInstance = new(userDao)

func init() {
	userDaoInstance.Init()
	userDaoInstance.Model = models.SystemUserModel{}
}
func GetUserDao() *userDao {
	return userDaoInstance
}

func (d userDao) GetByCondition(req models.LoginRequest, record *models.SystemUserModel) common.SwustError {
	tx := d.Gm.Model(d.Model).
		Select("id, name, account, password, salt").
		Where("account=?", req.Account).Take(record)
	if err := tx.Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return common.NewDaoError("用户未注册！")
		default:
			return common.NewDaoError(err.Error())
		}
	}
	return nil

}

func (d userDao) Login(record models.SystemUserModel) {
	tx := d.Gm.Model(d.Model).
		Select("id, account, password, salt").
		Where("account=?", record.Account)
	if err := tx.Error; err != nil {
		common.NewDaoError(err.Error())
	}
}

func (d userDao) GetUserList() (res []string, err common.SwustError) {
	tx := d.Gm.Model(d.Model).Select("name,phone_number")
	raws, e := tx.Rows()
	if e != nil {
		return res, common.NewDaoError(e.Error())
	}
	for raws.Next() {
		var name, phone_number string
		raws.Scan(&name, &phone_number)
		res = append(res, name+":"+phone_number)
	}
	return res, nil
}

// 格式为 张三:12333222222
func (d userDao) GetIdByNameAndPhone(req models.SystemUserQueryRequest) (id string, swustError common.SwustError) {

	tx := d.Gm.Model(d.Model).
		Select("id")
	if req.Name != "" {
		tx.Where("name = ?", req.Name, req.PhoneNumber).
			Take(&id)
	}
	if !utils.IsBlank(req.PhoneNumber) {
		tx.Where("phone_number = ?", req.PhoneNumber)
	}
	if e := tx.Error; e != nil {
		return "", common.NewDaoError(e.Error())
	}
	return id, nil
}
