// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameOrder = "tz_order"

// Order 订单表
type Order struct {
	OrderID       int64     `gorm:"column:order_id;primaryKey;autoIncrement:true;comment:订单ID" json:"order_id"`            // 订单ID
	ShopID        int64     `gorm:"column:shop_id;comment:店铺id" json:"shop_id"`                                            // 店铺id
	ProdName      string    `gorm:"column:prod_name;not null;comment:产品名称,多个产品将会以逗号隔开" json:"prod_name"`                   // 产品名称,多个产品将会以逗号隔开
	UserID        string    `gorm:"column:user_id;not null;comment:订购用户ID" json:"user_id"`                                 // 订购用户ID
	OrderNumber   string    `gorm:"column:order_number;not null;comment:订购流水号" json:"order_number"`                        // 订购流水号
	Total         float64   `gorm:"column:total;not null;default:0.00;comment:总值" json:"total"`                            // 总值
	ActualTotal   float64   `gorm:"column:actual_total;comment:实际总值" json:"actual_total"`                                  // 实际总值
	PayType       int32     `gorm:"column:pay_type;comment:支付方式 0 手动代付 1 微信支付 2 支付宝" json:"pay_type"`                      // 支付方式 0 手动代付 1 微信支付 2 支付宝
	Remarks       string    `gorm:"column:remarks;comment:订单备注" json:"remarks"`                                            // 订单备注
	Status        int32     `gorm:"column:status;not null;comment:订单状态 1:待付款 2:待发货 3:待收货 4:待评价 5:成功 6:失败" json:"status"`   // 订单状态 1:待付款 2:待发货 3:待收货 4:待评价 5:成功 6:失败
	DvyType       string    `gorm:"column:dvy_type;comment:配送类型" json:"dvy_type"`                                          // 配送类型
	DvyID         int64     `gorm:"column:dvy_id;comment:配送方式ID" json:"dvy_id"`                                            // 配送方式ID
	DvyFlowID     string    `gorm:"column:dvy_flow_id;comment:物流单号" json:"dvy_flow_id"`                                    // 物流单号
	FreightAmount float64   `gorm:"column:freight_amount;default:0.00;comment:订单运费" json:"freight_amount"`                 // 订单运费
	AddrOrderID   int64     `gorm:"column:addr_order_id;comment:用户订单地址Id" json:"addr_order_id"`                            // 用户订单地址Id
	ProductNums   int32     `gorm:"column:product_nums;comment:订单商品总数" json:"product_nums"`                                // 订单商品总数
	CreateTime    time.Time `gorm:"column:create_time;not null;comment:订购时间" json:"create_time"`                           // 订购时间
	UpdateTime    time.Time `gorm:"column:update_time;comment:订单更新时间" json:"update_time"`                                  // 订单更新时间
	PayTime       time.Time `gorm:"column:pay_time;comment:付款时间" json:"pay_time"`                                          // 付款时间
	DvyTime       time.Time `gorm:"column:dvy_time;comment:发货时间" json:"dvy_time"`                                          // 发货时间
	FinallyTime   time.Time `gorm:"column:finally_time;comment:完成时间" json:"finally_time"`                                  // 完成时间
	CancelTime    time.Time `gorm:"column:cancel_time;comment:取消时间" json:"cancel_time"`                                    // 取消时间
	IsPayed       bool      `gorm:"column:is_payed;comment:是否已经支付，1：已经支付过，0：，没有支付过" json:"is_payed"`                       // 是否已经支付，1：已经支付过，0：，没有支付过
	DeleteStatus  int32     `gorm:"column:delete_status;comment:用户订单删除状态，0：没有删除， 1：回收站， 2：永久删除" json:"delete_status"`      // 用户订单删除状态，0：没有删除， 1：回收站， 2：永久删除
	RefundSts     int32     `gorm:"column:refund_sts;comment:0:默认,1:在处理,2:处理完成" json:"refund_sts"`                         // 0:默认,1:在处理,2:处理完成
	ReduceAmount  float64   `gorm:"column:reduce_amount;comment:优惠总额" json:"reduce_amount"`                                // 优惠总额
	OrderType     int32     `gorm:"column:order_type;comment:订单类型" json:"order_type"`                                      // 订单类型
	CloseType     int32     `gorm:"column:close_type;comment:订单关闭原因 1-超时未支付 2-退款关闭 4-买家取消 15-已通过货到付款交易" json:"close_type"` // 订单关闭原因 1-超时未支付 2-退款关闭 4-买家取消 15-已通过货到付款交易
}

// TableName Order's table name
func (*Order) TableName() string {
	return TableNameOrder
}
