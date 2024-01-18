package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:51
//

type TzOrder struct {
	OrderId       uint64     `xorm:"not null pk autoincr comment('订单ID') UNSIGNED BIGINT" json:"orderId"`
	ShopId        int64      `xorm:"comment('店铺id') index BIGINT" json:"shopId"`
	ProdName      string     `xorm:"not null default '' comment('产品名称,多个产品将会以逗号隔开') VARCHAR(1000)" json:"prodName"`
	UserId        string     `xorm:"not null comment('订购用户ID') unique(order_number_userid_unique_ind) VARCHAR(36)" json:"userId"`
	OrderNumber   string     `xorm:"not null comment('订购流水号') unique unique(order_number_userid_unique_ind) VARCHAR(50)" json:"orderNumber"`
	Total         string     `xorm:"not null default 0.00 comment('总值') DECIMAL(15,2)" json:"total"`
	ActualTotal   string     `xorm:"comment('实际总值') DECIMAL(15,2)" json:"actualTotal"`
	PayType       int        `xorm:"comment('支付方式 0 手动代付 1 微信支付 2 支付宝') INT" json:"payType"`
	Remarks       string     `xorm:"comment('订单备注') VARCHAR(1024)" json:"remarks"`
	Status        int        `xorm:"not null default 0 comment('订单状态 1:待付款 2:待发货 3:待收货 4:待评价 5:成功 6:失败') INT" json:"status"`
	DvyType       string     `xorm:"comment('配送类型') VARCHAR(10)" json:"dvyType"`
	DvyId         int64      `xorm:"comment('配送方式ID') BIGINT" json:"dvyId"`
	DvyFlowId     string     `xorm:"default '' comment('物流单号') VARCHAR(100)" json:"dvyFlowId"`
	FreightAmount string     `xorm:"default 0.00 comment('订单运费') DECIMAL(15,2)" json:"freightAmount"`
	AddrOrderId   int64      `xorm:"comment('用户订单地址Id') BIGINT" json:"addrOrderId"`
	ProductNums   int        `xorm:"comment('订单商品总数') INT" json:"productNums"`
	CreateTime    *time.Time `xorm:"not null comment('订购时间') DATETIME" json:"createTime"`
	UpdateTime    *time.Time `xorm:"comment('订单更新时间') DATETIME" json:"updateTime"`
	PayTime       *time.Time `xorm:"comment('付款时间') DATETIME" json:"payTime"`
	DvyTime       *time.Time `xorm:"comment('发货时间') DATETIME" json:"dvyTime"`
	FinallyTime   *time.Time `xorm:"comment('完成时间') DATETIME" json:"finallyTime"`
	CancelTime    *time.Time `xorm:"comment('取消时间') DATETIME" json:"cancelTime"`
	IsPayed       int        `xorm:"comment('是否已经支付，1：已经支付过，0：，没有支付过') TINYINT(1)" json:"isPayed"`
	DeleteStatus  int        `xorm:"default 0 comment('用户订单删除状态，0：没有删除， 1：回收站， 2：永久删除') INT" json:"deleteStatus"`
	RefundSts     int        `xorm:"default 0 comment('0:默认,1:在处理,2:处理完成') INT" json:"refundSts"`
	ReduceAmount  string     `xorm:"comment('优惠总额') DECIMAL(15,2)" json:"reduceAmount"`
	OrderType     int        `xorm:"comment('订单类型') TINYINT" json:"orderType"`
	CloseType     int        `xorm:"comment('订单关闭原因 1-超时未支付 2-退款关闭 4-买家取消 15-已通过货到付款交易') TINYINT" json:"closeType"`
}
