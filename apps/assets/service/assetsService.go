package service

import (
	"BusinessServer/apps/assets/dao"
	model "BusinessServer/apps/assets/models"
	dao3 "BusinessServer/apps/assets/type/dao"
	"BusinessServer/apps/assets/type/service"
	dao4 "BusinessServer/apps/producer/dao"
	service2 "BusinessServer/apps/producer/service"
	logWriter "BusinessServer/apps/system/log/service"
	dao2 "BusinessServer/apps/system/user/dao"
	"BusinessServer/apps/system/user/models"
	service3 "BusinessServer/apps/system/user/service"
	"BusinessServer/common"
	"BusinessServer/common/Time"
	common2 "BusinessServer/common/abstract/models"
	eventService "BusinessServer/common/services"
	"BusinessServer/common/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

type assetsService struct {
}

var aService = new(assetsService)

func GetAssetsService() *assetsService {
	return aService
}

func (a *assetsService) Add(context *gin.Context, req any) (res any, err common.SwustError) {
	assetsModel := req.(*model.AssetsModel)
	assetsModel.SetNewId()
	logWriter.WriteLog(context, common.LOG_EVENT_UPDATE, "增加资产信息，设备编码为"+*assetsModel.Code)
	if *assetsModel.Code == "" || assetsModel.Name == "" {
		return nil, common.NewServiceError("资产编码或者资产名称不能为空")
	}
	res, err = dao.GetAssetDao().Add(assetsModel)
	if err != nil {
		return nil, common.NewServiceError(err.GetMsg())
	}
	return res, nil
}

func (a *assetsService) Update(context *gin.Context, req any) (err common.SwustError) {
	assetsModel := req.(*model.AssetsModel)
	logWriter.WriteLog(context, common.LOG_EVENT_UPDATE, "更新资产信息，设备编码为"+*assetsModel.Code)
	err = dao.GetAssetDao().Update(req)
	if err != nil {
		return common.NewServiceError(err.GetMsg())
	}
	return nil
}
func (a *assetsService) GetById(context *gin.Context, model any) (err common.SwustError) {
	id := context.Query("id")
	logWriter.WriteLog(context, common.LOG_EVENT_UPDATE, "查询资产信息，查询id为"+id)
	if id != "" {
		swustError := dao.GetAssetDao().GetById(id, model)
		if swustError != nil {
			return common.NewServiceError(swustError.GetMsg())
		}
		return nil
	}
	return common.NewServiceError("查询ID不能为空！")
}

func (a *assetsService) DeleteByIds(context *gin.Context, ids []string) (res int64, err common.SwustError) {
	if len(ids) < 0 {
		return 0, common.NewServiceError("请输入待删除资产ID")
	}
	affects, err := dao.GetAssetDao().Delete(ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	logWriter.WriteLog(context, common.LOG_EVENT_UPDATE, "删除资产信息，设备ID为"+strings.Join(ids, ","))
	return affects, nil

}

func (a *assetsService) Page(context *gin.Context, req any) (res common2.PageResponseModel, err common.SwustError) {
	request := req.(*model.AssetsQueryRequest)
	request.IfAbsent()
	var record []model.AssetsVO
	total, err := dao.GetAssetDao().Page(request, &record)

	if err != nil {
		return res, common.NewServiceError(err.GetMsg())
	}
	res = common2.PageResponseModel{
		Total: int(total),
		Page:  request.PageModel,
		Data:  record,
	}

	return res, nil
}

// 获取资产导入模板
func (a *assetsService) Template() (*excelize.File, common.SwustError) {
	file, swustError := utils.GetTemplate("AssetImportTemplate.xlsx")

	if swustError != nil {
		return nil, common.NewServiceError(swustError.GetMsg())
	}
	sheet := "Sheet1"

	utils.AddAllowBlank(file, sheet, false, fmt.Sprintf("%s2:%s1000", utils.ProductionTime, utils.ProductionTime))

	utils.SetExcelDropList(file, sheet, false, fmt.Sprintf("%s2:%s1000", utils.AssetStatus, utils.AssetStatus), []string{"在用", "故障", "维修", "报废", "闲置"})

	userList, swustError := service3.GetUserService().GetUserList()
	if swustError == nil {
		utils.SetExcelDropList(file, sheet, false, fmt.Sprintf("%s2:%s1000", utils.AssetHeader, utils.AssetHeader), userList)
	}
	typeList, swustError := service.GetTypeService().GetTypeList()
	if swustError == nil {
		utils.SetExcelDropList(file, sheet, false, fmt.Sprintf("%s2:%s1000", utils.AssetType, utils.AssetType), typeList)
	}

	pList, swustError := service2.GetProducerService().GetProducerList()
	if swustError == nil {
		utils.SetExcelDropList(file, sheet, false, fmt.Sprintf("%s2:%s1000", utils.Producer, utils.Producer), pList)
	}
	return file, nil
}

func (a *assetsService) Import(file *multipart.FileHeader) common.SwustError {
	f, err := file.Open()

	if err != nil {
		return common.NewServiceError(err.Error())
	}
	excel, err := excelize.OpenReader(f)
	if err != nil {
		return common.NewServiceError(err.Error())
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				msgBody := map[string]string{}
				msgBody["errMsg"] = "导入文件有误，请按照导入模板进行上传！"
				//common.NewResponse(ctx).ErrorWithMsg(err.(string))
				eventService.EB.Publish(common.COMMON_MSG, eventService.EventModel{
					Event: common.EVENT_IMPORT,
					Data:  msgBody,
				})
			}

		}()
		for _, sheet := range excel.GetSheetList() {

			rows, err := excel.GetRows(sheet)
			if err != nil {
				msgBody := map[string]string{}
				msgBody["errMsg"] = err.Error()
				//common.NewResponse(ctx).ErrorWithMsg(err.(string))
				eventService.EB.Publish(common.COMMON_MSG, eventService.EventModel{
					Event: common.EVENT_IMPORT,
					Data:  msgBody,
				})
			}
			rowLen := len(rows)
			for i := 2; i <= rowLen; i++ {
				msgBody := map[string]string{}
				cellIndex := strconv.Itoa(i)
				msgBody["rows"] = cellIndex
				err := updateByImport(excel, sheet, cellIndex)
				if err != nil {

					msgBody["errMsg"] = fmt.Sprintf("第%d行数据导入失败-%s", i, err.GetMsg())
					//log.Fatal("第", i, "行数据导入失败", err.GetMsg())
				}
				msgBody["progress"] = strconv.Itoa(i * 100 / rowLen)

				eventService.EB.Publish(common.COMMON_MSG, eventService.EventModel{
					Event: "EVENT_IMPORT",
					Data:  msgBody,
				})

			}
		}
	}()
	return nil
}

