package models

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:54
//

type TzPickAddr struct {
	AddrId     uint64 `xorm:"not null pk autoincr comment('ID') UNSIGNED BIGINT" json:"addrId"`
	AddrName   string `xorm:"comment('自提点名称') VARCHAR(36)" json:"addrName"`
	Addr       string `xorm:"comment('地址') VARCHAR(1000)" json:"addr"`
	Mobile     string `xorm:"comment('手机') VARCHAR(20)" json:"mobile"`
	ProvinceId int64  `xorm:"comment('省份ID') BIGINT" json:"provinceId"`
	Province   string `xorm:"comment('省份') VARCHAR(32)" json:"province"`
	CityId     int64  `xorm:"comment('城市ID') BIGINT" json:"cityId"`
	City       string `xorm:"comment('城市') VARCHAR(32)" json:"city"`
	AreaId     int64  `xorm:"comment('区/县ID') BIGINT" json:"areaId"`
	Area       string `xorm:"comment('区/县') VARCHAR(32)" json:"area"`
	ShopId     int64  `xorm:"comment('店铺id') BIGINT" json:"shopId"`
}
