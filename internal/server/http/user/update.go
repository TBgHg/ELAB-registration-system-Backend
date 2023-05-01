package user

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/user"
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func updateUser(ctx *gin.Context) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		slog.Error("oidc.GetToken failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	updateBody := user.User{}
	err = ctx.ShouldBind(&updateBody)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	if updateBody.OpenId != token.Subject {
		ctx.JSON(403, model.NewUpdateBodyOpenIdMismatchError())
		return
	}
	err = user.UpdateUser(ctx, token.Subject, updateBody)
	if err != nil {
		slog.Error("user.UpdateUser failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
}
