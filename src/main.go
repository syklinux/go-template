package main

import (
	"GoTemplate/src/dao"
	"GoTemplate/src/router"
	"GoTemplate/src/utils"
	"os"
	"os/signal"

	"github.com/syklinux/golib/log"
)

func main() {
	var err error

	// 获取环境
	utils.FlagsInit()

	// 初始化配置
	utils.LoadConf(utils.Env, utils.ConfPath)

	// 初始化log
	if err = utils.InitLog(); err != nil {
		log.Fatalf("InitLog Error: %s", err)
	}

	// 初始化mysql
	mysql := dao.NewDao(true)
	defer mysql.Close()

	// 初始化redis
	dao.InitRedis()
	defer dao.RedisClose()

	// 加载路由
	router.HTTPServerRun()

	// 捕捉信号
	quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(quit, os.Interrupt)
	<-quit
	router.HTTPServerStop()
	mysql.Close()
	dao.RedisClose()
}
