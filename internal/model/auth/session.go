package auth

import (
	"elab-backend/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"time"
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

type SessionNotFoundError struct{}

func (e SessionNotFoundError) Error() string {
	return "session_not_found"
}

func GetAuthSession(ctx *gin.Context, state string) (*Session, error) {
	svc := service.GetService()
	key := SessionUriScheme + state
	exists, err := svc.Redis.Client.Exists(ctx, key).Result()
	if err != nil {
		err = fmt.Errorf("redis.Exists failed: %w", err)
		slog.Error(err.Error())
		return nil, err
	}
	if exists == 0 {
		return nil, SessionNotFoundError{}
	}
	marshalledSession, err := svc.Redis.Client.Get(ctx, key).Result()
	if err != nil {
		err = fmt.Errorf("redis.Get failed: %w", err)
		slog.Error(err.Error())
		return nil, err
	}
	slog.Info(marshalledSession)
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

const SessionUriScheme = "auth-session://"
const SessionExpiration = time.Minute * 10

// NewSessionRequest 使用State首先存储RedirectUri和code_verifier
// 用于 POST /auth/new
type NewSessionRequest struct {
	State       string `json:"state" binding:"required"`
	RedirectUri string `json:"redirect_uri" binding:"required"`
	Verifier    string `json:"code_verifier" binding:"required"`
}

type Session struct {
	RedirectUri  string `json:"redirect_uri"`
	CodeVerifier string `json:"code_verifier"`
}

// NewSessionResponse 用于 POST /auth/new 的响应格式。
type NewSessionResponse struct {
	Ok bool `json:"ok"`
}
