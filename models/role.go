package models

type Role struct {
	Model
	Name   string `json:"name" form:"name"`     // 名称
	Power  uint8  `json:"power" form:"power"`   // 权力
	Status bool   `json:"status" form:"status"` // 状态
	Role   uint8  `json:"role" form:"role"`     // 角色
}

func (u *Role) TableName() string {
	return "role"
}
