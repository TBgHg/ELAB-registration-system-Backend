package comment

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/content/thread/comment"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func NewCommentNotFoundError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "comment_not_found",
		Message: "找不到要删除的评论",
	}
}

func updateCommentById(ctx *gin.Context) {
	request := comment.UpdateRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	exists, err := comment.CheckExistsById(ctx, &comment.CheckExistsByIdOptions{
		SpaceId:   request.SpaceId,
		CommentId: request.CommentId,
	})
	if err != nil {
		slog.Error("error in comment.CheckExistsById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if !exists {
		ctx.JSON(404, NewCommentNotFoundError())
		return
	}
	err = comment.Update(ctx, &request)
	if err != nil {
		slog.Error("error in comment.Update, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{Ok: true})
}
