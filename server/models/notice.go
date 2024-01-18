package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:50
//

type TzNotice struct {
	Id          int64      `xorm:"pk autoincr comment('公告id') BIGINT" json:"id"`
	ShopId      int64      `xorm:"comment('店铺id') BIGINT" json:"shopId"`
	Title       string     `xorm:"comment('公告标题') VARCHAR(36)" json:"title"`
	Content     string     `xorm:"comment('公告内容') TEXT" json:"content"`
	Status      int        `xorm:"comment('状态(1:公布 0:撤回)') TINYINT(1)" json:"status"`
	IsTop       int        `xorm:"comment('是否置顶') TINYINT" json:"isTop"`
	PublishTime *time.Time `xorm:"comment('发布时间') TIMESTAMP" json:"publishTime"`
	UpdateTime  *time.Time `xorm:"default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP" json:"updateTime"`
}
