package config

import "github.com/gin-gonic/gin"

var ConfigMysql = gin.H{
	"name":     "root",
	"password": "123",
	"url":      "8.134.204.77",
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
