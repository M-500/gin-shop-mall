package models

import (
	"time"
)

// Brand 模型（品牌品牌表）
type BrandModel struct {
	BaseModel
	BrandID    uint       `gorm:"primaryKey;autoIncrement" json:"brand_id"`
	BrandName  string     `gorm:"not null;size:30;default:''" json:"brand_name"`
	BrandPic   string     `gorm:"size:255" json:"brand_pic"`
	UserID     string     `gorm:"not null;size:36;default:''" json:"user_id"`
	Memo       string     `gorm:"size:50" json:"memo"`
	Seq        int        `gorm:"default:1" json:"seq"`
	Status     int        `gorm:"not null;default:1" json:"status"`
	Brief      string     `gorm:"size:100" json:"brief"`
	Content    string     `gorm:"type:text" json:"content"`
	RecTime    *time.Time `gorm:"type:datetime" json:"rec_time"`
	UpdateTime *time.Time `gorm:"type:datetime" json:"update_time"`
	FirstChar  string     `gorm:"size:1" json:"first_char"`
}

// TableName 指定表名
func (BrandModel) TableName() string {
	return "tz_brand"
}
