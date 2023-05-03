package member

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space/member"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func getMemberList(ctx *gin.Context) {
	spaceId := ctx.Param("space_id")
	members, err := member.GetSpaceMemberList(ctx, spaceId)
	if err != nil {
		slog.Error("error in space.GetSpaceMemberList, %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, member.ListResponse{
		Members: *members,
	})
}
