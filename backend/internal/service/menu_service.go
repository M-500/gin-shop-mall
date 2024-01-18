package service

import (
	"backend/internal/entity"
	"backend/internal/repositories/sys_repositories"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/9 19:19
//

type IMenuService interface {
	ListMenuByUserId(userId int64) ([]entity.SysMenu, error)
}

type MenuService struct {
	dao sys_repositories.IMenuRepository
}

func NewMenuService() IMenuService {
	return &MenuService{
		dao: sys_repositories.NewMenuRepository(),
	}
}

func (s *MenuService) ListMenuByUserId(id int64) ([]entity.SysMenu, error) {
	user := entity.SysUser{
		UserID: id,
	}
	return s.dao.ListMenuByUserId(&user)
}
