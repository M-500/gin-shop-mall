package na

import (
	"backend/internal/service/na"
	"backend/pkg/response"
	"github.com/gin-gonic/gin"
)

func GenerateCaptchaHandler(c *gin.Context) {
	captcha, err := na.CaptchaSer.MakeCaptcha()
	if err != nil {
		response.JsonFailMsg(c, "验证码生成错误")
		return
	}
	response.JsonSuccessData(c, "成功", captcha)
	return

}
