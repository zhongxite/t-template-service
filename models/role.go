package models

type Role struct {
	Model
	Name         string `json:"name" form:"name"`                 // 名称
	Power        int8   `json:"power" form:"power"`               // 权力
	Status       uint8  `json:"status" form:"status"`             // 状态
	MenusList    string `json:"menusList" from:"menusList"`       // 路由列表
	RegMenusList string `json:"regMenusList" from:"regMenusList"` // 回显路由列表
}

func (r *Role) TableName() string {
	return "role"
}
