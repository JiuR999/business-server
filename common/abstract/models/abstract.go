package common

import (
	"BusinessServer/common/utils"
	"strconv"
)

type Int64 int64

func (i *Int64) UnmarshalJSON(data []byte) error {
	num := string(data)
	res, err := strconv.ParseInt(num[1:len(num)-1], 10, 64)
	if err != nil {
		return err
	}
	*i = Int64(res)
	return nil
}

/*func (i *Int64) Value() (driver.Value, error) {

}

func (i *Int64) Scan(src interface{}) error  {

}*/

type Models interface {
	SetId(id string)
	GetId()
}

type IdModel struct {
	Id string `gorm:"column:id;primaryKey" json:"id"` // id
}

func (models *IdModel) SetId(id string) {
	models.Id = id
}

func (models *IdModel) GetId() string {
	return models.Id
}

func (models *IdModel) SetNewId() {
	id := utils.GenerateId()
	models.SetId(id)
}

type DefaultModel struct {
	IdModel
	Name string `gorm:"column:name;type:varchar(255)" json:"name"` //名称
}

func (models *DefaultModel) SetId(id string) {
	models.Id = id
}

func (models *DefaultModel) GetId() string {
	return models.Id
}
