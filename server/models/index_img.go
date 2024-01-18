package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:49
//

type TzIndexImg struct {
	ImgId      uint64     `xorm:"not null pk autoincr comment('主键') UNSIGNED BIGINT" json:"imgId"`
	ShopId     int64      `xorm:"comment('店铺ID') index BIGINT" json:"shopId"`
	ImgUrl     string     `xorm:"not null comment('图片') VARCHAR(200)" json:"imgUrl"`
	Des        string     `xorm:"default '' comment('说明文字,描述') VARCHAR(200)" json:"des"`
	Title      string     `xorm:"comment('标题') VARCHAR(200)" json:"title"`
	Link       string     `xorm:"comment('链接') VARCHAR(200)" json:"link"`
	Status     int        `xorm:"default 0 comment('状态') TINYINT(1)" json:"status"`
	Seq        int        `xorm:"default 0 comment('顺序') INT" json:"seq"`
	UploadTime *time.Time `xorm:"comment('上传时间') DATETIME" json:"uploadTime"`
	Relation   int64      `xorm:"comment('关联') BIGINT" json:"relation"`
	Type       int        `xorm:"comment('类型') INT" json:"type"`
}
