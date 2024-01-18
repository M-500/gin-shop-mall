package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:47
//

type TzCategory struct {
	CategoryId   uint64     `xorm:"not null pk autoincr comment('类目ID') UNSIGNED BIGINT" json:"categoryId"`
	ShopId       int64      `xorm:"not null comment('店铺ID') index BIGINT" json:"shopId"`
	ParentId     uint64     `xorm:"not null comment('父节点') index UNSIGNED BIGINT" json:"parentId"`
	CategoryName string     `xorm:"not null default '' comment('产品类目名称') VARCHAR(50)" json:"categoryName"`
	Icon         string     `xorm:"comment('类目图标') VARCHAR(255)" json:"icon"`
	Pic          string     `xorm:"comment('类目的显示图片') VARCHAR(300)" json:"pic"`
	Seq          int        `xorm:"not null comment('排序') INT" json:"seq"`
	Status       int        `xorm:"not null default 1 comment('默认是1，表示正常状态,0为下线状态') INT" json:"status"`
	RecTime      *time.Time `xorm:"not null comment('记录时间') DATETIME" json:"recTime"`
	Grade        int        `xorm:"not null comment('分类层级') INT" json:"grade"`
	UpdateTime   *time.Time `xorm:"comment('更新时间') DATETIME" json:"updateTime"`
}
