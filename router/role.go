package router

import (
	"github.com/gin-gonic/gin"
	"zhongxite/t-template/service"
)

func RoleRouter(r *gin.RouterGroup) {
	route := r.Group("/role")
	route.POST("/getRoleList", service.GetRoleList)
}
