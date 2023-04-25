package configs

import (
	log "ELAB-registration-system-Backend/logger"
	"os"

	"github.com/BurntSushi/toml"
)

type oauthConfig struct {
	Issuer string `toml:"issuer"`
}

type httpConfig struct {
	Addr  string       `toml:"addr"`
	OAuth *oauthConfig `toml:"oauth"`
}

type mobileConfig struct {
	Endpoint string `toml:"endpoint"`
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

type oidcConfig struct {
	Issuer       string `toml:"issuer"`
	ClientId     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
	RedirectUrl  string `toml:"redirect_url"`
}

type ossConfig struct {
	Endpoint        string `toml:"endpoint"`
	AccessKeyID     string `toml:"access_key_id"`
	AccessKeySecret string `toml:"access_key_secret"`
	OssPreURL       string `toml:"oss_pre_url"`
}

type Config struct {
	// 服务端口
	Http      *httpConfig   `toml:"http"`
	DB        *dbConfig     `toml:"db"`
	Redis     *redisConfig  `toml:"redis"`
	Oidc      *oidcConfig   `toml:"oidc"`
	Mobile    *mobileConfig `toml:"mobile"`
	OssConfig *ossConfig    `toml:"oss"`
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
}
