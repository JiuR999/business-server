package service

import (
	"BusinessServer/apps/system/role/dao"
	"BusinessServer/apps/system/role/models"
	"BusinessServer/common"
	common2 "BusinessServer/common/abstract/models"
	"github.com/gin-gonic/gin"
)

type roleService struct {
}

var rs = new(roleService)

func GetRoleService() *roleService {
	return rs
}

func (s *roleService) Add(context *gin.Context, req any) (res any, err common.SwustError) {

	return res, nil
}

func (s *roleService) Update(context *gin.Context, req any) (err common.SwustError) {
	return nil
}
func (s *roleService) GetById(context *gin.Context, model any) (err common.SwustError) {
	return nil
}

func (s *roleService) DeleteByIds(context *gin.Context, ids []string) (res int64, err common.SwustError) {

	return 0, nil

}

func (s *roleService) Del(context *gin.Context, ids []string) (int64, common.SwustError) {
	return 0, nil

}

func (s *roleService) Page(context *gin.Context, req any) (res common2.PageResponseModel, err common.SwustError) {

	var record []models.SwustSystemRole
	total, err := dao.GetRoleDao().Page(req, &record)
	if err != nil {
		return res, common.NewServiceError(err.GetMsg())
	}
	res = common2.PageResponseModel{
		Total: int(total),
		Data:  record,
	}
	return res, nil
}

func (s *roleService) GetRolesByUserId(userId string) ([]models.SystemRoleVO, common.SwustError) {
	records, err := dao.GetRoleDao().GetRolesByUserId(userId)
	if err != nil {
		return records, common.NewServiceError(err.GetMsg())
	}
	var res []models.SystemRoleVO
	for i, v := range records {

		for _, child := range records {
			if *child.ParentID == v.Id {
				if records[i].Children == nil {
					records[i].Children = make([]models.SystemRoleVO, 0)
				}
				records[i].Children = append(records[i].Children, child)
			}
		}

	}

	for _, record := range records {
		if *record.ParentID == "0" {
			res = append(res, record)
		}
	}
	return res, nil
}

// 修改权限
func (s *roleService) ModifyRole(userId string, roles []string) common.SwustError {
	err := dao.GetRoleDao().UpdateByUserId(userId, roles)
	if err != nil {
		return common.NewServiceError(err.GetMsg())
	}
	return nil
}

// 像用户权限表中添加数据
func (s *roleService) Add2UserRole(req *models.UserRoleModel) common.SwustError {
	return dao.GetRoleDao().Add2UserRole(req)
}
