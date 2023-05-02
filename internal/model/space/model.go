package space

import "gorm.io/gorm"

// Space 空间
//
// 空间是OneELAB组织人员的最小单位。
type Space struct {
	gorm.Model `json:"-"` // 隐藏gorm.Model
	// SpaceId 空间的唯一标识符
	SpaceId string `json:"space_id" binding:"uuid"`
	// Name 空间的名称
	Name string `json:"name"`
	// Description 空间的描述
	Description string `json:"description"`
	// Private 空间是否为私有空间
	Private bool `json:"private"`
}

type OperationResponse struct {
	Ok bool `json:"ok"`
}
