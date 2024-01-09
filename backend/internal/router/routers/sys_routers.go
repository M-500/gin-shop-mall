package routers

import (
	"backend/internal/handlers/admin/sys_controller"
	"github.com/gin-gonic/gin"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024/1/8 13:22
func InitSysRouter(router *gin.RouterGroup) {
	sysController := sys_controller.NewSysController()
	naRouter := router.Group("/")
	{
		naRouter.GET("/menu/nav", sysController.MenuNavHandler)
	}
}
