package dto

import (
	"gorm.io/gorm"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024/1/10 9:20
type ProdTagRecord struct {
	ID        int64          `json:"id"`
	CreatedAt time.Time      `json:"createdTime"`
	UpdatedAt time.Time      `json:"updatedTime"`
	DeletedAt gorm.DeletedAt `json:"deletedTime"`
	Title     string         `json:"title"`
	ShopId    int64          `json:"shopId"`
	Status    int            `json:"status"`
	IsDefault int            `json:"isDefault"`
	ProdCount int64          `json:"prodCount"`
	Style     int            `json:"style"`
	Seq       int            `json:"seq"`
}

type ProdTagDto struct {
	Current int64
	Orders  []interface{}
	Pages   int64
	Records []ProdTagRecord
	Size    int64
	Total   int64
}
