// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameUserAddr = "tz_user_addr"

// UserAddr 用户配送地址
type UserAddr struct {
	AddrID     int64     `gorm:"column:addr_id;primaryKey;autoIncrement:true;comment:ID" json:"addr_id"` // ID
	UserID     string    `gorm:"column:user_id;not null;default:0;comment:用户ID" json:"user_id"`          // 用户ID
	Receiver   string    `gorm:"column:receiver;comment:收货人" json:"receiver"`                            // 收货人
	ProvinceID int64     `gorm:"column:province_id;comment:省ID" json:"province_id"`                      // 省ID
	Province   string    `gorm:"column:province;comment:省" json:"province"`                              // 省
	City       string    `gorm:"column:city;comment:城市" json:"city"`                                     // 城市
	CityID     int64     `gorm:"column:city_id;comment:城市ID" json:"city_id"`                             // 城市ID
	Area       string    `gorm:"column:area;comment:区" json:"area"`                                      // 区
	AreaID     int64     `gorm:"column:area_id;comment:区ID" json:"area_id"`                              // 区ID
	PostCode   string    `gorm:"column:post_code;comment:邮编" json:"post_code"`                           // 邮编
	Addr       string    `gorm:"column:addr;comment:地址" json:"addr"`                                     // 地址
	Mobile     string    `gorm:"column:mobile;comment:手机" json:"mobile"`                                 // 手机
	Status     int32     `gorm:"column:status;not null;comment:状态,1正常，0无效" json:"status"`                // 状态,1正常，0无效
	CommonAddr int32     `gorm:"column:common_addr;not null;comment:是否默认地址 1是" json:"common_addr"`       // 是否默认地址 1是
	CreateTime time.Time `gorm:"column:create_time;not null;comment:建立时间" json:"create_time"`            // 建立时间
	Version    int32     `gorm:"column:version;not null;comment:版本号" json:"version"`                     // 版本号
	UpdateTime time.Time `gorm:"column:update_time;not null;comment:更新时间" json:"update_time"`            // 更新时间
}

// TableName UserAddr's table name
func (*UserAddr) TableName() string {
	return TableNameUserAddr
}
