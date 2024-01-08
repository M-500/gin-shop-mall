package models

import (
	"time"
)

// AttachFile 模型
type AttachFileModel struct {
	BaseModel
	FileID       uint       `gorm:"primaryKey;autoIncrement" json:"file_id"`
	FilePath     string     `gorm:"type:varchar(255)" json:"file_path"`
	FileType     string     `gorm:"type:varchar(20)" json:"file_type"`
	FileSize     int        `gorm:"type:int(10)" json:"file_size"`
	UploadTime   *time.Time `gorm:"type:datetime" json:"upload_time"`
	FileJoinID   uint       `gorm:"column:file_join_id" json:"file_join_id"`
	FileJoinType int        `gorm:"type:tinyint(2)" json:"file_join_type"`
}

// TableName 指定表名
func (AttachFileModel) TableName() string {
	return "tz_attach_file"
}
