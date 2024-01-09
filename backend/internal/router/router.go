package router

import (
	"backend/internal/middlewares"
	"backend/internal/router/routers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.CorsMiddleWare())       // 全局使用跨域中间件
	Router.Use(middlewares.PaginationMiddleware()) // 全局使用分页中间件
	naRouter := Router.Group("/api/v1/na")

	routers.InitNaRouter(naRouter)

	sysRouter := Router.Group("/sys")
	routers.InitSysRouter(sysRouter)
	return Router
}
