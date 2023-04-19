package service

import (
	"ELAB-registration-system-Backend/configs"
	"ELAB-registration-system-Backend/internal/dao"
	log "ELAB-registration-system-Backend/logger"
	"context"
	"fmt"

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

	//err = OAuth2Init(conf.Http.OAuth.Issuer)
	if err != nil {
		log.Logger.Error("NewService OAuth2Init failed err:" + err.Error())
		return nil, err
	}

	return &Service{
		Conf:         conf,
		db:           dao.Use(gormDB),
		Rdb:          rdb,
		Provider:     provider,
		OAuth2Config: &oauth2Config,
	}, nil
}
