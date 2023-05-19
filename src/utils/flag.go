package utils

import (
	"flag"
	"os"
)

var (
	//Env 环境名称
	Env string
	//ConfPath 配置文件所在位置
	ConfPath string
)

// FlagsInit FlagsInit
func FlagsInit() {
	flag.StringVar(&Env, "env", "", "running env")
	flag.StringVar(&ConfPath, "conf", "", "running env")
	flag.Parse()

	if Env == "" {
		Env = os.Getenv("env")
	}

	if ConfPath == "" {
		ConfPath = os.Getenv("conf")
	}
}
