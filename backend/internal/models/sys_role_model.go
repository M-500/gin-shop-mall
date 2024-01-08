//@Author: wulinlin
//@Description:
//@File:  sys_role_model
//@Version: 1.0.0
//@Date: 2024/01/08 20:32

package models

type SysRoleModel struct {
	BaseModel
	RoleName     string `gorm:"type:;size:;comment:角色名称;column:role_name" json:"role_name"`
	Remark       string `gorm:"type:;size:;comment:描述;column:remark" json:"remark"`
	CreateUserId int64  `gorm:"type:;size:;comment:创建人ID;column:create_user_id" json:"create_user_id"`
}

func (SysRoleModel) TableName() string {
	return "sys_role"
}
