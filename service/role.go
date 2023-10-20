package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
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
func RoleAddOrModify(c *gin.Context) {
	data := &models.Role{}
	data.Name = c.PostForm("name")
	if data.Name == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "名称不可为空",
		})
		return
	}
	status, err := strconv.Atoi(c.PostForm("status"))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "状态有误~",
			"err":  err.Error(),
		})
		return
	}
	data.Status = uint8(status)
	data.MenusList = c.PostForm("menusList")
	if data.MenusList == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "路由列表不可为空",
		})
		return
	}
	data.RegMenusList = c.PostForm("regMenusList")
	var resStr string
	var errStr string
	if c.PostForm("isChange") == "1" {
		uintId, _ := strconv.Atoi(c.PostForm("id"))
		data.ID = uint(uintId)
		err = common.DB.Model(data).Updates(data).Error
		resStr = "修改成功"
		errStr = "修改失败"
	} else {
		err = common.DB.Create(data).Error
		resStr = "新增成功"
		errStr = "新增失败"
	}
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  errStr,
			"err":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  resStr,
		})
	}
}
