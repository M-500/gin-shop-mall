package service

import (
	"backend/internal/models"
	"backend/internal/repositories/sys_repositories"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/9 19:19
//

type IMenuService interface {
	ListMenuByUserId(userId int64) ([]models.SysMenuModel, error)
}

type MenuService struct {
	repo sys_repositories.IMenuRepository
}

func NewMenuService() IMenuService {
	return &MenuService{
		repo: sys_repositories.NewMenuRepository(),
	}
}

func (s *MenuService) ListMenuByUserId(userId int64) ([]models.SysMenuModel, error) {
	user := models.SysUserModel{
		BaseModel: models.BaseModel{ID: userId},
	}
	return s.repo.ListMenuByUserId(&user)
}
