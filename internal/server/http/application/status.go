package application

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/application"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"golang.org/x/sync/errgroup"
)

func getStatus(ctx *gin.Context) {
	longTextFormExists := make(chan bool)
	roomSelected := make(chan bool)
	g := errgroup.Group{}
	g.Go(func() error {
		exists, err := application.CheckLongTextFormExists(ctx)
		if err != nil {
			return err
		}
		longTextFormExists <- exists
		return nil
	})
	g.Go(func() error {
		selected, err := application.CheckRoomSelected(ctx)
		if err != nil {
			return err
		}
		roomSelected <- selected
		return nil
	})
	err := g.Wait()
	if err != nil {
		slog.Error("MySQL.Model failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, application.StatusResponse{
		LongTextForm:  <-longTextFormExists,
		RoomSelection: <-roomSelected,
	})
}
