// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameSmsLog = "tz_sms_log"

// SmsLog 短信记录表
type SmsLog struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID" json:"id"` // ID
	UserID       string    `gorm:"column:user_id;comment:用户id" json:"user_id"`                   // 用户id
	UserPhone    string    `gorm:"column:user_phone;not null;comment:手机号码" json:"user_phone"`    // 手机号码
	Content      string    `gorm:"column:content;not null;comment:短信内容" json:"content"`          // 短信内容
	MobileCode   string    `gorm:"column:mobile_code;not null;comment:手机验证码" json:"mobile_code"` // 手机验证码
	Type         int32     `gorm:"column:type;not null;comment:短信类型  1:注册  2:验证" json:"type"`    // 短信类型  1:注册  2:验证
	RecDate      time.Time `gorm:"column:rec_date;not null;comment:发送时间" json:"rec_date"`        // 发送时间
	ResponseCode string    `gorm:"column:response_code;comment:发送短信返回码" json:"response_code"`    // 发送短信返回码
	Status       int32     `gorm:"column:status;not null;comment:状态  1:有效  0：失效" json:"status"`  // 状态  1:有效  0：失效
}

// TableName SmsLog's table name
func (*SmsLog) TableName() string {
	return TableNameSmsLog
}
