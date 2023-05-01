package application

import (
	"elab-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func CheckLongTextFormExists(ctx *gin.Context) (bool, error) {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	if err != nil {
		return false, err
	}
	openid := token.Subject
	var counts int64
	err = svc.MySQL.WithContext(ctx).Model(
		&LongTextForm{
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

func UpdateLongTextForm(ctx *gin.Context, form *LongTextForm) error {
	svc := service.GetService()
	token, err := svc.Oidc.GetToken(ctx)
	openid := token.Subject
	if err != nil {
		return err
	}
	// 先检查是否存在
	exists, err := CheckLongTextFormExists(ctx)
	if err != nil {
		return err
	}
	form.OpenId = openid
	if !exists {
		err = svc.MySQL.WithContext(ctx).Create(&form).Error
		if err != nil {
			return err
		}
		return nil
	}
	err = svc.MySQL.WithContext(ctx).Model(&LongTextForm{
		OpenId: openid,
	}).Save(form).Error
	return err
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
		Id: roomId,
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
		Id: roomSelection.RoomId,
	}
	// 然后再更新InterviewRoom
	err = svc.MySQL.WithContext(ctx).Model(&InterviewRoom{
		Id: roomSelection.RoomId,
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
		Id: roomId,
	}
	err = svc.MySQL.WithContext(ctx).First(&targetRoom).Error
	if err != nil {
		return err
	}
	targetRoom.CurrentOccupancy--
	err = svc.MySQL.WithContext(ctx).Save(&targetRoom).Error
	return err
}
