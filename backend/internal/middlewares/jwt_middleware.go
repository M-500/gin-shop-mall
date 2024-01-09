package middlewares

import (
	"backend/internal/config"
	"backend/pkg/constant"
	"backend/pkg/response"
	"backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	jwtConfig := config.ConfigInstance.Jwt
	return func(c *gin.Context) {
		// 从请求头中获取token
		token := c.Request.Header.Get(jwtConfig.JwtHeaderKey)
		if utils.IsBlank(token) {
			response.JsonFailCode(c, constant.NOT_LOGIN)
			c.Abort()
			return
		}
		tokenClean := strings.Replace(token, jwtConfig.JwtPrefix, "", -1)
		newJWT := utils.NewJwt()
		claims, err := newJWT.ParseToken(tokenClean)
		if err != nil {
			response.JsonFailMsg(c, err.Error())
			c.Abort()
			return
		}
		// 将用户信息插入上下文中
		c.Set(constant.JWT_INFO_UNAME, claims.UserName)
		c.Set(constant.JWT_INFO_EMAIL, claims.Email)
		c.Set(constant.JWT_INFO_UID, claims.ID)
		c.Next()
	}
}
