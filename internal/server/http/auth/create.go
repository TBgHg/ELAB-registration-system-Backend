package auth

import (
	"crypto/sha256"
	"elab-backend/internal/model"
	"elab-backend/internal/model/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

// createLoginSession 创建一个新的登录会话
//
// 该接口通过输入state和code_verifier来创建一个新的登录会话，
// code_verifier存储在Redis库中，Key是 login://<state>，Value是code_verifier
func createLoginSession(ctx *gin.Context) {
	params := auth.NewSessionRequestParams{}
	err := ctx.ShouldBind(&params)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	// 使用SHA256对codeChallenge进行编码
	hash := sha256.New()
	hash.Write([]byte(params.CodeChallenge))
	codeChallenge := hash.Sum(nil)
	// 构造Redis Key
	key := auth.SessionUriScheme + params.State
	session := auth.Session{
		RedirectUri:  params.RedirectUri,
		CodeVerifier: string(codeChallenge),
	}
	// 将codeVerifier存储到Redis中
	err = auth.SetAuthSession(ctx, key, &session)
	if err != nil {
		err = fmt.Errorf("svc.Redis.SetAuthSession failed: %w", err)
		slog.Error(err.Error())
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, auth.NewSessionResponse{Ok: true})
}
