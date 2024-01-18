// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameMessage = "tz_message"

// Message mapped from table <tz_message>
type Message struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;comment:留言创建时间" json:"create_time"` // 留言创建时间
	UserName   string    `gorm:"column:user_name;comment:姓名" json:"user_name"`         // 姓名
	Email      string    `gorm:"column:email;comment:邮箱" json:"email"`                 // 邮箱
	Contact    string    `gorm:"column:contact;comment:联系方式" json:"contact"`           // 联系方式
	Content    string    `gorm:"column:content;comment:留言内容" json:"content"`           // 留言内容
	Reply      string    `gorm:"column:reply;comment:留言回复" json:"reply"`               // 留言回复
	Status     int32     `gorm:"column:status;comment:状态：0:未审核 1审核通过" json:"status"`   // 状态：0:未审核 1审核通过
}

// TableName Message's table name
func (*Message) TableName() string {
	return TableNameMessage
}