package models

import "time"

type CategoryModel struct {
	BaseModel
	ShopId       int64      `gorm:"comment:;column:shop_id" json:"shop_id"`
	ParentId     int64      `gorm:"comment:;column:parent_id" json:"parent_id"`
	CategoryName string     `gorm:"comment:;column:category_name" json:"category_name"`
	Icon         string     `gorm:"comment:;column:icon" json:"icon"`
	Picture      string     `gorm:"comment:;column:picture" json:"picture"`
	SeqNum       string     `gorm:"comment:;column:seq_num" json:"seq_num"`
	Status       int8       `gorm:"comment:;column:status" json:"status"`
	RecTime      *time.Time `gorm:"comment:;column:rec_time" json:"rec_time"`
	Grade        int8       `gorm:"comment:;column:grade" json:"grade"`
}

func (CategoryModel) TableName() string {
	return "tz_category"
}

type CategoryToBrand struct {
	ID         int64 `gorm:"primarykey;type bigint;comment:主键自增长ID" json:"id"`
	CategoryID int64 `gorm:"comment:分类ID;column:category_id"`
	BrandID    int64 `gorm:"comment:品牌ID;column:brand_id"`
}

func (CategoryToBrand) TableName() string {
	return "tz_category_brand"
}
