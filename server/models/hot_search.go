package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:49
//

type TzHotSearch struct {
	HotSearchId uint64     `xorm:"not null pk autoincr comment('主键') UNSIGNED BIGINT" json:"hotSearchId"`
	ShopId      int64      `xorm:"comment('店铺ID 0为全局热搜') index BIGINT" json:"shopId"`
	Content     string     `xorm:"not null comment('内容') VARCHAR(255)" json:"content"`
	RecDate     *time.Time `xorm:"not null comment('录入时间') DATETIME" json:"recDate"`
	Seq         int        `xorm:"comment('顺序') INT" json:"seq"`
	Status      int        `xorm:"not null default 1 comment('状态 0下线 1上线') TINYINT" json:"status"`
	Title       string     `xorm:"not null comment('热搜标题') VARCHAR(255)" json:"title"`
}
