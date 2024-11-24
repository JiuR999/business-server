package service

import (
	"BusinessServer/apps/system/menu/dao"
	"BusinessServer/apps/system/menu/models"
	"BusinessServer/common"
	common2 "BusinessServer/common/abstract/models"
	"github.com/gin-gonic/gin"
)

type systemMenuService struct {
}

var menuService = new(systemMenuService)

func GetSystemMenuService() *systemMenuService {
	return menuService
}

func (s systemMenuService) Add(context *gin.Context, req any) (res any, err common.SwustError) {
	//TODO implement me
	panic("implement me")
}

func (s systemMenuService) DeleteByIds(context *gin.Context, ids []string) (affects int64, err common.SwustError) {
	//TODO implement me
	panic("implement me")
}

func (s systemMenuService) Update(context *gin.Context, req any) (err common.SwustError) {
	//TODO implement me
	panic("implement me")
}

func (s systemMenuService) GetById(context *gin.Context, model any) (err common.SwustError) {
	//TODO implement me
	panic("implement me")
}

func (s systemMenuService) Page(context *gin.Context, req any) (res common2.PageResponseModel, err common.SwustError) {
	//TODO implement me
	panic("implement me")
}

func (s systemMenuService) List() ([]models.SwustMenuVO, common.SwustError) {
	list, err := dao.GetSystemMenuDao().List()
	if err != nil {
		return nil, common.NewServiceError(err.GetMsg())
	}
	buildMenuTree(list)
	return list, nil
}

func buildMenuTree(list []models.SwustMenuVO) {
	for i, v := range list {
		for _, v2 := range list {
			if *v2.ParentID == v.Id {
				if list[i].Children == nil {
					list[i].Children = make([]models.SwustMenuVO, 0)
				}
				list[i].Children = append(list[i].Children, v2)
			}
		}
	}
}
