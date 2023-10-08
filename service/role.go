package service

import (
	"github.com/gin-gonic/gin"
	"zhongxite/t-template/common"
	"zhongxite/t-template/models"
)

func GetRoleList(c *gin.Context) {
	data := []models.Role{}
	common.DB.Find(&data)
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "获取成功",
	})
}
