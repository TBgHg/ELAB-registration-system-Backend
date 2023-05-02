package comment

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/content/thread/comment"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func NewCommentAlreadyLikedError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "comment_already_liked",
		Message: "已经点赞过该评论",
	}
}

func likeCommentById(ctx *gin.Context) {
	request := comment.LikeRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	exists, err := comment.CheckLike(ctx, &request)
	if err != nil {
		slog.Error("error in comment.CheckExistsById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if !exists {
		ctx.JSON(400, NewCommentAlreadyLikedError())
		return
	}
	err = comment.Like(ctx, &request)
	if err != nil {
		slog.Error("error in comment.Like, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{Ok: true})
}

func dislikeCommentById(ctx *gin.Context) {
	request := comment.LikeRequest{}
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	exists, err := comment.CheckLike(ctx, &request)
	if err != nil {
		slog.Error("error in comment.CheckExistsById, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if exists {
		ctx.JSON(400, NewCommentAlreadyLikedError())
		return
	}
	err = comment.Dislike(ctx, &request)
	if err != nil {
		slog.Error("error in comment.Dislike, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, space.OperationResponse{Ok: true})
}
