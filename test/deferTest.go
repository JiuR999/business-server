package main

import (
	"fmt"
	"sync"
)

func main() {
	group := sync.WaitGroup{}
	group.Add(1)
	defer func() {
		fmt.Println("正常退出")
		//group.Done()
	}()
	go func() {
		//恢复同一协程
		defer func() {
			if e := recover(); e != nil {
				fmt.Println(e)
				group.Done()
			}
		}()
		panic("抛出异常")
	}()
	group.Wait()
}
