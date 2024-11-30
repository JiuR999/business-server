package Websocket

import (
	"BusinessServer/common"
	publisher "BusinessServer/common/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var ws = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
	return true
}}

func HandFunc(ctx *gin.Context) {
	conn, err := ws.Upgrade(ctx.Writer, ctx.Request, nil)
	closeFlag := false
	if err != nil {

	}
	conn.SetCloseHandler(func(code int, text string) error {
		fmt.Println("关闭连接")
		closeFlag = true
		return nil
	})

	eventChan := publisher.EB.Subscribe(common.COMMON_MSG)
	go func() {
		for {
			select {
			case event := <-eventChan:
				bytes, _ := json.Marshal(event)
				conn.WriteMessage(websocket.TextMessage, bytes)
			}

			if closeFlag {
				break
			}

			time.Sleep(1 * time.Millisecond)
		}
	}()
}

func HandNoticeFunc(ctx *gin.Context) {
	conn, err := ws.Upgrade(ctx.Writer, ctx.Request, nil)
	closeFlag := false
	if err != nil {

	}
	conn.SetCloseHandler(func(code int, text string) error {
		fmt.Println("关闭连接")
		closeFlag = true
		return nil
	})

	eventChan := publisher.EB.Subscribe(common.EVENT_NOTICE)
	go func() {
		for {
			select {
			case event := <-eventChan:
				log.Print("收到通知", event.Data)
				bytes, _ := json.Marshal(event)
				conn.WriteMessage(websocket.TextMessage, bytes)
			}

			if closeFlag {
				break
			}

			time.Sleep(1 * time.Millisecond)
		}
	}()
}
