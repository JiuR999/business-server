package dao

import (
	"BusinessServer/apps/system/log/models"
	"BusinessServer/common"
	"BusinessServer/common/Time"
	abstract "BusinessServer/common/abstract/dao"
)

type logDao struct {
	abstract.Dao
}

var logDaoInstance = new(logDao)

func init() {
	logDaoInstance.Init()
	logDaoInstance.Model = models.SwustSystemLog{}
}
func GetLogDao() *logDao {
	return logDaoInstance
}

func (d *logDao) Page(req any, record any) (int64, common.SwustError) {
	request := req.(*models.SystemLogQueryRequest)
	var total int64
	tx := d.Gm.Model(d.Model).
		Table("swust_system_log l").
		Select("l.id,l.content,l.operate_time,l.event,su.name operator_name")

	if request.Event != 0 {
		tx.Where("event = ?", request.Event)
	}
	if request.OperateUser != "" {
		tx.Where("user_id = ?", request.OperateUser)
	}
	if request.StartTime != nil &&
		(*request.StartTime).String() != Time.NIL_TIME &&
		request.EndTime != nil &&
		(*request.EndTime).String() != Time.NIL_TIME {
		tx.Where("operate_time between ? and ?", request.StartTime, request.EndTime)
	}
	tx.Count(&total)

	tx.Joins("LEFT JOIN swust_system_user su ON su.id = l.user_id ").
		Limit(request.PageSize).
		Offset((request.PageNum - 1) * request.PageSize).
		Find(record)
	if err := tx.Error; err != nil {
		return 0, common.NewDaoError(err.Error())
	}
	return total, nil
}

func (d *logDao) GetLogList() (res []string, err common.SwustError) {
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
