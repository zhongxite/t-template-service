package models

import (
	"gorm.io/gorm"
)

type Model struct {
	ID      uint           `gorm:"primaryKey" json:"id" form:"id"`
	Created int64          `gorm:"autoCreateTime:milli" json:"created" form:"created"` // 使用时间戳秒数填充创建时间
	Updated int64          `gorm:"autoUpdateTime:milli" json:"updated" form:"updated"` // 使用时间戳毫秒数填充更新时间
	Deleted gorm.DeletedAt `gorm:"index" json:"deleted" form:"deleted"`
}
