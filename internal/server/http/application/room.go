package application

import (
	"elab-backend/internal/model"
	"elab-backend/internal/model/application"
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func NewRoomAlreadySelectedError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "room_already_selected",
		Message: "已经选择过面试场地",
	}
}

func NewRoomUnselectedError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "room_unselected",
		Message: "未选择面试场地",
	}
}

func NewRoomCapacityExceededError() model.ErrorResponse {
	return model.ErrorResponse{
		Error:   "room_capacity_exceeded",
		Message: "面试场地容量已满",
	}
}

func getRoom(ctx *gin.Context) {
	rooms, err := application.GetRoom(ctx)
	if err != nil {
		slog.Error("MySQL.Model failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, application.GetInterviewRoomResponse{
		Rooms: *rooms,
	})
}

func setRoom(ctx *gin.Context) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		slog.Error("oidc.GetToken failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	roomSelection := application.InterviewRoomSelection{
		OpenId: token.Subject,
	}
	// 首先检测是否已经存在面试场地选择
	alreadySelected, err := application.CheckRoomSelected(ctx)
	if err != nil {
		slog.Error("MySQL.Model failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if alreadySelected {
		ctx.JSON(403, NewRoomAlreadySelectedError())
		return
	}
	err = ctx.ShouldBind(&roomSelection)
	if err != nil {
		ctx.JSON(400, model.NewInvalidParamError())
		return
	}
	lockKey := application.InterviewRoomUriScheme + roomSelection.RoomId
	unLock, err := svc.Redis.GetLock(ctx, lockKey)
	if err != nil {
		slog.Error("redis.GetLock failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	defer unLock()
	// 检测面试场地是否满足容量条件
	room, err := application.GetRoomById(ctx, roomSelection.RoomId)
	if err != nil {
		slog.Error("MySQL.Model failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if room.Capacity-room.CurrentOccupancy <= 0 {
		ctx.JSON(403, NewRoomCapacityExceededError())
		return
	}
	// 更新面试场地容量，分两步：先新增一条InterviewRoomSelection记录，再更新InterviewRoom记录
	err = application.SetRoom(ctx, &roomSelection)
	if err != nil {
		slog.Error("MySQL.Model failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, application.InterviewRoomOperateResponse{
		Ok: true,
	})
}

func deleteRoom(ctx *gin.Context) {
	// 1. 检测是否实际存在作为选择
	// 2. 加锁
	// 3. 删除面试场地选择
	svc := service.GetService()
	selected, err := application.CheckRoomSelected(ctx)
	if err != nil {
		slog.Error("MySQL.Model failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	if !selected {
		ctx.JSON(403, NewRoomUnselectedError())
		return
	}
	roomSelection, err := application.GetRoomSelection(ctx)
	if err != nil {
		slog.Error("MySQL.Model failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	lockKey := application.InterviewRoomUriScheme + roomSelection.RoomId
	unLock, err := svc.Redis.GetLock(ctx, lockKey)
	if err != nil {
		slog.Error("redis.GetLock failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	defer unLock()
	err = application.DeleteRoom(ctx)
	if err != nil {
		slog.Error("MySQL.Model failed %w", err)
		ctx.JSON(500, model.NewInternalServerError())
		return
	}
	ctx.JSON(200, application.InterviewRoomOperateResponse{
		Ok: true,
	})
}
