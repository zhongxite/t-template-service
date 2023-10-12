package cmd

import (
	"zhongxite/t-template/common"
	"zhongxite/t-template/config"
	"zhongxite/t-template/router"
)

func Start() {
	var err error
	common.DB, err = config.InitDB()
	if err != nil {
		panic("连接数据库失败：" + err.Error())
	}
	common.RC, err = config.InitRedis()
	if err != nil {
		panic("连接redis失败：" + err.Error())
	}
	router.InitRouter()
}
func Clean() {

}
