package main

import (
	"ELAB-registration-system-Backend/configs"
	"ELAB-registration-system-Backend/internal/server/http"
	"ELAB-registration-system-Backend/internal/service"
	log "ELAB-registration-system-Backend/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	// 调用Init函数，读取配置文件
	configs.Init()
	conf, err := configs.GetConfig()
	if err != nil {
		log.Logger.Error("main configs.GetConfig failed err:" + err.Error())
		return
	}

	// 创建一个默认的路由引擎
	svc, err := service.NewService(conf)
	if err != nil {
		log.Logger.Error("main service.NewService failed err:" + err.Error())
		return
	}

	r := gin.Default()
	http.Init(r, svc)
	// 启动服务
	err = r.Run(conf.Http.Addr)
	if err != nil {
		log.Logger.Error("main r.Run failed err:" + err.Error())
		return
	}
}
