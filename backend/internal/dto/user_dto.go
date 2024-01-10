package dto

import (
	"gorm.io/gorm"
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
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Email       string         `json:"email"`
	Mobile      string         `json:"mobile"`
	Status      int8           `json:"status"`
	CreatUserId int64          `json:"creatUserId"`
	ShopId      int64          `json:"shopId"`
}
