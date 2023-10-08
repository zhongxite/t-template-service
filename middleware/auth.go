package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"zhongxite/t-template/common"
	"zhongxite/t-template/models"
	"zhongxite/t-template/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenRequst := c.Request.Header.Get("Authorization") // 获取token
		if tokenRequst == "" {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "token为空",
			})
			c.Abort()
			return
		}
		if len(tokenRequst) < 7 || !strings.HasPrefix(tokenRequst, "Bearer") {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "非法token",
			})
			c.Abort()
			return
		}
		tokenRequst = strings.Fields(tokenRequst)[1] // 分割空白获取Bearer后的token值
		token, userId, err := utils.ParseTokenHs256(tokenRequst)
		if err != nil || !token.Valid {
			c.JSON(200, gin.H{
				"code": 401,
				"msg":  "token无效",
			})
			c.Abort()
			return
		}
		user := &models.User{}
		common.DB.Where("id = ?", userId).Find(&user)
		c.Set("user", user)
		c.Next()
	}
}
