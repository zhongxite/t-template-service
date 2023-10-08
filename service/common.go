package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"path"
	"strconv"
	"strings"
	"time"
	"zhongxite/t-template/common"
	"zhongxite/t-template/models"
	"zhongxite/t-template/utils"
)

func UploadOss(c *gin.Context) {
	file, err := c.FormFile("file")
	url, err := utils.UploadOss(file)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{
			"url": url,
		},
		"msg": "上传成功",
	})
}
func ExportFile(c *gin.Context) {
	exportName := c.Query("exportName")
	if exportName == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "导出文件表名不可为空",
		})
		return
	}
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	if exportName == "user" {
		formData := &models.User{}
		err := c.ShouldBind(formData)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  "获取数据失败",
			})
		}
		type exportFileData struct {
			Name        string `json:"name" form:"name"`               // 姓名
			Avatar      string `json:"avatar" form:"avatar"`           // 头像
			Sex         int    `json:"sex" form:"sex"`                 // 性别
			Phone       string `json:"phone" form:"phone"`             // 手机号
			Email       string `json:"email" form:"email"`             // 邮箱
			AccountName string `json:"accountName" form:"accountName"` // 账号名
		}
		data := []exportFileData{}
		if startTime == "" || endTime == "" {
			common.DB.Table(exportName).Where(formData).Find(&data)
		} else {
			common.DB.Table(exportName).Where("created_at BETWEEN ? AND ?", startTime, endTime).Where(formData).Find(&data)
		}
		var returnData []interface{}
		for _, v := range data {
			returnData = append(returnData, &exportFileData{
				Name:        v.Name,
				Avatar:      v.Avatar,
				Sex:         v.Sex,
				Phone:       v.Phone,
				Email:       v.Email,
				AccountName: v.AccountName,
			})
		}
		content := utils.ToExcel([]string{`姓名`, `头像`, `性别`, `手机号`, `邮箱`, `账号名`}, returnData)
		utils.ResponseXls(c, content, "文件名")
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "导出文件表名错误",
		})
		return
	}

}
func ImportFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{
			"msg":  "导入失败",
			"code": 400,
		})
		return
	}
	extName := path.Ext(file.Filename)
	if extName != ".xlsx" {
		c.JSON(200, gin.H{
			"msg":  "文件格式错误，仅支持xlsx",
			"code": 400,
		})
		return
	}

	guid := time.Now().String()
	filePath := "static/files/" + guid + ".xlsx"
	if filePath == "" {
		c.JSON(200, gin.H{
			"msg":  "导入失败",
			"code": 400,
		})
		return
	}
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "导入失败",
		})
		return
	}
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "导入失败",
		})
		return
	}
	var userDataList = make([]models.User, 0)
	var existList = make([]string, 0)
	for _, sheet := range xlFile.Sheets {
		for index, row := range sheet.Rows {
			if index == 0 {
				continue
			}
			user := models.User{}
			user.Name = row.Cells[0].Value
			user.Avatar = row.Cells[1].Value
			user.Sex, _ = strconv.Atoi(row.Cells[2].Value)
			user.Phone = row.Cells[3].Value
			user.Email = row.Cells[4].Value
			user.AccountName = row.Cells[5].Value
			var count int64
			count = common.DB.Where("name = ?", user.Name).First(&user).RowsAffected
			if count > 0 {
				existList = append(existList, strconv.Itoa(index+1))
			} else {
				userDataList = append(userDataList, user)
			}
		}
	}
	if len(userDataList) > 0 {
		var createDataList = [][]models.User{}
		num := 1500
		lenNum := len(userDataList) / num
		if lenNum > 1500 {
			for i := 0; i < lenNum; i++ {
				if lenNum-1 == i {
					createDataList = append(createDataList, userDataList[num*i:len(userDataList)])
					break
				}
				createDataList = append(createDataList, userDataList[num*i:num*(i+1)])
			}
		} else {
			createDataList = append(createDataList, userDataList)
		}
		for _, data := range createDataList {
			err = common.DB.Create(data).Error
			if err != nil {
				c.JSON(200, gin.H{
					"code": 400,
					"msg":  "导入失败",
					"err":  err.Error(),
				})
				return
			}
		}
		if len(existList) > 0 {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "导入成功，重复数据行数为:" + strings.Join(existList, ","),
			})
			return
		} else {
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "导入成功",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "导入失败，重复的数据",
		})
		return
	}
}
