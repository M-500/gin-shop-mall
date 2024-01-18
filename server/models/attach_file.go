package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:46
//

type TzAttachFile struct {
	FileId       int64      `xorm:"not null pk autoincr BIGINT" json:"fileId"`
	FilePath     string     `xorm:"comment('文件路径') VARCHAR(255)" json:"filePath"`
	FileType     string     `xorm:"comment('文件类型') VARCHAR(20)" json:"fileType"`
	FileSize     int        `xorm:"comment('文件大小') INT" json:"fileSize"`
	UploadTime   *time.Time `xorm:"comment('上传时间') DATETIME" json:"uploadTime"`
	FileJoinId   int64      `xorm:"comment('文件关联的表主键id') BIGINT" json:"fileJoinId"`
	FileJoinType int        `xorm:"comment('文件关联表类型：1 商品表  FileJoinType') TINYINT" json:"fileJoinType"`
}