func (a *assetsService) Export(ids []string) (file *excelize.File, msgs []string) {
	var list []model.AssetsVO
	swustError := dao.GetAssetDao().List(&model.AssetsQueryRequest{Ids: ids}, &list)
	if swustError != nil {
		msgs = append(msgs, swustError.GetMsg())
		return nil, msgs
	}
	file, swustError = utils.GetTemplate("AssetImportTemplate.xlsx")
	if swustError != nil {
		msgs = append(msgs, swustError.GetMsg())
		return nil, msgs
	}
	sheet := "Sheet1"
	for i, v := range list {
		index := strconv.Itoa(i + 2)
		status := "未知"
		comment := ""
		tname := "/"
		pname := "/"
		//productionTime := "/"
		file.SetCellStr(sheet, utils.Name+index, v.Name)
		file.SetCellStr(sheet, utils.AssetCode+index, *v.Code)
		file.SetCellStr(sheet, utils.AssetHeader+index, v.Header)
		//if v.ProducerName != nil {
		//	productionTime = v.ProductionTime.String()
		//}
		file.SetCellStr(sheet, utils.ProductionTime+index, v.ProductionTime.String())
		file.SetCellInt(sheet, utils.ServiceLength+index, int(*v.ServiceLength))
		file.SetCellFloat(sheet, utils.Price+index, float64(*v.Price), -1, 64)

		str2IntMap := utils.GetStr2IntMap()
		for k, index := range str2IntMap {
			if index == int(*v.Status) {
				status = k
			}
		}
		file.SetCellStr(sheet, utils.AssetStatus+index, status)

		if v.TypeName != nil {
			tname = *v.TypeName
		}
		file.SetCellStr(sheet, utils.AssetType+index, tname)
		if v.ProducerName != nil {
			pname = *v.ProducerName
		}
		file.SetCellStr(sheet, utils.Producer+index, pname)
		if v.Comment != nil {
			comment = *v.Comment
		}
		file.SetCellStr(sheet, utils.Comment+index, comment)
	}

	return file, msgs
}

