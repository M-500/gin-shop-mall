package sys_repositories

import (
	"backend/internal/entity"
	databasenani "backend/pkg/database"
	"gorm.io/gorm"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/9 19:05
//

type IMenuRepository interface {
	ListMenuByUserId(user *entity.SysUser) ([]entity.SysMenu, error)
}

type MenuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository() IMenuRepository {
	return &MenuRepository{DB: databasenani.GetDB()}
}

// 通过用户ID获取用户关联的菜单信息
func (r *MenuRepository) ListMenuByUserId(user *entity.SysUser) ([]entity.SysMenu, error) {
	/*
		select * from tz_sys_menu
		join tz_sys_role_menu on tz_sys_menu.id = tz_sys_role_menu.menu_id
		join tz_sys_user_role on tz_sys_user_role.role_id = tz_sys_role_menu.role_id
		join tz_sys_user on tz_sys_user_role.user_id = tz_sys_user.id
		where tz_sys_user.id=1
	*/
	var menuList []entity.SysMenu
	err := r.DB.Table("tz_sys_menu").
		Joins("join tz_sys_role_menu on tz_sys_menu.id = tz_sys_role_menu.menu_id").
		Joins("join tz_sys_user_role on tz_sys_user_role.role_id = tz_sys_role_menu.role_id").
		Joins("join tz_sys_user on tz_sys_user_role.user_id = tz_sys_user.id").
		Where("tz_sys_user.id=?", user.UserID).Distinct().Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	return menuList, nil
}
