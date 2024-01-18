package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:50
//

type TzMessage struct {
	Id         uint64     `xorm:"not null pk autoincr UNSIGNED BIGINT" json:"id"`
	CreateTime *time.Time `xorm:"comment('留言创建时间') DATETIME" json:"createTime"`
	UserName   string     `xorm:"comment('姓名') VARCHAR(36)" json:"userName"`
	Email      string     `xorm:"comment('邮箱') VARCHAR(64)" json:"email"`
	Contact    string     `xorm:"comment('联系方式') VARCHAR(36)" json:"contact"`
	Content    string     `xorm:"comment('留言内容') TEXT" json:"content"`
	Reply      string     `xorm:"comment('留言回复') TEXT" json:"reply"`
	Status     int        `xorm:"comment('状态：0:未审核 1审核通过') TINYINT" json:"status"`
}
