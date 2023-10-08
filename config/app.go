package config

import "github.com/gin-gonic/gin"

var ConfigMysql = gin.H{
	"name":     "root",
	"password": "Zxite998*",
	"url":      "localhost",
	"port":     "3306",
	"dbName":   "t_template",
}
var ConfigRedis = map[string]interface{}{
	"addr":     "localhost:6379",
	"password": "",
	"db":       0,
}
var ConfigRouter = map[string]interface{}{
	"host": "80",
}
