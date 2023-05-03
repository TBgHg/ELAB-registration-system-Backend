package thread

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space/content/thread"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func NewNotFoundError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "thread_not_found",
		Message: "未找到该 thread",
	}
}

func getThreadHistoryById(ctx *gin.Context) {
	request := thread.GetContentByHistoryIdOptions{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	exists, err := thread.CheckHistoryExistsById(ctx, request.SpaceId, request.ThreadId, request.HistoryId)
	if err != nil {
		slog.Error("error in thread.CheckHistoryExistsById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if !exists {
		ctx.JSON(400, NewNotFoundError())
		return
	}
	response, err := thread.GetContentByHistoryId(ctx, &request)
	if err != nil {
		slog.Error("error in thread.CheckHistoryExistsById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, thread.ContentResponse{
		Content: response,
	})
}
