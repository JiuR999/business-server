package middleware

import (
	"BusinessServer/common/abstract/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"
	"strings"
)

func Recovery() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				printStackTrace()
				common.NewResponse(context).ErrorWithMsg("内部错误!")
			}
		}()
		context.Next()
	}
}

func printStackTrace() {
	stackBuf := make([]byte, 4096)
	stackSize := runtime.Stack(stackBuf, false)
	stackTrace := strings.Split(string(stackBuf[:stackSize]), "\n")

	fmt.Println("Stack Trace:")
	for _, line := range stackTrace {
		if strings.HasPrefix(line, "\t") {
			fmt.Println(line)
		}
	}
}
