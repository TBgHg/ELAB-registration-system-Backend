package redis

import (
	"elab-backend/configs"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	Client *redis.Client
}

func NewService() *Service {
	config := configs.GetConfig()
	return &Service{
		Client: redis.NewClient(&redis.Options{
			Addr:     config.Redis.Address,
			Password: config.Redis.Password,
			DB:       config.Redis.Database,
		}),
	}
}
