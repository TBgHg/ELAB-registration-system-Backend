package application

import (
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const InterviewRoomUriScheme = "interview-room://"

// InterviewRoom 用于面试的房间
//
// 在Gin当中并不会被用作Request，而是作为Response.
type InterviewRoom struct {
	gorm.Model `json:"-"` // 不会被序列化
	// Id 房间Id
	RoomId string `json:"id"`
	// Name 房间名称
	Name string `json:"name"`
	// Time 面试时间，应该是UNIX时间戳以秒为单位
	Time int64 `json:"time"`
	// Capacity 房间容量
	Capacity int32 `json:"capacity"`
	// CurrentOccupancy 房间已报名人数
	CurrentOccupancy int32 `json:"current_occupancy"`
	// Location 房间地点
	Location string `json:"location"`
}

type InterviewRoomSelection struct {
	gorm.Model `json:"-"` // 不会被序列化
	// OpenId 用户OpenId
	OpenId string `json:"openid" binding:"required,uuid" gorm:"column:openid"`
	// RoomId 房间Id
	RoomId string `json:"room_id" binding:"required,uuid" gorm:"column:room_id"`
}

type GetInterviewRoomResponse struct {
	Rooms []InterviewRoom `json:"rooms"`
}

type InterviewRoomOperateResponse struct {
	Ok bool `json:"ok"`
}

func GetRoom(ctx *gin.Context) (*[]InterviewRoom, error) {
	var rooms []InterviewRoom
	svc := service.GetService()
	err := svc.MySQL.WithContext(ctx).Model(&InterviewRoom{}).Find(&rooms).Error
	if err != nil {
		return nil, err
	}
	return &rooms, nil
}

func GetRoomById(ctx *gin.Context, roomId string) (*InterviewRoom, error) {
	room := InterviewRoom{
		RoomId: roomId,
	}
	svc := service.GetService()
	err := svc.MySQL.WithContext(ctx).First(&room).Error
	return &room, err
}

func GetRoomSelection(ctx *gin.Context) (*InterviewRoomSelection, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return nil, err
	}
	openid := token.Subject
	var roomSelection InterviewRoomSelection
	err = svc.MySQL.WithContext(ctx).Model(&InterviewRoomSelection{
		OpenId: openid,
	}).First(&roomSelection).Error
	if err != nil {
		return nil, err
	}
	return &roomSelection, nil
}

func CheckRoomSelected(ctx *gin.Context) (bool, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return false, err
	}
	openid := token.Subject
	var counts int64
	err = svc.MySQL.WithContext(ctx).Model(
		&InterviewRoomSelection{
			OpenId: openid,
		}).Count(&counts).Error
	if err != nil {
		return false, err
	}
	if counts == 0 {
		return false, nil
	}
	return true, nil
}

func SetRoom(ctx *gin.Context, roomSelection *InterviewRoomSelection) error {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return err
	}
	openid := token.Subject
	roomSelection.OpenId = openid
	// 已经通过Handler检测是否已经选择过面试场地
	err = svc.MySQL.WithContext(ctx).Create(roomSelection).Error
	if err != nil {
		return err
	}
	targetRoom := InterviewRoom{
		RoomId: roomSelection.RoomId,
	}
	// 然后再更新InterviewRoom
	err = svc.MySQL.WithContext(ctx).Model(&InterviewRoom{
		RoomId: roomSelection.RoomId,
	}).First(&targetRoom).Error
	if err != nil {
		return err
	}
	targetRoom.CurrentOccupancy++
	err = svc.MySQL.WithContext(ctx).Save(&targetRoom).Error
	return err
}

func DeleteRoom(ctx *gin.Context) error {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return err
	}
	openid := token.Subject
	roomSelection := InterviewRoomSelection{
		OpenId: openid,
	}
	err = svc.MySQL.WithContext(ctx).First(&roomSelection).Error
	if err != nil {
		return err
	}
	roomId := roomSelection.RoomId
	err = svc.MySQL.WithContext(ctx).Delete(&roomSelection).Error
	if err != nil {
		return err
	}
	// 同时再从InterviewRoom中减去1
	targetRoom := InterviewRoom{
		RoomId: roomId,
	}
	err = svc.MySQL.WithContext(ctx).First(&targetRoom).Error
	if err != nil {
		return err
	}
	targetRoom.CurrentOccupancy--
	err = svc.MySQL.WithContext(ctx).Save(&targetRoom).Error
	return err
}
