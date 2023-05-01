package main

import (
	"flag"
)

var configDirectory string

// FlagInit 初始化命令行参数
func flagInit() {
	flag.StringVar(&configDirectory, "config", "configs/config.toml", "config directory")
	flag.Parse()
}

func main() {
	flagInit()
	//// 调用Init函数，读取配置文件
	//err := configs.Init(configDirectory)
	//if err != nil {
	//	slog.Error("无法读取配置：%w", err)
	//	return
	//}
	//conf := configs.GetConfig()
	//if conf == nil {
	//	slog.Error("config.conf似乎是一个空指针，请确认是否已经调用了Init函数")
	//	return
	//}
	//
	//// 创建一个默认的路由引擎
	//svc, err := service.NewService(conf)
	//if err != nil {
	//	slog.Error("无法创建服务：%w", err)
	//	return
	//}
	//// 创建一个Gin实例
	//r := gin.Default()
	//http.Init(r, svc)
	//// 启动服务
	//err = r.Run(conf.Http.BindAddress)
	//if err != nil {
	//	slog.Error("无法启动服务：%w", err)
	//	return
	//}
}
