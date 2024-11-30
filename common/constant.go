package common

import (
	"BusinessServer/env"
)

/*
 * @Desc:所有静态定义
 * @author:zhangx
 * @version: v1.0.0
 */

var TokenMap = make(map[string]string)

// 采购订单状态
const (
	STATUS_PREPARE  = 0
	STATUS_REFUSE   = -1
	STATUS_ORDING   = 1
	STATUS_COMPLETE = 2

	APPROVE_REFUSE = -1
	APPROVE_AGREE  = 1
)

// 日志相关常量
const (
	LOG_EVENT_ADD = iota + 1
	LOG_EVENT_DELETE
	LOG_EVENT_UPDATE
	LOG_EVENT_READ
	LOG_EVENT_OA
	LOG_EVENT_LOGIN
)
const (
	COMMON_PROJECT_NAME = env.COMMON_PROJECT_NAME
	COMMON_AUTH_SALT    = "swust2021"
	COMMON_AUTH_CURRENT = "currentUser"

	STATUS_SUCCESS  = "操作成功！"
	REQUEST_SUCCESS = "请求成功！"
)

const (
	COMMON_MSG        = "event_common"
	EVENT_NOTICE      = "event_notice"
	EVENT_IMPORT      = "event_import"
	EVENT_ORDER       = "event_order"
	EVENT_FINISH_SAVE = "event_finish_save"
	EVENT_LOG         = "event_log"
	EVENT_ASYNC       = "event_async"
)

const (
	RETURN_FAILED          = 0  //失败
	RETURN_SUCCESS         = 1  //成功
	COMMON_STATUS_INIT     = 0  //初始的
	COMMON_STATUS_NEGATIVE = 1  //失效的
	COMMON_STATUS_POSITIVE = 2  //有效的
	COMMON_STATUS_DELETE   = -1 //标记删除的
)
