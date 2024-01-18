// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameShopDetail = "tz_shop_detail"

// ShopDetail mapped from table <tz_shop_detail>
type ShopDetail struct {
	ShopID           int64     `gorm:"column:shop_id;primaryKey;autoIncrement:true;comment:店铺id" json:"shop_id"`      // 店铺id
	ShopName         string    `gorm:"column:shop_name;comment:店铺名称(数字、中文，英文(可混合，不可有特殊字符)，可修改)、不唯一" json:"shop_name"` // 店铺名称(数字、中文，英文(可混合，不可有特殊字符)，可修改)、不唯一
	UserID           string    `gorm:"column:user_id;comment:店长用户id" json:"user_id"`                                  // 店长用户id
	ShopType         int32     `gorm:"column:shop_type;comment:店铺类型" json:"shop_type"`                                // 店铺类型
	Intro            string    `gorm:"column:intro;comment:店铺简介(可修改)" json:"intro"`                                   // 店铺简介(可修改)
	ShopNotice       string    `gorm:"column:shop_notice;comment:店铺公告(可修改)" json:"shop_notice"`                       // 店铺公告(可修改)
	ShopIndustry     int32     `gorm:"column:shop_industry;comment:店铺行业(餐饮、生鲜果蔬、鲜花等)" json:"shop_industry"`           // 店铺行业(餐饮、生鲜果蔬、鲜花等)
	ShopOwner        string    `gorm:"column:shop_owner;comment:店长" json:"shop_owner"`                                // 店长
	Mobile           string    `gorm:"column:mobile;comment:店铺绑定的手机(登录账号：唯一)" json:"mobile"`                          // 店铺绑定的手机(登录账号：唯一)
	Tel              string    `gorm:"column:tel;comment:店铺联系电话" json:"tel"`                                          // 店铺联系电话
	ShopLat          string    `gorm:"column:shop_lat;comment:店铺所在纬度(可修改)" json:"shop_lat"`                           // 店铺所在纬度(可修改)
	ShopLng          string    `gorm:"column:shop_lng;comment:店铺所在经度(可修改)" json:"shop_lng"`                           // 店铺所在经度(可修改)
	ShopAddress      string    `gorm:"column:shop_address;comment:店铺详细地址" json:"shop_address"`                        // 店铺详细地址
	Province         string    `gorm:"column:province;comment:店铺所在省份（描述）" json:"province"`                            // 店铺所在省份（描述）
	City             string    `gorm:"column:city;comment:店铺所在城市（描述）" json:"city"`                                    // 店铺所在城市（描述）
	Area             string    `gorm:"column:area;comment:店铺所在区域（描述）" json:"area"`                                    // 店铺所在区域（描述）
	PcaCode          string    `gorm:"column:pca_code;comment:店铺省市区代码，用于回显" json:"pca_code"`                          // 店铺省市区代码，用于回显
	ShopLogo         string    `gorm:"column:shop_logo;comment:店铺logo(可修改)" json:"shop_logo"`                         // 店铺logo(可修改)
	ShopPhotos       string    `gorm:"column:shop_photos;comment:店铺相册" json:"shop_photos"`                            // 店铺相册
	OpenTime         string    `gorm:"column:open_time;comment:每天营业时间段(可修改)" json:"open_time"`                        // 每天营业时间段(可修改)
	ShopStatus       int32     `gorm:"column:shop_status;comment:店铺状态(-1:未开通 0: 停业中 1:营业中)，可修改" json:"shop_status"`   // 店铺状态(-1:未开通 0: 停业中 1:营业中)，可修改
	TransportType    int32     `gorm:"column:transport_type;comment:0:商家承担运费; 1:买家承担运费" json:"transport_type"`        // 0:商家承担运费; 1:买家承担运费
	FixedFreight     float64   `gorm:"column:fixed_freight;comment:固定运费" json:"fixed_freight"`                        // 固定运费
	FullFreeShipping float64   `gorm:"column:full_free_shipping;comment:满X包邮" json:"full_free_shipping"`              // 满X包邮
	CreateTime       time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`                            // 创建时间
	UpdateTime       time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                            // 更新时间
	IsDistribution   int32     `gorm:"column:is_distribution;comment:分销开关(0:开启 1:关闭)" json:"is_distribution"`         // 分销开关(0:开启 1:关闭)
}

// TableName ShopDetail's table name
func (*ShopDetail) TableName() string {
	return TableNameShopDetail
}