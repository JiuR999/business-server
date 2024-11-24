package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func main() {
	var path string
	//"D:\\Downloads\\data_2"
	fmt.Println("请输入游戏文件地址,例如：D:\\Star Rail\\Game\\StarRail_Data\\webCaches\\Cache\\Cache_Data\\data_2")
	fmt.Scan(&path)
	file, err := os.Open(path)
	if err == nil {
		defer file.Close()
		bytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		s := string(bytes)
		compile, err := regexp.Compile("https.*/getGachaLog.*")
		if err != nil {
			fmt.Println(err.Error())
		}
		allString := compile.FindAllString(s, -1)
		fmt.Println("共匹配到", len(allString))
		gachaUrl := allString[len(allString)-1]
		gachaUrls := strings.Split(gachaUrl, "/0/")

		fmt.Println("~~~~~ 抽卡链接 ~~~~~~~")
		const prefix = "https://api-takumi.mihoyo.com/common/gacha_record"
		path := gachaUrls[len(gachaUrls)-1]
		r2, _ := regexp.Compile("/api.*")
		findString := r2.FindString(path)
		fmt.Println(prefix + findString)
	}
}
