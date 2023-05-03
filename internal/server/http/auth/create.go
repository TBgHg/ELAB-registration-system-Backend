package auth

import (
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
	params := auth.NewSessionRequest{}
	err := ctx.ShouldBind(&params)
	if err != nil {
		slog.Error("ctx.ShouldBind failed", "error", err)
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	// 构造Redis Key
	session := auth.Session{
		RedirectUri:  params.RedirectUri,
		CodeVerifier: params.Verifier,
	}
	// 将codeVerifier存储到Redis中
	err = auth.SetAuthSession(ctx, params.State, &session)
	if err != nil {
		err = fmt.Errorf("svc.Redis.SetAuthSession failed: %w", err)
		slog.Error(err.Error())
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, auth.NewSessionResponse{Ok: true})
}
