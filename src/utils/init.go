package utils

import (
	"GoTemplate/src/config"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/syklinux/golib/log"
)

var (
	//Conf 配置加载
	Conf config.Conf
)

// LoadConf 读取配置文件
func LoadConf(env string, path string) {

	if env == "" {
		panic("no env")
	}
	if path == "" {
		panic("no conf path")
	}

	path = path + "/conf_" + env + ".json"

	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	var conf config.Conf

	err = json.Unmarshal(byteValue, &conf)
	if err != nil {
		panic("load conf err" + err.Error())
	}

	Conf = conf
}

// InitLog 初始化日志通道
func InitLog() (err error) {
	err = log.InitByConf(Conf.Log)
	if err != nil {
		return fmt.Errorf("InitLog failed: %s", err.Error())
	}
	return
}
