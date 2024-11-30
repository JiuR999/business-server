package service

import (
	dao2 "BusinessServer/apps/order/dao"
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

const (
	APPLYING       = 0  //申请中
	ORDERING       = 1  //采购中
	ORDER_FINISHED = 2  //采购完成
	ORDER_REFUSE   = -1 //拒绝
)

// CountAssetsByType 根据资产类型统计
func (a *statisticService) CountAssetsByType() (res []models.StatisticModel, err common.SwustError) {
	record, err := dao.GetStatisticDao().CountAssetsByType()
	if err != nil {
		return res, err
	}
	return record, nil
}

// CountAssetsByStatus 根据资产状态统计
func (a *statisticService) CountAssetsByStatus() (res []models.StatisticModel, err common.SwustError) {
	record, err := dao.GetStatisticDao().CountAssetsByStatus()
	if err != nil {
		return res, err
	}
	return record, nil
}

// CountAssetsApplyTrend 获取资产申请趋势统计
func (a *statisticService) CountAssetsApplyTrend() (res []models.StatisticModel, err common.SwustError) {
	record, err := dao.GetStatisticDao().CountAssetsApplyTrend()
	if err != nil {
		return res, err
	}
	return record, nil
}

// CountAssetsDepTrend 获取资产报废趋势统计
func (a *statisticService) CountAssetsDepTrend() (res []models.StatisticModel, err common.SwustError) {
	record, err := dao.GetStatisticDao().CountAssetsDepTrend()
	if err != nil {
		return res, err
	}
	return record, nil
}

func (a *statisticService) CountOrderDeTail() (res models.StatisticModel, err common.SwustError) {
	orderTotal, swustError := dao2.GetOrderDao().CountByStatus(APPLYING)
	if swustError != nil {
		return res, swustError
	}
	res = models.StatisticModel{
		Name:  "申请总数",
		Value: int(orderTotal),
	}

	return res, nil
}
