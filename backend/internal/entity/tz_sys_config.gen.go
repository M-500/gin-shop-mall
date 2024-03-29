// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameSysConfig = "tz_sys_config"

// SysConfig 系统配置信息表
type SysConfig struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ParamKey   string `gorm:"column:param_key;comment:key" json:"param_key"`       // key
	ParamValue string `gorm:"column:param_value;comment:value" json:"param_value"` // value
	Remark     string `gorm:"column:remark;comment:备注" json:"remark"`              // 备注
}

// TableName SysConfig's table name
func (*SysConfig) TableName() string {
	return TableNameSysConfig
}
