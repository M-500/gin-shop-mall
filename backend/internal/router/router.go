package router

import (
	"backend/internal/middlewares"
	"backend/internal/router/routers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	Router := gin.New()
	Router.Use(middlewares.CorsMiddleWare())       // 全局使用跨域中间件
	Router.Use(middlewares.PaginationMiddleware()) // 全局使用分页中间件
	naRouter := Router.Group("/api/v1/na")

	routers.InitNaRouter(naRouter)
	return Router
}
