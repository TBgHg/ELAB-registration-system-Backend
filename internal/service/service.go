package service

import (
	"ELAB-registration-system-Backend/configs"
	"ELAB-registration-system-Backend/internal/dao"
	log "ELAB-registration-system-Backend/logger"
	"context"
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/redis/go-redis/v9"
	"golang.org/x/oauth2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	Conf         *configs.Config
	db           *dao.Query
	Rdb          *redis.Client
	Provider     *oidc.Provider
	OAuth2Config *oauth2.Config
	OssBucket    *oss.Bucket
}

func NewService(conf *configs.Config) (*Service, error) {
	dsn := fmt.Sprintf(conf.DB.DSN, conf.DB.User, conf.DB.Pwd, conf.DB.Addr, conf.DB.Database)
	gormDB, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Logger.Error("NewService gorm.Open failed err:" + err.Error())
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Pwd,
		DB:       conf.Redis.DB,
	})

	provider, err := oidc.NewProvider(context.Background(), conf.Oidc.Issuer)
	if err != nil {
		log.Logger.Error("NewService oidc.NewProvider failed err:" + err.Error())
		return nil, err
	}
	oauth2Config := oauth2.Config{
		ClientID:     conf.Oidc.ClientId,
		ClientSecret: conf.Oidc.ClientSecret,
		RedirectURL:  conf.Oidc.RedirectUrl,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	if err != nil {
		log.Logger.Error("NewService OAuth2Init failed err:" + err.Error())
		return nil, err
	}

	bucket, err := ossInit(conf.OssConfig.Endpoint, conf.OssConfig.AccessKeyID, conf.OssConfig.AccessKeySecret)
	if err != nil {
		log.Logger.Error("NewService ossInit failed err:" + err.Error())
		return nil, err
	}

	return &Service{
		Conf:         conf,
		db:           dao.Use(gormDB),
		Rdb:          rdb,
		Provider:     provider,
		OAuth2Config: &oauth2Config,
		OssBucket:    bucket,
	}, nil
}

// ossInit 初始化，将ConnQuery与数据库绑定
func ossInit(endpoint, accessKeyID, accessKeySecret string) (bucket *oss.Bucket, err error) {
	// 连接OSS账户
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		log.Logger.Error("连接OSS账户失败" + err.Error())
		return nil, err
	} else { // OSS账户连接成功
		// 连接存储空间
		bucket, err = client.Bucket("byte-dance-01")
		if err != nil {
			log.Logger.Error("连接存储空间失败" + err.Error())
			return nil, err
		} else { // 存储空间连接成功
			log.Logger.Info("OSS初始化完成")
		}
	}
	return bucket, nil
}
