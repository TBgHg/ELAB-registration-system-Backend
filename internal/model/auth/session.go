package auth

import (
	"elab-backend/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func SetAuthSession(ctx *gin.Context, state string, session *Session) error {
	svc := service.GetService()
	key := SessionUriScheme + state
	// 使用JSON序列化
	marshalledSession, err := json.Marshal(session)
	if err != nil {
		err = fmt.Errorf("json.Marshal failed: %w", err)
		slog.Error(err.Error())
		return err
	}
	err = svc.Redis.Client.Set(ctx, key, string(marshalledSession), SessionExpiration).Err()
	if err != nil {
		err = fmt.Errorf("redis.Set failed: %w", err)
		slog.Error(err.Error())
		return err
	}
	return nil
}

func GetAuthSession(ctx *gin.Context, state string) (*Session, error) {
	svc := service.GetService()
	key := SessionUriScheme + state
	marshalledSession, err := svc.Redis.Client.Get(ctx, key).Result()
	if err != nil {
		err = fmt.Errorf("redis.Get failed: %w", err)
		slog.Error(err.Error())
		return nil, err
	}
	session := Session{}
	// 使用JSON解析
	err = json.Unmarshal([]byte(marshalledSession), &session)
	if err != nil {
		err = fmt.Errorf("json.Unmarshal failed: %w", err)
		slog.Error(err.Error())
		return nil, err
	}
	err = svc.Redis.Client.Del(ctx, key).Err()
	if err != nil {
		err = fmt.Errorf("redis.Del failed: %w", err)
		slog.Error(err.Error())
		return nil, err
	}
	return &session, nil
}
