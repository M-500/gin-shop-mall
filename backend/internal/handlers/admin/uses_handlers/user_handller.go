package uses_handlers

import (
	"backend/internal/forms"
	"backend/internal/repositories/users_repositories"
	"backend/internal/service"
	"backend/pkg/response"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/8 15:18
//

type UserController struct {
	userService service.IUserService
}

func NewUserController() *UserController {
	repo := users_repositories.NewUserRepository()
	return &UserController{
		userService: service.NewUserService(repo),
	}
}

/* 创建后台用户 */
func (u *UserController) CreateUserHandler(c *gin.Context) {
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
