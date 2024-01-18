package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:52
//

type TzOrderRefund struct {
	RefundId       int64      `xorm:"not null pk autoincr comment('记录ID') BIGINT" json:"refundId"`
	ShopId         int64      `xorm:"not null comment('店铺ID') BIGINT" json:"shopId"`
	OrderId        int64      `xorm:"not null comment('订单ID') BIGINT" json:"orderId"`
	OrderNumber    string     `xorm:"not null comment('订单流水号') index VARCHAR(50)" json:"orderNumber"`
	OrderAmount    float64    `xorm:"not null comment('订单总金额') DOUBLE(12,2)" json:"orderAmount"`
	OrderItemId    int64      `xorm:"not null default 0 comment('订单项ID 全部退款是0') BIGINT" json:"orderItemId"`
	RefundSn       string     `xorm:"not null comment('退款编号') unique VARCHAR(50)" json:"refundSn"`
	FlowTradeNo    string     `xorm:"not null comment('订单支付流水号') VARCHAR(100)" json:"flowTradeNo"`
	OutRefundNo    string     `xorm:"comment('第三方退款单号(微信退款单号)') VARCHAR(200)" json:"outRefundNo"`
	PayType        int        `xorm:"comment('订单支付方式 1 微信支付 2 支付宝') INT" json:"payType"`
	PayTypeName    string     `xorm:"comment('订单支付名称') VARCHAR(50)" json:"payTypeName"`
	UserId         string     `xorm:"not null comment('买家ID') VARCHAR(50)" json:"userId"`
	GoodsNum       int        `xorm:"comment('退货数量') INT" json:"goodsNum"`
	RefundAmount   string     `xorm:"not null comment('退款金额') DECIMAL(10,2)" json:"refundAmount"`
	ApplyType      int        `xorm:"not null default 0 comment('申请类型:1,仅退款,2退款退货') INT" json:"applyType"`
	RefundSts      int        `xorm:"not null default 0 comment('处理状态:1为待审核,2为同意,3为不同意') INT" json:"refundSts"`
	ReturnMoneySts int        `xorm:"not null default 0 comment('处理退款状态: 0:退款处理中 1:退款成功 -1:退款失败') INT" json:"returnMoneySts"`
	ApplyTime      *time.Time `xorm:"not null comment('申请时间') DATETIME" json:"applyTime"`
	HandelTime     *time.Time `xorm:"comment('卖家处理时间') DATETIME" json:"handelTime"`
	RefundTime     *time.Time `xorm:"comment('退款时间') DATETIME" json:"refundTime"`
	PhotoFiles     string     `xorm:"comment('文件凭证json') VARCHAR(150)" json:"photoFiles"`
	BuyerMsg       string     `xorm:"comment('申请原因') VARCHAR(300)" json:"buyerMsg"`
	SellerMsg      string     `xorm:"comment('卖家备注') VARCHAR(300)" json:"sellerMsg"`
	ExpressName    string     `xorm:"comment('物流公司名称') VARCHAR(50)" json:"expressName"`
	ExpressNo      string     `xorm:"comment('物流单号') VARCHAR(50)" json:"expressNo"`
	ShipTime       *time.Time `xorm:"comment('发货时间') DATETIME" json:"shipTime"`
	ReceiveTime    *time.Time `xorm:"comment('收货时间') DATETIME" json:"receiveTime"`
	ReceiveMessage string     `xorm:"comment('收货备注') VARCHAR(300)" json:"receiveMessage"`
}
