//@Author: wulinlin
//@Description:
//@File:  delivery_model
//@Version: 1.0.0
//@Date: 2024/01/08 20:24

package models

import "time"

// 物流公司表
type DeliveryModel struct {
	BaseModel
	DvyName        string
	CompanyHomeURL string
	RecTime        *time.Time
	QueryUrl       string
}
