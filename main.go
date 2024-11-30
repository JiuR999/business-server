package main

import (
	orderService "BusinessServer/apps/order/service"
	"BusinessServer/apps/system/log/service"
	"BusinessServer/router"
)

func main() {
	go func() {
		service.StartLogService()
	}()
	go func() {
		orderService.StartAsyncService()
	}()
	router.Init()
}
