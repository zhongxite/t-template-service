package main

import (
	"zhongxite/t-template/cmd"
)

func main() {
	/*
		gin：go get github.com/gin-gonic/gin
		跨域：go get github.com/gin-contrib/cors
		gorm： go get gorm.io/gorm
		mysql数据库：go get gorm.io/driver/mysql
		redis：go get github.com/redis/go-redis/v9
		jwt：go get github.com/golang-jwt/jwt/v5
		文档导出导入：go get github.com/tealeg/xlsx
	*/
	cmd.Start()
	defer cmd.Clean()
}
