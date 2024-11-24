package service

import (
	"BusinessServer/apps/statistic/dao"
	"BusinessServer/apps/statistic/models"
	"BusinessServer/common"
)

type statisticService struct {
}

var rService = new(statisticService)

func GetStatisticService() *statisticService {
	return rService
}

/*
根据资产类型统计
*/
func (a *statisticService) CountAssetsByType() (res []models.AssetTypeModel, err common.SwustError) {
	record, err := dao.GetStatisticDao().CountAssetsByType()
	if err != nil {
		return res, err
	}
	return record, nil
}
