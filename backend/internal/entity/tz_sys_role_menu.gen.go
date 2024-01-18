// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSysRoleMenu = "tz_sys_role_menu"

// SysRoleMenu 角色与菜单对应关系
type SysRoleMenu struct {
	ID     int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RoleID int64 `gorm:"column:role_id;comment:角色ID" json:"role_id"` // 角色ID
	MenuID int64 `gorm:"column:menu_id;comment:菜单ID" json:"menu_id"` // 菜单ID
}

// TableName SysRoleMenu's table name
func (*SysRoleMenu) TableName() string {
	return TableNameSysRoleMenu
}