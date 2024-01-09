package dto

import (
	"gorm.io/gorm"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024/1/9 20:16
type SysMenuDTO struct {
	ID        int64          `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	ParentId  int64          `json:"parentId"`
	Name      string         `json:"name"`
	Url       string         `json:"url"`
	Perms     string         `json:"perms"`
	Type      int            `json:"type"`
	Icon      string         `json:"icon"`
	OrderNum  int            `json:"orderNum"`
	List      *[]SysMenuDTO  `json:"list"`
}

type NavDTO struct {
	MenuList    []SysMenuDTO `json:"menuList"`
	Authorities []string     `json:"authorities"`
}
