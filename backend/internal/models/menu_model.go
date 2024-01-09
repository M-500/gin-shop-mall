package models

type SysMenuModel struct {
	BaseModel
	ParentId int64  `gorm:"default:NULL;type:bigint;comment:父ID;column:parent_id" json:"parentId"`
	Name     string `gorm:"default:NULL;type:varchar(50);comment:菜单名称;column:name" json:"name"`
	Url      string `gorm:"default:NULL;type:varchar(200);comment:菜单URL;column:url" json:"url"`
	Perms    string `gorm:"default:NULL;type:varchar(500);comment:授权(多个用逗号分隔，如：user:list,user:create);column:perms" json:"perms"`
	Type     int    `gorm:"default:NULL;type:int;comment:类型   0：目录   1：菜单   2：按钮;column:type" json:"type"`
	Icon     string `gorm:"default:NULL;type:varchar(50);comment:菜单图标;column:icon" json:"icon"`
	OrderNum int    `gorm:"default:NULL;type:int;comment:排序;column:order_num" json:"orderNum"`
}

func (SysMenuModel) TableName() string {
	return "m4_sys_menu"
}

type SysRoleMenu struct {
	BaseModel
	RoleId int64 `gorm:"type:bigint;column:role_id" json:"roleId"`
	MenuId int64 `gorm:"type:bigint;column:menu_id" json:"menuId"`
}

func (SysRoleMenu) TableName() string {
	return "m4_sys_role_menu"
}
