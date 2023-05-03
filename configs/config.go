package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
	"strings"
)

// HttpConfig 有关HTTP服务器的配置。
type HttpConfig struct {
	// Gin绑定的地址。
	BindAddress string `mapstructure:"bind_address"`
}

// MySQLConfig MySQL配置。
type MySQLConfig struct {
	// MySQL数据库地址。
	Address string `mapstructure:"address"`
	// MySQL数据库用户名。
	Username string `mapstructure:"username"`
	// MySQL数据库密码。
	Password string `mapstructure:"password"`
	// MySQL数据库名。
	Database string `mapstructure:"database"`
	// MySQL的DSN。
	DSN string `mapstructure:"dsn"`
}

// RedisConfig Redis配置。
type RedisConfig struct {
	// Redis地址。
	Address string `mapstructure:"address"`
	// Redis密码。
	Password string `mapstructure:"password"`
	// Redis数据库。
	Database int `mapstructure:"database"`
}

// AliyunOSSConfig 阿里云OSS配置。
type AliyunOSSConfig struct {
	// 阿里云OSS的Endpoint。
	Endpoint string `mapstructure:"endpoint"`
	// 阿里云OSS的AccessKeyID。
	AccessKeyID string `mapstructure:"access_key_id"`
	// 阿里云OSS的AccessKeySecret。
	AccessKeySecret string `mapstructure:"access_key_secret"`
}

// OidcConfig OIDC配置。
type OidcConfig struct {
	// Oidc的Issuer。
	Issuer string `mapstructure:"issuer"`
	// Oidc的ClientID。
	ClientID string `mapstructure:"client_id"`
	// Oidc的ClientSecret。
	ClientSecret string `mapstructure:"client_secret"`
	// Oidc的RedirectURL。
	RedirectURL string `mapstructure:"redirect_url"`
}

// Config 应用的整体配置。
type Config struct {
	// HTTP服务器的配置。
	Http HttpConfig `mapstructure:"http"`
	// MySQL的配置。
	MySQL MySQLConfig `mapstructure:"mysql"`
	// Redis的配置。
	Redis RedisConfig `mapstructure:"redis"`
	// 阿里云OSS的配置。
	AliyunOSS AliyunOSSConfig `mapstructure:"aliyun_oss"`
	// OIDC的配置。
	Oidc OidcConfig `mapstructure:"oidc"`
}

var conf Config

// BindEnvs 结构体绑定环境变量
//
// https://github.com/spf13/viper/issues/188
//
// viper对于嵌套的结构体，无法自动绑定环境变量。本函数通过Reflect实现了对嵌套结构体的绑定。
func BindEnvs(target interface{}, parts ...string) error {
	ifv := reflect.ValueOf(target)
	ift := reflect.TypeOf(target)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			err := BindEnvs(v.Interface(), append(parts, tv)...)
			if err != nil {
				return fmt.Errorf("failed to bind env: %w", err)
			}
		default:
			err := viper.BindEnv(strings.Join(append(parts, tv), "."))
			if err != nil {
				return fmt.Errorf("failed to bind env: %w", err)
			}
		}
	}
	return nil
}

// loadConfig 从viper中读取配置。
func loadConfig() error {
	if err := viper.Unmarshal(&conf); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return nil
}

// Init 初始化配置。
//
// 本应用主要从两个渠道读取配置。
//   - 配置文件（configs/config.toml）
//   - 环境变量，环境变量命名规则遵循Viper。
func Init(configFilePath string) error {
	// 读取配置文件
	viper.SetConfigFile(configFilePath)
	// 读取环境变量
	viper.AutomaticEnv()
	// 对环境变量的处理。
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("ELAB")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}
	if err := BindEnvs(conf); err != nil {
		return fmt.Errorf("failed to bind envs: %w", err)
	}
	if err := loadConfig(); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	return nil
}

func GetConfig() *Config {
	return &conf
}
