package models

type SysRoleModel struct {
	BaseModel
	RoleName     string `gorm:"type: varchar(100);comment:角色名称;column:role_name" json:"roleName"`
	Remark       string `gorm:"type:varchar(100);comment:备注;column:remark" json:"remark"`
	CreateUserId int64  `gorm:"type:bigint;comment:创建人ID;column:create_user_id" json:"createUserId"`
}

func (SysRoleModel) TableName() string {
	return "m4_sys_role"
}

type SysUserRoleModel struct {
	BaseModel
	UserId int64 `gorm:"type:bigint;column:user_id" json:"userId"`
	RoleId int64 `gorm:"type:bigint;column:role_id" json:"roleId"`
}

func (SysUserRoleModel) TableName() string {
	return "m4_sys_user_role"
}
