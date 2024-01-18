package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:47
//

type TzBrand struct {
	BrandId    uint64    `xorm:"not null pk autoincr comment('主键') UNSIGNED BIGINT" json:"brandId"`
	BrandName  string    `xorm:"not null default '' comment('品牌名称') VARCHAR(30)" json:"brandName"`
	BrandPic   string    `xorm:"comment('图片路径') VARCHAR(255)" json:"brandPic"`
	UserId     string    `xorm:"not null default '' comment('用户ID') VARCHAR(36)" json:"userId"`
	Memo       string    `xorm:"comment('备注') VARCHAR(50)" json:"memo"`
	Seq        int       `xorm:"default 1 comment('顺序') INT" json:"seq"`
	Status     int       `xorm:"not null default 1 comment('默认是1，表示正常状态,0为下线状态') INT" json:"status"`
	Brief      string    `xorm:"comment('简要描述') VARCHAR(100)" json:"brief"`
	Content    string    `xorm:"comment('内容') TEXT" json:"content"`
	RecTime    time.Time `xorm:"comment('记录时间') DATETIME" json:"recTime"`
	UpdateTime time.Time `xorm:"comment('更新时间') DATETIME" json:"updateTime"`
	FirstChar  string    `xorm:"comment('品牌首字母') VARCHAR(1)" json:"firstChar"`
}

type TzCategoryBrand struct {
	Id         int64 `xorm:"pk autoincr BIGINT" json:"id"`
	CategoryId int64 `xorm:"comment('分类id') index BIGINT" json:"categoryId"`
	BrandId    int64 `xorm:"comment('品牌id') index BIGINT" json:"brandId"`
}
