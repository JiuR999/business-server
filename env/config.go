package env

import (
	"gopkg.in/yaml.v3"
	"os"
	"runtime"
)

const (
	COMMON_PROJECT_NAME = "BusinessServer"
)

type config struct {
	ClusterID  int64  `yaml:"clusterId"`
	Port       string `yaml:"port"`
	EnableAuth bool   `yaml:"auth"`
	Database
	Jwt
	Default
}

var serverConfig = new(config)

func GetConfig() *config {
	return serverConfig
}

type Jwt struct {
	ExpireTime int    `yaml:"expireTime"` //有效时间
	SignUser   string `yaml:"signUser"`   //签发人
	Secrect    string `yaml:"secrect"`    //密钥
}

type Database struct {
	UserName     string `yaml:"username"` //数据库用户名
	Password     string `yaml:"password"` //密码
	Address      string `yaml:"address"`  //服务器地址
	DataBasePort string `yaml:"dataBasePort"`
	Name         string `yaml:"name"`
}

type Default struct {
	PageSize int `yaml:"pageSize"`
}

func init() {
	err := loadConfig(serverConfig)
	if err != nil {
		return
	}
}

func loadConfig(conf *config) error {
	goos := runtime.GOOS
	filePrefix := "./"
	if goos == "linux" {
		filePrefix = "/etc/business/"
	}
	yamlFile, err := os.ReadFile(filePrefix + "config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return err
	}
	return nil
}
