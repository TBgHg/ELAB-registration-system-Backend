package configs

import (
	log "ELAB-registration-system-Backend/logger"
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type oauthConfig struct {
	Issuer string `mapstructure:"issuer"`
}

type httpConfig struct {
	Addr  string       `mapstructure:"addr"`
	OAuth *oauthConfig `mapstructure:"oauth"`
}

type mobileConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

type dbConfig struct {
	Addr     string `mapstructure:"addr"`
	User     string `mapstructure:"user"`
	Pwd      string `mapstructure:"pwd"`
	Database string `mapstructure:"database"`
	DSN      string `mapstructure:"dsn"`
}

type redisConfig struct {
	Addr string `mapstructure:"addr"`
	Pwd  string `mapstructure:"pwd"`
	DB   int    `mapstructure:"db"`
}

type oidcConfig struct {
	Issuer       string `mapstructure:"issuer"`
	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
	RedirectUrl  string `mapstructure:"redirect_url"`
}

type ossConfig struct {
	Endpoint        string `toml:"endpoint"`
	AccessKeyID     string `toml:"access_key_id"`
	AccessKeySecret string `toml:"access_key_secret"`
	OssPreURL       string `toml:"oss_pre_url"`
}

type Config struct {
	// 服务端口
	Http      httpConfig   `mapstructure:"http"`
	DB        dbConfig     `mapstructure:"db"`
	Redis     redisConfig  `mapstructure:"redis"`
	Oidc      oidcConfig   `mapstructure:"oidc"`
	Mobile    mobileConfig `mapstructure:"mobile"`
	Cool      string       `mapstructure:"cool"`
  OssConfig ossConfig    `mapstructure:"oss"`
}

var conf Config

// https://github.com/spf13/viper/issues/188
// viper对于嵌套的结构体，无法自动绑定环境变量
func BindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			BindEnvs(v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}

func Init() {
	viper.SetConfigFile("configs/config.toml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("ELAB")
	if err := viper.ReadInConfig(); err != nil {
		log.Logger.Error("failed to read config file: ")
		log.Logger.Error(err.Error())
		return
	}
	BindEnvs(conf)
	// 试着输出conf的一些东西
	if err := loadConfig(); err != nil {
		log.Logger.Error("failed to load config: ")
		log.Logger.Error(err.Error())
		return
	}
}

func GetConfig() (*Config, error) {
	return &conf, nil
}

func loadConfig() error {
	if err := viper.Unmarshal(&conf); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return nil
}
