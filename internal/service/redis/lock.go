package redis

import (
	"context"
	"time"
)

const Timeout = time.Second * 10
const RetryTimes = 100
const RetryInterval = time.Millisecond * 100

type GetLockTimeoutError struct{}

func (err *GetLockTimeoutError) Error() string {
	return "get lock timeout"
}

// GetLock 使用Redis获取分布式锁
func (s *Service) GetLock(ctx context.Context, key string) (func(), error) {
	// 重试次数
	for i := 0; i < RetryTimes; i++ {
		ok, err := s.Client.SetNX(ctx, key, "1", Timeout).Result()
		if err != nil {
			return nil, err
		}
		if ok {
			return func() {
				s.Client.Del(ctx, key)
			}, nil
		}
	}
	return nil, &GetLockTimeoutError{}
}
