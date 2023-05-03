package main

import (
	"elab-backend/configs"
	"elab-backend/internal/model"
	"elab-backend/internal/server/http/server"
	"elab-backend/internal/service"
	"flag"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"os"
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
	gin.SetMode(gin.DebugMode)
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr)))
	err = model.Init()
	if err != nil {
		panic(err)
	}
	err = engine.Run(configs.GetConfig().Http.BindAddress)
	if err != nil {
		panic(err)
	}
}
