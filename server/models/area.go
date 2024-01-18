package models

// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:45
type TzArea struct {
	AreaId   int64  `xorm:"not null pk autoincr BIGINT" json:"areaId"`
	AreaName string `xorm:"VARCHAR(50)" json:"areaName"`
	ParentId int64  `xorm:"index BIGINT" json:"parentId"`
	Level    int    `xorm:"INT" json:"level"`
}
