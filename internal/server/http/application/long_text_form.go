package application

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/application"
	"elab-backend/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

func NewLongTextFormNotFoundError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "long_text_form_not_found",
		Message: "无法找到长文本表单",
	}
}

func getLongTextForm(ctx *gin.Context) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		slog.Error("oidc.GetToken failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	longTextForm := application.LongTextForm{}
	err = svc.MySQL.Model(&application.LongTextForm{
		OpenId: token.Subject,
	}).First(&longTextForm).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(404, NewLongTextFormNotFoundError())
			return
		}
		slog.Error("MySQL.Model failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, longTextForm)
}

func updateLongTextForm(ctx *gin.Context) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		slog.Error("oidc.GetToken failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	updateBody := application.LongTextForm{}
	err = ctx.ShouldBind(&updateBody)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	if updateBody.OpenId != token.Subject {
		ctx.JSON(403, model.NewUpdateBodyOpenIdMismatchError())
		return
	}
	err = application.UpdateLongTextForm(ctx, &updateBody)
	if err != nil {
		slog.Error("application.UpdateLongTextForm failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, application.UpdateLongTextFormResponse{
		Ok: true,
	})
}
