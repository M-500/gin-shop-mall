package na

import (
	"backend/pkg/response"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/8 15:21
//

func PwdLoginHandler(ctx *gin.Context) {
	response.JsonSuccessData(ctx, "成功", gin.H{})
}
