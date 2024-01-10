package sys_controller

import "backend/internal/service"

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
