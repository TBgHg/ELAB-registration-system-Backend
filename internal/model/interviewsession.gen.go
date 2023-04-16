// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameInterviewSession = "InterviewSession"

// InterviewSession mapped from table <InterviewSession>
type InterviewSession struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // 主键
	StartTime  time.Time `gorm:"column:start_time;not null" json:"start_time"`                           // 面试开始时间
	EndTime    time.Time `gorm:"column:end_time;not null" json:"end_time"`                               // 面试结束时间
	Location   string    `gorm:"column:location;not null" json:"location"`                               // 面试地点
	Capacity   int32     `gorm:"column:capacity;not null" json:"capacity"`                               // 可参加人数
	AppliedNum int32     `gorm:"column:applied_num;not null" json:"applied_num"`                         // 已报名人数
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // 最后修改时间
}

// TableName InterviewSession's table name
func (*InterviewSession) TableName() string {
	return TableNameInterviewSession
}