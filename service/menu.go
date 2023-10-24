package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"zhongxite/t-template/common"
	"zhongxite/t-template/models"
)

func GetMenusList(c *gin.Context) {
	userJson, ok := c.Get("user")
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取用户信息失败",
		})
		return
	}
	user := userJson.(*models.User)
	if user.Role == 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取失败",
		})
		return
	}
	role := &models.Role{}
	role.ID = user.Role
	err := common.DB.First(role).Error
	if err != nil || role.MenusList == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取失败",
		})
		return
	}
	data := make([]models.Menus, 0)
	if user.Role == 1 {
		common.DB.Order("id").Find(&data)
	} else {
		arr := strings.Split(role.MenusList, ",")
		common.DB.Where("id IN ?", arr).Order("id").Find(&data)
	}

	list := initMenusLit(data, 0)
	c.JSON(200, gin.H{
		"code": 200,
		"data": list,
		"msg":  "获取成功",
	})
}
func initMenusLit(stuAll []models.Menus, pid uint64) []models.InitMenusList {
	var goodArr []models.InitMenusList
	for _, v := range stuAll {
		if v.Pid == pid {
			child := initMenusLit(stuAll, uint64(v.ID))
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
func MenusAddOrModify(c *gin.Context) {
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
	var err error
	data.Pid, err = strconv.ParseUint(c.PostForm("pid"), 10, 0)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "上级菜单不可为空",
		})
		return
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
func DeleteMenus(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "id不可为空",
		})
		return
	}
	uintId, _ := strconv.Atoi(c.PostForm("id"))
	data := &models.Menus{}
	data.ID = uint(uintId)
	count := common.DB.First(data).RowsAffected
	if count == 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "id不存在",
		})
		return
	}
	err := common.DB.Delete(data).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "删除失败",
			"err":  err.Error(),
		})
		return
	}
	list := make([]models.Menus, 0)
	count = common.DB.Where("pid = ?", uintId).Find(&list).RowsAffected
	if count != 0 {
		err := common.DB.Where("pid = ?", uintId).Delete(&list).Error
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "删除失败",
			})
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "删除成功",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	}

}
