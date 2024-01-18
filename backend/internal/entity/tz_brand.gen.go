// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameBrand = "tz_brand"

// Brand 品牌表
type Brand struct {
	BrandID    int64     `gorm:"column:brand_id;primaryKey;autoIncrement:true;comment:主键" json:"brand_id"`  // 主键
	BrandName  string    `gorm:"column:brand_name;not null;comment:品牌名称" json:"brand_name"`                 // 品牌名称
	BrandPic   string    `gorm:"column:brand_pic;comment:图片路径" json:"brand_pic"`                            // 图片路径
	UserID     string    `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                       // 用户ID
	Memo       string    `gorm:"column:memo;comment:备注" json:"memo"`                                        // 备注
	Seq        int32     `gorm:"column:seq;default:1;comment:顺序" json:"seq"`                                // 顺序
	Status     int32     `gorm:"column:status;not null;default:1;comment:默认是1，表示正常状态,0为下线状态" json:"status"` // 默认是1，表示正常状态,0为下线状态
	Brief      string    `gorm:"column:brief;comment:简要描述" json:"brief"`                                    // 简要描述
	Content    string    `gorm:"column:content;comment:内容" json:"content"`                                  // 内容
	RecTime    time.Time `gorm:"column:rec_time;comment:记录时间" json:"rec_time"`                              // 记录时间
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                        // 更新时间
	FirstChar  string    `gorm:"column:first_char;comment:品牌首字母" json:"first_char"`                         // 品牌首字母
}

// TableName Brand's table name
func (*Brand) TableName() string {
	return TableNameBrand
}