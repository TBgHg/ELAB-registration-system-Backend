package service

import (
	log "ELAB-registration-system-Backend/common/logger"
	"ELAB-registration-system-Backend/configs"
	"ELAB-registration-system-Backend/internal/dao"
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

var svc *Service

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

	return &Service{
		Conf: conf,
		db:   dao.Use(gormDB),
		Rdb:  rdb,
	}, nil
}
