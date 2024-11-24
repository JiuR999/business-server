package main

import (
	"BusinessServer/apps/system/log/service"
	"BusinessServer/router"
)

func main() {
	go func() {
		service.StartLogService()
	}()
	router.Init()
	/*tx := db.Orm.DB().
		First(&m)
	if err := tx.ErrorWithMsg; err != nil {
		fmt.Println(err.ErrorWithMsg())
	} else {
		fmt.Println(*m.Name)
	}*/
}
