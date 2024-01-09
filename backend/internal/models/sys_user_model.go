package models

import "time"

// 后台管理用户表
type SysUserModel struct {
	BaseModel
	Username    string `gorm:"uniqueIndex;type:varchar(50);comment:用户账号;column:username" json:"username"`
	Password    string `gorm:"type:varchar(128);comment:密码;column:password" json:"password"`
	Email       string `gorm:"type:varchar(128);comment:邮箱;column:email" json:"email"`
	Mobile      string `gorm:"type:varchar(11);comment:手机号;column:mobile" json:"mobile"`
	Status      int8   `gorm:"type:tinyint;comment:状态 0：禁用 1:正常;column:status" json:"status"`
	CreatUserId int64  `gorm:"type:bigint;comment:创建人ID;column:creatUserId" json:"creatUserId"`
	ShopId      int64  `gorm:"type:bigint;comment:用户所在的商场ID;column:shopId" json:"shopId"`
}

func (SysUserModel) TableName() string {
	return "m4_sys_user"
}

// 前台用户表
type UserModel struct {
	BaseModel
	UserId       string     `gorm:"uniqueIndex;type:varchar(32);comment:用户ID;column:user_id" json:"userId"`
	NickName     string     `gorm:"type:varchar(50);comment:昵称;column:nick_name" json:"nickName"`
	RealName     string     `gorm:"type:varchar(50);comment:真实姓名;column:real_name" json:"realName"`
	UserEmail    string     `gorm:"type:varchar(128);comment:邮箱;column:user_email" json:"userEmail"`
	LoginPWD     string     `gorm:"type: varchar(255);comment:登录密码;column:login_pwd" json:"loginPWD"`
	PayPwd       string     `gorm:"type:varchar(50);comment:支付密码;column:pay_pwd" json:"payPwd"`
	UserMobile   string     `gorm:"type:varchar(11);comment:手机号;column:user_mobile" json:"userMobile"`
	ModifyTime   *time.Time `gorm:"type:datetime;comment:修改时间;column:modify_time" json:"modifyTime"`
	UserRegIP    string     `gorm:"type:varchar(32);default:'NULL';comment:注册IP;column:user_reg_ip" json:"userRegIP"`
	UserLastTime *time.Time `gorm:"type:datetime;comment:最后登录时间;column:user_last_time" json:"userLastTime"`
	UserMemo     string     `gorm:"type:varchar(500);comment:用户备注;column:user_memo" json:"userMemo"`
	Sex          string     `gorm:"type:char(1);comment:性别M(男) or F(女);column:sex" json:"sex"`
	BirthDay     string     `gorm:"type:char(10);comment:例如：2009-11-27;column:birth_day" json:"birthDay"`
	CoverUrl     string     `gorm:"type:varchar(255);comment:头像图片路径;column:cover_url" json:"coverUrl"`
	Score        int        `gorm:"type:int;comment:用户积分;column:score" json:"score"`
	Status       int        `gorm:"type:int;comment:状态 1 正常 0 无效;column:status" json:"status"`
}

func (UserModel) TableName() string {
	return "m4_user"
}
