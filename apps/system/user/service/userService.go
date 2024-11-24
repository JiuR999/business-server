package service

import (
	logModels "BusinessServer/apps/system/log/models"
	"BusinessServer/apps/system/log/service"
	roleModels "BusinessServer/apps/system/role/models"
	roleService "BusinessServer/apps/system/role/service"
	"BusinessServer/apps/system/user/dao"
	"BusinessServer/apps/system/user/models"
	"BusinessServer/common"
	common2 "BusinessServer/common/abstract/models"
	publisher "BusinessServer/common/services"
	"BusinessServer/common/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type userService struct {
}

var uService = new(userService)

func GetUserService() *userService {
	return uService
}

func (a *userService) Add(context *gin.Context, req any) (res any, err common.SwustError) {
	userModel := req.(*models.SystemUserModel)
	userModel.SetNewId()
	salt := utils.GenerateSalt(16)
	userModel.Salt = &salt
	digest := utils.Md5Digest(*userModel.Password, salt)
	userModel.Password = &digest
	res, err = dao.GetUserDao().Add(userModel)
	if err != nil {
		return nil, common.NewServiceError(err.GetMsg())
	}
	service.WriteLog(context, common.LOG_EVENT_ADD, fmt.Sprintf("增加用户-%s", userModel.Name))
	//为用户给予初始权限
	defaultRoles := []string{"2", "6", "9", "11", "16"}
	roleReq := &roleModels.UserRoleModel{
		UserId: userModel.Id,
		Roles:  defaultRoles,
	}
	roleService.GetRoleService().Add2UserRole(roleReq)
	return res, nil
}

func (a *userService) Update(context *gin.Context, req any) (err common.SwustError) {
	err = dao.GetUserDao().Update(req)
	if err != nil {
		return common.NewServiceError(err.GetMsg())
	}
	return nil
}
func (a *userService) GetById(context *gin.Context, model any) (err common.SwustError) {
	id := context.Query("id")
	if id != "" {
		swustError := dao.GetUserDao().GetById(id, model)
		if swustError != nil {
			return common.NewServiceError(swustError.GetMsg())
		}
		return nil
	}
	return common.NewServiceError("查询ID不能为空！")
}

func (a *userService) DeleteByIds(context *gin.Context, ids []string) (res int64, err common.SwustError) {

	if len(ids) < 0 {
		return 0, common.NewServiceError("请输入待删除资产ID")
	}
	affects, err := a.Del(context, ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	return affects, nil

}

func (a *userService) Del(context *gin.Context, ids []string) (int64, common.SwustError) {
	affects, err := dao.GetUserDao().Delete(ids)
	if err != nil {
		return 0, common.NewServiceError(err.GetMsg())
	}
	return affects, nil

}

func (a *userService) Page(context *gin.Context, req any) (res common2.PageResponseModel, err common.SwustError) {
	request := req.(*common2.PageModel)
	request.IfAbsent()
	var record []models.SystemUserModel
	total, err := dao.GetUserDao().Page(request, &record)
	if err != nil {
		return res, common.NewServiceError(err.GetMsg())
	}
	res = common2.PageResponseModel{
		Total: int(total),
		Page:  *request,
		Data:  record,
	}
	return res, nil
}

func (a *userService) Login(req models.LoginRequest) (string, common.SwustError) {
	var record models.SystemUserModel
	err := dao.GetUserDao().GetByCondition(req, &record)
	if err != nil {
		return "", common.NewServiceError(err.GetMsg())
	}
	digest := utils.Md5Digest(req.Password, *record.Salt)
	if digest != *record.Password {
		return "", common.NewServiceError("用户名或密码有误!")
	}
	jwt, err2 := utils.GenerateJWT(*record.Account, record.Id)
	if err2 != nil {
		return "", common.NewServiceError("生成Token失败 请重试!")
	}
	claims, _ := utils.ParseToken(jwt)
	log := logModels.NewSystemLog(common.LOG_EVENT_LOGIN, claims.UserId, claims.UserName+"登陆系统")
	publisher.EB.Publish(common.EVENT_LOG, publisher.EventModel{
		Event: "写入日志",
		Data:  log,
	})
	fmt.Println(claims)
	//TODO 存Redis
	return jwt, nil
}

func (a *userService) GetCurrentUser(context *gin.Context) (res *models.SystemUserModel, err common.SwustError) {
	token := context.Request.Header.Get("token")
	if strings.TrimSpace(token) == "" {
		return nil, common.NewServiceError("用户未登录!")
	}
	claims, e := utils.ParseToken(token)
	if e != nil {
		return nil, common.NewServiceError(e.Error())
	}
	context.Set(common.COMMON_AUTH_CURRENT, &utils.CurrentUser{
		ID:   claims.UserId,
		Name: claims.UserName,
	})
	record := &models.SystemUserModel{}
	err = dao.GetUserDao().GetById(claims.UserId, record)
	if err != nil {
		return nil, common.NewServiceError(err.GetMsg())
	}
	return record, nil
}

func (a *userService) GetUserList() ([]string, common.SwustError) {
	return dao.GetUserDao().GetUserList()
}
