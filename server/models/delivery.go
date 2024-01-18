package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:48
//

type TzDelivery struct {
	DvyId          uint64     `xorm:"not null pk autoincr comment('ID') UNSIGNED BIGINT" json:"dvyId"`
	DvyName        string     `xorm:"not null default '' comment('物流公司名称') VARCHAR(50)" json:"dvyName"`
	CompanyHomeUrl string     `xorm:"comment('公司主页') VARCHAR(255)" json:"companyHomeUrl"`
	RecTime        *time.Time `xorm:"not null comment('建立时间') DATETIME" json:"recTime"`
	ModifyTime     *time.Time `xorm:"not null comment('修改时间') DATETIME" json:"modifyTime"`
	QueryUrl       string     `xorm:"not null comment('物流查询接口') VARCHAR(520)" json:"queryUrl"`
}
