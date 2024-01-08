package models

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/8 13:45
//

// Area 模型
type AreaModel struct {
	BaseModel
	AreaID   uint   `gorm:"primaryKey;autoIncrement" json:"area_id"`
	AreaName string `gorm:"type:varchar(50)" json:"area_name"`
	ParentID uint   `gorm:"index" json:"parent_id"`
	Level    int    `gorm:"type:int(1)" json:"level"`
}

// TableName 指定表名
func (AreaModel) TableName() string {
	return "tz_area"
}
