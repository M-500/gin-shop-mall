// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameHotSearch = "tz_hot_search"

// HotSearch 热搜
type HotSearch struct {
	HotSearchID int64     `gorm:"column:hot_search_id;primaryKey;autoIncrement:true;comment:主键" json:"hot_search_id"` // 主键
	ShopID      int64     `gorm:"column:shop_id;comment:店铺ID 0为全局热搜" json:"shop_id"`                                  // 店铺ID 0为全局热搜
	Content     string    `gorm:"column:content;not null;comment:内容" json:"content"`                                  // 内容
	RecDate     time.Time `gorm:"column:rec_date;not null;comment:录入时间" json:"rec_date"`                              // 录入时间
	Seq         int32     `gorm:"column:seq;comment:顺序" json:"seq"`                                                   // 顺序
	Status      int32     `gorm:"column:status;not null;default:1;comment:状态 0下线 1上线" json:"status"`                  // 状态 0下线 1上线
	Title       string    `gorm:"column:title;not null;comment:热搜标题" json:"title"`                                    // 热搜标题
}

// TableName HotSearch's table name
func (*HotSearch) TableName() string {
	return TableNameHotSearch
}