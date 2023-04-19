package configs

import (
	log "ELAB-registration-system-Backend/logger"
	"github.com/BurntSushi/toml"
	"os"
)

type oauthConfig struct {
	Issuer string `toml:"issuer"`
}

type httpConfig struct {
	Addr  string       `toml:"addr"`
	OAuth *oauthConfig `toml:"oauth"`
}

type dbConfig struct {
	Addr     string `toml:"addr"`
	User     string `toml:"user"`
	Pwd      string `toml:"pwd"`
	Database string `toml:"database"`
	DSN      string `toml:"dsn"`
}

type redisConfig struct {
	Addr string `toml:"addr"`
	Pwd  string `toml:"pwd"`
	DB   int    `toml:"db"`
}

type Config struct {
	// 服务端口
	Http  *httpConfig  `toml:"http"`
	DB    *dbConfig    `toml:"db"`
	Redis *redisConfig `toml:"redis"`
}

var (
	conf Config
	err  error
)

func Init() {
	loadConfig()
}

func GetConfig() (*Config, error) {
	return &conf, err
}

func loadConfig() {
	var data []byte
	// 读取配置文件
	data, err = os.ReadFile("./configs/config.toml")
	if err != nil {
		log.Logger.Error("loadConfig os.ReadFile failed err:" + err.Error())
		return
	}
	err = toml.Unmarshal(data, &conf)
	if err != nil {
		log.Logger.Error("loadConfig toml.Unmarshal failed err:" + err.Error())
		return
	}
	return
}
