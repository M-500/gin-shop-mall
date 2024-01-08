package middlewares

import (
	"backend/internal/forms"
	"backend/pkg/constant"
	"github.com/gin-gonic/gin"
)

func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从请求中提取分页对象
		var pagination forms.PaginateForm
		if err := c.ShouldBindQuery(&pagination); err == nil {
			// 解析参数
			if pagination.PageSize < 1 || pagination.PageSize > 100 {
				pagination.PageSize = 15
			}
			if pagination.PageNum < 1 {
				pagination.PageNum = 1
			}
			c.Set(constant.PAGINATION_KEY, pagination)
		}
		c.Next() // 放行继续执行后面的逻辑
	}
}
