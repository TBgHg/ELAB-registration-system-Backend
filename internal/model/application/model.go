package application

import "gorm.io/gorm"

const InterviewRoomUriScheme = "interview-room://"

// LongTextForm 长文本表单
//
// 在Gin当中，会被用作Request
type LongTextForm struct {
	gorm.Model `json:"-"` // 不会被序列化
	// OpenId 用户OpenId
	OpenId string `json:"openid" binding:"required,uuid" gorm:"column:openid"`
	// 加入原因
	Reason string `json:"reason"`
	// 个人经历
	Experience string `json:"experience"`
	// 个人自我评价
	SelfEvaluation string `json:"self_evaluation"`
}

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

type UpdateLongTextFormResponse struct {
	Ok bool `json:"ok"`
}

type StatusResponse struct {
	LongTextForm  bool `json:"long_text_form"`
	RoomSelection bool `json:"room_selection"`
}
