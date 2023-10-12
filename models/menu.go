package models

import "github.com/gin-gonic/gin"

type Menus struct {
	Model
	Type      string `json:"type" form:"type"`           // 类型：0菜单目录 1菜单项 2页面按钮 3外链
	Pid       uint64 `json:"pid" form:"pid"`             // 父级Id
	Name      string `json:"name" form:"name"`           // 名称
	Mark      string `json:"mark" form:"mark"`           // 路由标识
	Path      string `json:"path" form:"path"`           // 访问路由
	Icon      string `json:"icon" form:"icon"`           // 组件图标
	Component string `json:"component" form:"component"` // 组件路径
	Title     string `json:"title" form:"title"`         // 标题
	Status    string `json:"status" form:"status"`       // 是否启动：0不启动1启动
}
type InitMenusList struct {
	Model
	Type      string          `json:"type" form:"type"`           // 类型：0菜单目录 1菜单项 2页面按钮 3外链
	Pid       uint64          `json:"pid" form:"pid"`             // 父级Id
	Name      string          `json:"name" form:"name"`           // 名称
	Mark      string          `json:"mark" form:"mark"`           // 路由标识
	Path      string          `json:"path" form:"path"`           // 访问路由
	Icon      string          `json:"icon" form:"icon"`           // 组件图标
	Component string          `json:"component" form:"component"` // 组件路径
	Meta      gin.H           `json:"meta" form:"meta"`           // meta对象
	Status    string          `json:"status" form:"status"`       // 是否启动：0不启动1启动
	Children  []InitMenusList `json:"children" form:"children"`
}

func (m *Menus) TableName() string {
	return "menu"
}
