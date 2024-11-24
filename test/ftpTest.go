package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"sync"
)

func main() {
	client, err := ssh.Dial("tcp", "103.243.26.204:22", &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("6fc15103c4b2")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ftp, err := sftp.NewClient(client)
	if err != nil {
		fmt.Println("连接失败")
	}
	srcDir := "/home/swust"
	dstDir := "D:\\ftp\\"

	file, err := ftp.Open(srcDir + "/" + "docker-compose.yaml")
	if err != nil {
		fmt.Println("Open Err!")
	}
	os.Mkdir(dstDir, 0644)
	// 创建本地文件的等待拷贝
	logFile, _ := os.Create(dstDir + "\\" + "1.yaml")
	// 创建资源记得关闭
	defer logFile.Close()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		io.Copy(logFile, file)
		wg.Done()
	}()
	wg.Wait()
	//e := session.Run("cat " + srcDir + "/server/" + "BS")

	//_, err = io.Copy(logFile, file)
	//if err != nil {
	//	fmt.Println("Writer error!")
	//	return
	//}

}
