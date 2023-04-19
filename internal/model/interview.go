package model

import "time"

type InterviewSelectReq struct {
	OldInterviewSessionID int32 `json:"old_interview_session_id"`
	InterviewSessionID    int32 `json:"interview_session_id"`
}

// Validate 验证参数，简单校验一下，就不上Validate库了
func (req *InterviewSelectReq) Validate() bool {
	if req.OldInterviewSessionID == 0 && req.InterviewSessionID == 0 {
		return false
	}
	if req.OldInterviewSessionID == req.InterviewSessionID {
		return false
	}
	if req.OldInterviewSessionID < 0 || req.InterviewSessionID < 0 {
		return false
	}
	return true
}

type InterviewSelectResp struct {
	*CommonResp
}

type InterviewGetResp struct {
	*CommonResp
	InterviewID int32     `json:"interview_id"` // 主键
	StartTime   time.Time `json:"start_time"`   // 面试开始时间
	EndTime     time.Time `json:"end_time"`     // 面试结束时间
	Location    string    `json:"location"`     // 面试地点
	Capacity    int32     `json:"capacity"`     // 可参加人数
	AppliedNum  int32     `json:"applied_num"`  // 已报名人数
}
