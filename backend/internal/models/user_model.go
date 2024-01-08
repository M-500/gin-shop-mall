package models

type SysUserModel struct {
	BaseModel
	Username    string `gorm:"size:type:;comment:用户账号;column:username" json:"username"`
	Password    string `gorm:"size:type:;comment:密码;column:password" json:"password"`
	Email       string `gorm:"size:type:;comment:邮箱;column:email" json:"email"`
	Mobile      string `gorm:"size:type:;comment:手机号;column:mobile" json:"mobile"`
	Status      int8   `gorm:"size:type:;comment:状态 0：禁用 1:正常;column:status" json:"status"`
	CreatUserId int64  `gorm:"size:type:;comment:创建人ID;column:creatUserId" json:"creatUserId"`
	ShopId      int64  `gorm:"size:type:;comment:用户所在的商场ID;column:shopId" json:"shopId"`
}

func (SysUserModel) TableName() string {
	return "m4_sys_user"
}
