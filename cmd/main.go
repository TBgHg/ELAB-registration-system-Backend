package main

import (
	"elab-backend/configs"
	"elab-backend/internal/server/http/server"
	"elab-backend/internal/service"
	"flag"
	"github.com/gin-gonic/gin"
)

var configDirectory string

// FlagInit 初始化命令行参数
func flagInit() {
	flag.StringVar(&configDirectory, "config", "configs/config.toml", "config directory")
	flag.Parse()
}

func main() {
	flagInit()
	err := configs.Init(configDirectory)
	if err != nil {
		panic(err)
	}
	err = service.Init(configs.GetConfig())
	if err != nil {
		panic(err)
	}
	engine := gin.Default()
	server.Init(engine)
	err = engine.Run(configs.GetConfig().Http.BindAddress)
	if err != nil {
		panic(err)
	}
}
