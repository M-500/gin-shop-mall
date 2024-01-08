package na

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/8 15:21
//

func PwdLoginHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"accessToken": "hahahaah",
	})
}
