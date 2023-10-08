package common

import (
	"gorm.io/gorm"
	"zhongxite/t-template/config"
)

var (
	DB *gorm.DB
	RC *config.Client
)
