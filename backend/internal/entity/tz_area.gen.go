// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameArea = "tz_area"

// Area mapped from table <tz_area>
type Area struct {
	AreaID   int64  `gorm:"column:area_id;primaryKey;autoIncrement:true" json:"area_id"`
	AreaName string `gorm:"column:area_name" json:"area_name"`
	ParentID int64  `gorm:"column:parent_id" json:"parent_id"`
	Level    int32  `gorm:"column:level" json:"level"`
}

// TableName Area's table name
func (*Area) TableName() string {
	return TableNameArea
}
