package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"zhongxite/t-template/config"
)

func InitRouter() {
	r := gin.Default()
	r.Use(cors.Default())
	api := r.Group("/api")
	UserRouter(api)
	CommonRouter(api)
	RoleRouter(api)
	MenusRouter(api)
	panic(r.Run(":" + config.ConfigRouter["host"].(string)))
}
