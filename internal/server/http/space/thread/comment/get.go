package comment

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space/content/thread/comment"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func getComment(ctx *gin.Context) {
	request := comment.GetRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	response, err := comment.Get(ctx, &request)
	if err != nil {
		slog.Error("error in comment.Get, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, response)
}
