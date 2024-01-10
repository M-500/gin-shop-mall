package dto

import (
	"time"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/8 13:23
//

type PwdLoginDTO struct {
	AccessToken string `json:"accessToken"`
}

type UserInfoDTO struct {
	UserId    int64     `json:"userId"`
	CreatedAt time.Time `json:"createdTime"`
	UpdatedAt time.Time `json:"updatedTime"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	Status    int8      `json:"status"`
	ShopId    int64     `json:"shopId"`
}
