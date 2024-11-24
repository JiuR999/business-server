package middleware

import (
	"BusinessServer/common"
	common2 "BusinessServer/common/abstract/models"
	"BusinessServer/common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("token")
		if strings.TrimSpace(token) == "" {
			common2.NewResponse(context).ErrorWithMsg("用户未登录!")
			context.Abort()
		} else {
			if _, ok := common.TokenMap[token]; ok {
				common2.NewResponse(context).ErrorWithCode(http.StatusUnauthorized, "登录信息已失效!")
				context.Abort()
			}
			claims, err := utils.ParseToken(token)
			if err != nil {
				common2.NewResponse(context).ErrorWithCode(http.StatusUnauthorized, "登录信息已过期!")
				context.Abort()
			}
			//context.Request.Header.Set(common.COMMON_AUTH_CURRENT, claims.UserId)
			context.Set(common.COMMON_AUTH_CURRENT, claims)
			context.Next()

		}
	}
}
