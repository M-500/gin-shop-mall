package sys_repositories

import (
	"backend/internal/models"
	databasenani "backend/pkg/database"
	"gorm.io/gorm"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/9 19:05
//

type IMenuRepository interface {
	ListMenuByUserId(user *models.SysUserModel) ([]models.SysMenuModel, error)
}

type MenuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository() IMenuRepository {
	return &MenuRepository{DB: databasenani.GetDB()}
}

// 通过用户ID获取用户关联的菜单信息
func (r *MenuRepository) ListMenuByUserId(user *models.SysUserModel) ([]models.SysMenuModel, error) {
	/*
		select * from m4_sys_menu
		join m4_sys_role_menu on m4_sys_menu.id = m4_sys_role_menu.menu_id
		join m4_sys_user_role on m4_sys_user_role.role_id = m4_sys_role_menu.role_id
		join m4_sys_user on m4_sys_user_role.user_id = m4_sys_user.id
		where m4_sys_user.id=1
	*/
	var menuList []models.SysMenuModel
	err := r.DB.Table("m4_sys_menu").
		Joins("join m4_sys_role_menu on m4_sys_menu.id = m4_sys_role_menu.menu_id").
		Joins("join m4_sys_user_role on m4_sys_user_role.role_id = m4_sys_role_menu.role_id").
		Joins("join m4_sys_user on m4_sys_user_role.user_id = m4_sys_user.id").
		Where("m4_sys_user.id=?", user.ID).Distinct().Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	return menuList, nil
}
