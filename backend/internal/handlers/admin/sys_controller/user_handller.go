package sys_controller

import (
	"backend/internal/forms"
	"backend/internal/service"
	"backend/pkg/response"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/8 15:18
//

type SysController struct {
	userService service.IUserService
	menuService service.IMenuService
}

func NewSysController() *SysController {
	return &SysController{
		userService: service.NewUserService(),
		menuService: service.NewMenuService(),
	}
}

/* 创建后台用户 */
func (u *SysController) CreateUserHandler(c *gin.Context) {
	resp := forms.CreateAdminUserForm{}
	err := c.ShouldBindJSON(&resp)
	if err != nil {
		return
	}
	user, err := u.userService.CreateUser(&resp)
	if err != nil {
		response.JsonFailMsg(c, err.Error())
		return
	}
	response.JsonSuccessData(c, "成功！", user)
	return
}

// 获取用户信息
func (u *SysController) AdminUserInfo(c *gin.Context) {
	response.JsonSuccessData(c, "成功！", nil)
}
