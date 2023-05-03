package thread

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/space"
	"elab-backend/internal/model/space/content/thread"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func deleteThreadById(c *gin.Context) {
	threadId := c.GetString("thread_id")
	spaceId := c.GetString("space_id")
	exists, err := thread.CheckExistsById(c, spaceId, threadId)
	if err != nil {
		slog.Error("error in thread.CheckExistsById, %w", err)
		c.JSON(500, model.NewInternalServerError())
		return
	}
	if !exists {
		c.JSON(404, NewNotFoundError())
		return
	}
	err = thread.HardDelete(c, spaceId, threadId)
	if err != nil {
		slog.Error("error in thread.HardDelete, %w", err)
		c.JSON(500, model.NewInternalServerError())
		return
	}
	c.JSON(200, space.OperationResponse{
		Ok: true,
	})
}
