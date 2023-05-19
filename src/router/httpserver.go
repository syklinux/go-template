package router

import (
	"GoTemplate/src/middleware"
	"GoTemplate/src/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	//HTTPSrvHandler HTTPSrvHandler
	HTTPSrvHandler *http.Server
)

// HTTPServerRun HTTPServerRun
func HTTPServerRun() {
	var httpConf = utils.Conf.HTTPConf
	r := InitRouter(middleware.Cors())
	port := ":" + utils.Conf.Port
	HTTPReadTime := httpConf.ReadTimeout
	HTTPWriteTime := httpConf.WriteTimeout
	HTTPMaxBytes := httpConf.MaxHeaderBytes
	HTTPSrvHandler = &http.Server{
		Addr:           port,
		Handler:        r,
		ReadTimeout:    time.Duration(HTTPReadTime) * time.Second,
		WriteTimeout:   time.Duration(HTTPWriteTime) * time.Second,
		MaxHeaderBytes: 1 << uint(HTTPMaxBytes),
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", port)
		if err := HTTPSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", port, err)
		}
	}()
}

// HTTPServerStop HTTPServerStop
func HTTPServerStop() {
	fmt.Println("5s后关闭http")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := HTTPSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
