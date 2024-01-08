//@Author: wulinlin
//@Description:
//@File:  sys_menu_model
//@Version: 1.0.0
//@Date: 2024/01/08 20:33

package models

type SysMenuModel struct {
	BaseModel
	ParentId int64
	Name     string
	Url      string
	Perms    string
	Type     int8
	Icon     string
	OrderNum int
}

func (SysMenuModel) TableName() string {
	return "sys_menu"
}
