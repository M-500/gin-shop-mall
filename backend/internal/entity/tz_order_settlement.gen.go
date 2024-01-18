// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameOrderSettlement = "tz_order_settlement"

// OrderSettlement mapped from table <tz_order_settlement>
type OrderSettlement struct {
	SettlementID int64     `gorm:"column:settlement_id;primaryKey;autoIncrement:true;comment:支付结算单据ID" json:"settlement_id"` // 支付结算单据ID
	PayNo        string    `gorm:"column:pay_no;comment:支付单号" json:"pay_no"`                                                 // 支付单号
	BizPayNo     string    `gorm:"column:biz_pay_no;comment:外部订单流水号" json:"biz_pay_no"`                                      // 外部订单流水号
	OrderNumber  string    `gorm:"column:order_number;comment:order表中的订单号" json:"order_number"`                              // order表中的订单号
	PayType      int32     `gorm:"column:pay_type;comment:支付方式 1 微信支付 2 支付宝" json:"pay_type"`                                // 支付方式 1 微信支付 2 支付宝
	PayTypeName  string    `gorm:"column:pay_type_name;comment:支付方式名称" json:"pay_type_name"`                                 // 支付方式名称
	PayAmount    float64   `gorm:"column:pay_amount;comment:支付金额" json:"pay_amount"`                                         // 支付金额
	IsClearing   int32     `gorm:"column:is_clearing;comment:是否清算 0:否 1:是" json:"is_clearing"`                               // 是否清算 0:否 1:是
	UserID       string    `gorm:"column:user_id;comment:用户ID" json:"user_id"`                                               // 用户ID
	CreateTime   time.Time `gorm:"column:create_time;not null;comment:创建时间" json:"create_time"`                              // 创建时间
	ClearingTime time.Time `gorm:"column:clearing_time;comment:清算时间" json:"clearing_time"`                                   // 清算时间
	Version      int32     `gorm:"column:version;comment:版本号" json:"version"`                                                // 版本号
	PayStatus    int32     `gorm:"column:pay_status;comment:支付状态" json:"pay_status"`                                         // 支付状态
}

// TableName OrderSettlement's table name
func (*OrderSettlement) TableName() string {
	return TableNameOrderSettlement
}
