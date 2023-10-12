package router

import (
	"github.com/gin-gonic/gin"
	"zhongxite/t-template/middleware"
	"zhongxite/t-template/service"
)

func MenusRouter(r *gin.RouterGroup) {
	route := r.Group("/menu")
	route.POST("/getMenusList", service.GetMenusList)
	route.POST("/addMenus", middleware.Auth(), service.AddMenus)
}
