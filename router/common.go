package router

import (
	"github.com/gin-gonic/gin"
	"zhongxite/t-template/middleware"
	"zhongxite/t-template/service"
)

func CommonRouter(router *gin.RouterGroup) {
	route := router.Group("/common")
	route.POST("/uploadOss", service.UploadOss)
	route.GET("/exportFile", middleware.Auth(), service.ExportFile)
	route.POST("/importFile", service.ImportFile)
}
