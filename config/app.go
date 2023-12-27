package config

import "github.com/gin-gonic/gin"

var ConfigMysql = gin.H{
	"name":     "root",
	"password": "123456",
	"url":      "127.0.0.1",
	"port":     "3306",
	"dbName":   "t_template",
}
var ConfigRedis = map[string]interface{}{
	"addr":     "localhost:6379",
	"password": "",
	"db":       0,
}
var ConfigRouter = map[string]interface{}{
	"host": "8080",
}
