package models

import "time"

// OrderModel 订单表
type OrderModel struct {
	BaseModel
	ShopId        int64      `gorm:"size:;type:;comment:;column:shop_id" json:"shop_id"`
	ProductName   string     `gorm:"size:;type:;comment:;column:product_name" json:"product_name"`
	UserId        int64      `gorm:"size:;type:;comment:;column:user_id" json:"user_id"`
	OrderNum      string     `gorm:"size:;type:;comment:;column:order_num" json:"order_num"`
	Total         float64    `gorm:"size:;type:;comment:;column:total" json:"total"`
	ActualTotal   float64    `gorm:"size:;type:;comment:;column:actual_total" json:"actual_total"`
	PayType       int8       `gorm:"size:;type:;comment:;column:pay_type" json:"pay_type"`
	Remarks       string     `gorm:"size:;type:;comment:;column:remarks" json:"remarks"`
	Status        int8       `gorm:"size:;type:;comment:;column:status" json:"status"`
	DvyType       string     `gorm:"size:;type:;comment:;column:dvy_type" json:"dvy_type"`
	DvyId         int64      `gorm:"size:;type:;comment:;column:dvy_id" json:"dvy_id"`
	DvyFlowId     string     `gorm:"size:;type:;comment:;column:dvy_flow_id" json:"dvy_flow_id"`
	FreightAmount float64    `gorm:"size:;type:;comment:;column:freight_amount" json:"freight_amount"`
	AddrOrderId   string     `gorm:"size:;type:;comment:;column:addr_order_id" json:"addr_order_id"`
	ProductNuns   int        `gorm:"size:;type:;comment:;column:product_nuns" json:"product_nuns"`
	PayTime       *time.Time `gorm:"size:;type:;comment:;column:pay_time" json:"pay_time"`
	DvyTime       *time.Time `gorm:"size:;type:;comment:;column:dvy_time" json:"dvy_time"`
	FinallyTime   *time.Time `gorm:"size:;type:;comment:;column:finally_time" json:"finally_time"`
	CancelTime    *time.Time `gorm:"size:;type:;comment:;column:cancel_time" json:"cancel_time"`
	IsPayed       int8       `gorm:"size:;type:;comment:;column:is_payed" json:"is_payed"`
	DeleteStatus  int8       `gorm:"size:;type:;comment:;column:delete_status" json:"delete_status"`
	RefundSts     int8       `gorm:"size:;type:;comment:;column:refund_sts" json:"refund_sts"`
	ReduceAmount  float64    `gorm:"size:;type:;comment:;column:reduce_amount" json:"reduce_amount"`
	OrderType     int8       `gorm:"size:;type:;comment:;column:order_type" json:"order_type"`
	CloseType     int8       `gorm:"size:;type:;comment:;column:close_type" json:"close_type"`
}

func (OrderModel) TableName() string {
	return ""
}
