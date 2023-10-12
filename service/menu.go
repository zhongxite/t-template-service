package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"zhongxite/t-template/common"
	"zhongxite/t-template/models"
)

func GetMenusList(c *gin.Context) {
	data := make([]models.Menus, 0)
	common.DB.Order("id").Find(&data)
	smida := test1(data, 0)
	c.JSON(200, gin.H{
		"code": 200,
		"data": smida,
		"msg":  "获取成功",
	})
}
func test1(stuAll []models.Menus, pid uint64) []models.InitMenusList {
	var goodArr []models.InitMenusList
	for _, v := range stuAll {
		if v.Pid == pid {
			child := test1(stuAll, uint64(v.ID))
			node := models.InitMenusList{
				Model:     models.Model{ID: v.ID, Created: v.Created, Updated: v.Updated, Deleted: v.Deleted},
				Type:      v.Type,
				Pid:       v.Pid,
				Name:      v.Name,
				Mark:      v.Mark,
				Path:      v.Path,
				Icon:      v.Icon,
				Component: v.Component,
				Meta: gin.H{
					"title": v.Title,
				},
				Status:   v.Status,
				Children: child,
			}
			goodArr = append(goodArr, node)
		}
	}
	return goodArr
}
func AddMenus(c *gin.Context) {
	data := &models.Menus{}
	data.Name = c.PostForm("name")
	if data.Name == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "名称不可为空",
		})
		return
	}
	data.Type = c.PostForm("type")
	if data.Type == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "类型不可为空",
		})
		return
	}
	if c.PostForm("pid") != "" {
		data.Pid, _ = strconv.ParseUint(c.PostForm("uppId"), 10, 0)
	} else {
		data.Pid = 0
	}
	data.Path = c.PostForm("path")
	if data.Path == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "组件路由不可为空",
		})
		return
	}
	data.Mark = c.PostForm("mark")
	if data.Mark == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "标识不可为空",
		})
		return
	}
	data.Icon = c.PostForm("icon")
	if data.Icon == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "图标不可为空",
		})
		return
	}
	data.Component = c.PostForm("component")
	if (data.Type == "1" || data.Type == "2") && data.Component == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "组件路径不可为空",
		})
		return
	}
	data.Title = c.PostForm("title")
	data.Status = c.PostForm("status")
	err := common.DB.Create(data).Error
	fmt.Println(err, "================")
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "新增失败",
			"err":  err,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "新增成功",
		})
	}
}
