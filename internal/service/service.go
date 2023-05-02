package service

import (
	"elab-backend/configs"
	"elab-backend/internal/service/mysql"
	"elab-backend/internal/service/oidc"
	"elab-backend/internal/service/redis"
	"gorm.io/gorm"
)

var svc *Service

type Service struct {
	Oidc  *oidc.Service
	Redis *redis.Service
	MySQL *gorm.DB
}

func Init(conf *configs.Config) error {
	oidcService, err := oidc.NewService(conf)
	if err != nil {
		return err
	}
	redisService := redis.NewService(conf)
	if err != nil {
		return err
	}
	mysqlService, err := mysql.NewService(conf)
	if err != nil {
		return err
	}
	svc = &Service{
		Oidc:  oidcService,
		Redis: redisService,
		MySQL: mysqlService,
	}
	return nil
}

func GetService() *Service {
	return svc
}
