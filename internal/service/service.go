package service

import (
	"elab-backend/configs"
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
	svc = &Service{
		Oidc: nil,
	}
	return nil
}

func GetService() *Service {
	return svc
}
