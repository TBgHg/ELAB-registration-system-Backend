package service

import (
	"ELAB-registration-system-Backend/configs"
	"ELAB-registration-system-Backend/internal/dao"
	log "ELAB-registration-system-Backend/logger"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Service struct {
	Conf *configs.Config
	db   *dao.Query
	Rdb  *redis.Client
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

	err = OAuth2Init(conf.Http.OAuth.Issuer)
	if err != nil {
		log.Logger.Error("NewService OAuth2Init failed err:" + err.Error())
		return nil, err
	}

	return &Service{
		Conf: conf,
		db:   dao.Use(gormDB),
		Rdb:  rdb,
	}, nil
}
