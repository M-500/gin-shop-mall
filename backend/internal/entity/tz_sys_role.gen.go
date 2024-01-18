// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameSysRole = "tz_sys_role"

// SysRole 角色
type SysRole struct {
	RoleID       int64     `gorm:"column:role_id;primaryKey;autoIncrement:true" json:"role_id"`
	RoleName     string    `gorm:"column:role_name;comment:角色名称" json:"role_name"`            // 角色名称
	Remark       string    `gorm:"column:remark;comment:备注" json:"remark"`                    // 备注
	CreateUserID int64     `gorm:"column:create_user_id;comment:创建者ID" json:"create_user_id"` // 创建者ID
	CreateTime   time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`        // 创建时间
}

// TableName SysRole's table name
func (*SysRole) TableName() string {
	return TableNameSysRole
}