// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameApplication = "Application"

// Application mapped from table <Application>
type Application struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                      // 主键
	UserID       int32     `gorm:"column:user_id;not null" json:"user_id"`                                 // 用户ID
	CreateTime   time.Time `gorm:"column:create_time;not null" json:"create_time"`                         // 报名时间
	InterviewID  int32     `gorm:"column:interview_id;not null" json:"interview_id"`                       // 面试场次ID
	State        int32     `gorm:"column:state;not null" json:"state"`                                     // 状态：0表示已取消，1表示正常状态
	InterviewRes int32     `gorm:"column:interview_res;not null" json:"interview_res"`                     // 面试结果：0表示评审中/未面试，-1表示未通过，1表示通过
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"` // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"` // 最后修改时间
}

// TableName Application's table name
func (*Application) TableName() string {
	return TableNameApplication
}
