package redis

import (
	"elab-backend/configs"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	Client *redis.Client
}

func NewService(conf *configs.Config) *Service {
	return &Service{
		Client: redis.NewClient(&redis.Options{
			Addr:     conf.Redis.Address,
			Password: conf.Redis.Password,
			DB:       conf.Redis.Database,
		}),
	}
}