func updateByImport(excel *excelize.File, sheet string, rowNum string) common.SwustError {
	var (
		userId, typeId string
		err            common.SwustError
		productionTime Time.LocalDay
		serviceLength  int64
		price          float32
		status         int64
	)
	name := utils.GetCellValue(excel, sheet, utils.Name+rowNum)
	code := utils.GetCellValue(excel, sheet, utils.AssetCode+rowNum)
	header := utils.GetCellValue(excel, sheet, utils.AssetHeader+rowNum)

	if header != utils.CELL_NOT_VALUE {
		userInfo := strings.Split(header, ":")
		userId, err = dao2.GetUserDao().GetIdByNameAndPhone(models.SystemUserQueryRequest{
			Name:        userInfo[0],
			PhoneNumber: &userInfo[1],
		})
		if err != nil {
			//用户未找到
			return common.NewServiceError("负责人" + userInfo[0] + "匹配失败" + err.GetMsg())
		}
	}
	producTimeStr := utils.GetCellValue(excel, sheet, utils.ProductionTime+rowNum)
	if producTimeStr != utils.CELL_NOT_VALUE {
		parse, err2 := time.Parse(Time.LOCAL_DAY_FORMAT, producTimeStr)
		if err2 != nil {
			//productionTime = Time.LocalDay{}
		} else {
			productionTime = Time.LocalDay(parse)
		}

	}

	serviceLengthStr := utils.GetCellValue(excel, sheet, utils.ServiceLength+rowNum)
	if serviceLengthStr != utils.CELL_NOT_VALUE {
		serviceLength = utils.ConvertStr2Int64(serviceLengthStr)
	}
	priceStr := utils.GetCellValue(excel, sheet, utils.Price+rowNum)
	price = utils.ConvertStr2Float32(priceStr)
	statusNumber := utils.GetCellValue(excel, sheet, utils.AssetStatus+rowNum)
	status = int64(utils.GetStr2IntMap()[statusNumber])
	typeStr := utils.GetCellValue(excel, sheet, utils.AssetType+rowNum)
	typeId, err = dao3.GetTypeDao().GetIdByName(typeStr)
	if err != nil {
		return common.NewServiceError(err.GetMsg())
	}
	producerStr := utils.GetCellValue(excel, sheet, utils.Producer+rowNum)
	producerInfo := strings.Split(producerStr, ":")
	producerId, err := dao4.GetProducerDao().GetIdByNameAndAddr(producerInfo[0], producerInfo[1])
	if err != nil {
		return common.NewServiceError(err.GetMsg())
	}
	comment := utils.GetCellValue(excel, sheet, utils.Comment+rowNum)

	asset := &model.AssetsModel{}
	asset.Name = name
	asset.Code = &code
	asset.UserID = &userId
	asset.ProductionTime = &productionTime
	asset.ServiceLength = &serviceLength
	asset.Price = &price
	asset.Status = &status
	asset.TypeID = &typeId
	asset.ProducerID = &producerId
	asset.Comment = &comment
	asset.SetNewId()
	err = dao.GetAssetDao().AddOrUpdate(asset)
	if err != nil {
		return common.NewServiceError(err.GetMsg())
	}
	return nil
}
