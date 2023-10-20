package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
	"zhongxite/t-template/common"
	"zhongxite/t-template/models"
	"zhongxite/t-template/utils"
)

func CreateUser(c *gin.Context) {
	user := &models.User{}
	accountName := c.PostForm("accountName")
	name := c.PostForm("name")
	password := c.PostForm("password")
	if name == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "名称不可为空",
		})
		return
	}
	user.Name = name
	if accountName == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "账号名不可为空",
		})
		return
	}
	user.AccountName = accountName
	count := common.DB.Where("account_name = ?", accountName).First(user).RowsAffected
	if count > 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "账号名已存在",
		})
		return
	}
	if password == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "密码不可为空",
		})
		return
	}
	salt := string(rand.Int31())                       // 获取随机数
	user.Password = utils.MakePassword(password, salt) // 使用md5加密密码
	user.Salt = salt
	user.Status = 1
	common.DB.Create(user)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "创建成功",
	})
}
func DeleteUser(c *gin.Context) {
	id := c.PostForm("id")
	if id == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "id不可为空",
		})
		return
	}
	uintId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "删除失败",
		})
		return
	}
	user := &models.User{}
	user.ID = uint(uintId)
	count := common.DB.First(user).RowsAffected
	if count == 0 {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "用户不存在",
		})
		return
	}
	err = common.DB.Delete(user).Error
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
}
func Login(c *gin.Context) {
	user := &models.User{}
	accountName := c.PostForm("accountName")
	password := c.PostForm("password")
	if accountName == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "账号名不可为空",
		})
		return
	}
	user.AccountName = accountName
	common.DB.Where("account_name = ?", accountName).Find(user)
	fmt.Println(accountName, user.ID)
	if user.ID == 0 {
		fmt.Println(user.Name)
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "账号名不存在",
		})
		return
	}
	flag := utils.ValidPassword(password, user.Salt, user.Password) // 使用md5解密密码
	if !flag {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "密码错误",
		})
		return
	}
	token, err := utils.GenerateTokenUsingHs256(user) // 获取token
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取token失败",
		})
		return
	}
	user.LoginTime = time.Now().UnixNano() / 1e6
	user.LoginIp = c.ClientIP()
	err = common.DB.Model(user).Updates(user).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "登录失败",
			"err":  err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"token": token,
		},
		"msg": "登录成功",
	})
}

func GetUserInfo(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取用户信息失败",
		})
		return
	}
	data := user.(*models.User)
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"id":          data.ID,
			"accountName": data.AccountName,
			"avatar":      data.Avatar,
			"email":       data.Email,
			"name":        data.Name,
			"phone":       data.Phone,
			"sex":         data.Sex,
			"role":        data.Role,
		},
		"msg": "获取成功",
	})
}

func UpdateUserInfo(c *gin.Context) {
	user := &models.User{}
	id := c.PostForm("id")
	if id == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "用户id不可为空！",
		})
		return
	}
	err := common.DB.First(user, id).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "用户不存在",
			"err":  err.Error(),
		})
		return
	}
	user.Avatar = c.PostForm("avatar")
	user.Name = c.PostForm("name")
	user.Sex, _ = strconv.Atoi(c.PostForm("sex"))
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	err = common.DB.Model(user).Updates(user).Error
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "修改失败",
			"err":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}
func GetUserList(c *gin.Context) {
	data := []models.User{}
	form := &models.User{}
	err := c.ShouldBindQuery(form)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "获取参数失败",
			"err":  err.Error(),
		})
	}
	createStartTime := c.Query("createStartTime")
	createEndTime := c.Query("createEndTime")
	loginStartTime := c.Query("loginStartTime")
	loginEndTime := c.Query("loginEndTime")

	dbSql := common.DB.Where(form)
	if createStartTime != "" || createEndTime != "" {
		dbSql = dbSql.Where("created_at BETWEEN ? AND ?", createStartTime, createEndTime)
	}
	if loginStartTime != "" || loginEndTime != "" {
		dbSql = dbSql.Where("login_time BETWEEN ? AND ?", loginStartTime, loginEndTime)
	}
	var total int64
	limit, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	offset := (pageNum - 1) * limit
	dbSql.Limit(limit).Offset(offset).Order("id asc").Find(&data).Offset(-1).Limit(-1).Count(&total)
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"list":  data,
			"total": total,
		},
		"msg": "获取成功",
	})
}
