package routers

import (
	"backend/internal/handlers/admin/uses_handlers"
	"backend/internal/handlers/na"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/8 15:55
//

func InitNaRouter(router *gin.RouterGroup) {
	UserController := uses_handlers.NewUserController()
	naRouter := router.Group("/")
	{
		naRouter.GET("captcha", na.GenerateCaptchaHandler)
		naRouter.POST("login", na.PwdLoginHandler)
		naRouter.POST("register", UserController.CreateUserHandler)
	}
}
