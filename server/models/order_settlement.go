package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:53
//

type TzOrderSettlement struct {
	SettlementId uint64     `xorm:"not null pk autoincr comment('支付结算单据ID') UNSIGNED BIGINT" json:"settlementId"`
	PayNo        string     `xorm:"comment('支付单号') VARCHAR(36)" json:"payNo"`
	BizPayNo     string     `xorm:"comment('外部订单流水号') VARCHAR(255)" json:"bizPayNo"`
	OrderNumber  string     `xorm:"comment('order表中的订单号') unique VARCHAR(36)" json:"orderNumber"`
	PayType      int        `xorm:"comment('支付方式 1 微信支付 2 支付宝') INT" json:"payType"`
	PayTypeName  string     `xorm:"comment('支付方式名称') VARCHAR(50)" json:"payTypeName"`
	PayAmount    string     `xorm:"comment('支付金额') DECIMAL(15,2)" json:"payAmount"`
	IsClearing   int        `xorm:"comment('是否清算 0:否 1:是') INT" json:"isClearing"`
	UserId       string     `xorm:"comment('用户ID') index VARCHAR(36)" json:"userId"`
	CreateTime   *time.Time `xorm:"not null comment('创建时间') DATETIME" json:"createTime"`
	ClearingTime *time.Time `xorm:"comment('清算时间') DATETIME" json:"clearingTime"`
	Version      int        `xorm:"comment('版本号') INT" json:"version"`
	PayStatus    int        `xorm:"comment('支付状态') INT" json:"payStatus"`
}
