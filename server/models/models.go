package models

import (
	"time"
)

type TzCategoryProp struct {
	Id         int64 `xorm:"pk autoincr BIGINT"`
	CategoryId int64 `xorm:"comment('分类id') index BIGINT"`
	PropId     int64 `xorm:"comment('商品属性id即表tz_prod_prop中的prop_id') index BIGINT"`
}

type TzProdFavorite struct {
	FavoriteId uint64    `xorm:"not null pk autoincr comment('主键') UNSIGNED BIGINT"`
	ProdId     uint64    `xorm:"not null comment('商品ID') index UNSIGNED BIGINT"`
	RecTime    time.Time `xorm:"not null comment('收藏时间') DATETIME"`
	UserId     string    `xorm:"not null default '' comment('用户ID') VARCHAR(36)"`
}

type TzProdProp struct {
	PropId   int64  `xorm:"not null pk autoincr comment('属性id') BIGINT"`
	PropName string `xorm:"comment('属性名称') VARCHAR(20)"`
	Rule     int    `xorm:"comment('ProdPropRule 1:销售属性(规格); 2:参数属性;') TINYINT"`
	ShopId   int64  `xorm:"comment('店铺id') index BIGINT"`
}

type TzProdPropValue struct {
	ValueId   int64  `xorm:"not null pk autoincr comment('属性值ID') BIGINT"`
	PropValue string `xorm:"comment('属性值名称') VARCHAR(20)"`
	PropId    int64  `xorm:"comment('属性ID') index BIGINT"`
}

type TzProdTag struct {
	Id         int64     `xorm:"pk autoincr comment('分组标签id') BIGINT"`
	Title      string    `xorm:"comment('分组标题') VARCHAR(36)"`
	ShopId     int64     `xorm:"comment('店铺Id') BIGINT"`
	Status     int       `xorm:"comment('状态(1为正常,0为删除)') TINYINT(1)"`
	IsDefault  int       `xorm:"comment('默认类型(0:商家自定义,1:系统默认)') TINYINT(1)"`
	ProdCount  int64     `xorm:"comment('商品数量') BIGINT"`
	Style      int       `xorm:"comment('列表样式(0:一列一个,1:一列两个,2:一列三个)') INT"`
	Seq        int       `xorm:"comment('排序') INT"`
	CreateTime time.Time `xorm:"comment('创建时间') TIMESTAMP"`
	UpdateTime time.Time `xorm:"comment('修改时间') TIMESTAMP"`
	DeleteTime time.Time `xorm:"comment('删除时间') TIMESTAMP"`
}

type TzProdTagReference struct {
	ReferenceId int64     `xorm:"not null pk autoincr comment('分组引用id') BIGINT"`
	ShopId      int64     `xorm:"comment('店铺id') BIGINT"`
	TagId       int64     `xorm:"comment('标签id') BIGINT"`
	ProdId      int64     `xorm:"comment('商品id') BIGINT"`
	Status      int       `xorm:"comment('状态(1:正常,0:删除)') TINYINT(1)"`
	CreateTime  time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
}

type TzShopDetail struct {
	ShopId           int64     `xorm:"not null pk autoincr comment('店铺id') BIGINT"`
	ShopName         string    `xorm:"comment('店铺名称(数字、中文，英文(可混合，不可有特殊字符)，可修改)、不唯一') VARCHAR(50)"`
	UserId           string    `xorm:"comment('店长用户id') unique VARCHAR(36)"`
	ShopType         int       `xorm:"comment('店铺类型') TINYINT"`
	Intro            string    `xorm:"comment('店铺简介(可修改)') VARCHAR(200)"`
	ShopNotice       string    `xorm:"comment('店铺公告(可修改)') VARCHAR(50)"`
	ShopIndustry     int       `xorm:"comment('店铺行业(餐饮、生鲜果蔬、鲜花等)') TINYINT"`
	ShopOwner        string    `xorm:"comment('店长') VARCHAR(20)"`
	Mobile           string    `xorm:"comment('店铺绑定的手机(登录账号：唯一)') unique VARCHAR(20)"`
	Tel              string    `xorm:"comment('店铺联系电话') VARCHAR(20)"`
	ShopLat          string    `xorm:"comment('店铺所在纬度(可修改)') VARCHAR(20)"`
	ShopLng          string    `xorm:"comment('店铺所在经度(可修改)') VARCHAR(20)"`
	ShopAddress      string    `xorm:"comment('店铺详细地址') VARCHAR(100)"`
	Province         string    `xorm:"comment('店铺所在省份（描述）') VARCHAR(10)"`
	City             string    `xorm:"comment('店铺所在城市（描述）') VARCHAR(10)"`
	Area             string    `xorm:"comment('店铺所在区域（描述）') VARCHAR(10)"`
	PcaCode          string    `xorm:"comment('店铺省市区代码，用于回显') VARCHAR(20)"`
	ShopLogo         string    `xorm:"comment('店铺logo(可修改)') VARCHAR(200)"`
	ShopPhotos       string    `xorm:"comment('店铺相册') VARCHAR(1000)"`
	OpenTime         string    `xorm:"comment('每天营业时间段(可修改)') VARCHAR(100)"`
	ShopStatus       int       `xorm:"comment('店铺状态(-1:未开通 0: 停业中 1:营业中)，可修改') TINYINT"`
	TransportType    int       `xorm:"comment('0:商家承担运费; 1:买家承担运费') TINYINT"`
	FixedFreight     string    `xorm:"comment('固定运费') DECIMAL(15,2)"`
	FullFreeShipping string    `xorm:"comment('满X包邮') DECIMAL(15,2)"`
	CreateTime       time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateTime       time.Time `xorm:"comment('更新时间') DATETIME"`
	IsDistribution   int       `xorm:"comment('分销开关(0:开启 1:关闭)') TINYINT"`
}

type TzSku struct {
	SkuId        uint64    `xorm:"not null pk autoincr comment('单品ID') UNSIGNED BIGINT"`
	ProdId       uint64    `xorm:"not null comment('商品ID') index UNSIGNED BIGINT"`
	Properties   string    `xorm:"default '' comment('销售属性组合字符串 格式是p1:v1;p2:v2') VARCHAR(2000)"`
	OriPrice     string    `xorm:"comment('原价') DECIMAL(15,2)"`
	Price        string    `xorm:"comment('价格') DECIMAL(15,2)"`
	Stocks       int       `xorm:"not null comment('商品在付款减库存的状态下，该sku上未付款的订单数量') INT"`
	ActualStocks int       `xorm:"comment('实际库存') INT"`
	UpdateTime   time.Time `xorm:"not null comment('修改时间') DATETIME"`
	RecTime      time.Time `xorm:"not null comment('记录时间') DATETIME"`
	PartyCode    string    `xorm:"comment('商家编码') VARCHAR(100)"`
	ModelId      string    `xorm:"comment('商品条形码') VARCHAR(100)"`
	Pic          string    `xorm:"comment('sku图片') VARCHAR(500)"`
	SkuName      string    `xorm:"comment('sku名称') VARCHAR(120)"`
	ProdName     string    `xorm:"comment('商品名称') VARCHAR(255)"`
	Version      int       `xorm:"not null default 0 comment('版本号') INT"`
	Weight       float64   `xorm:"comment('商品重量') DOUBLE"`
	Volume       float64   `xorm:"comment('商品体积') DOUBLE"`
	Status       int       `xorm:"default 1 comment('0 禁用 1 启用') TINYINT"`
	IsDelete     int       `xorm:"comment('0 正常 1 已被删除') TINYINT"`
}

type TzSmsLog struct {
	Id           uint64    `xorm:"not null pk autoincr comment('ID') UNSIGNED BIGINT"`
	UserId       string    `xorm:"comment('用户id') VARCHAR(50)"`
	UserPhone    string    `xorm:"not null default '' comment('手机号码') VARCHAR(20)"`
	Content      string    `xorm:"not null default '' comment('短信内容') VARCHAR(100)"`
	MobileCode   string    `xorm:"not null default '' comment('手机验证码') VARCHAR(50)"`
	Type         int       `xorm:"not null default 0 comment('短信类型  1:注册  2:验证') INT"`
	RecDate      time.Time `xorm:"not null comment('发送时间') DATETIME"`
	ResponseCode string    `xorm:"comment('发送短信返回码') VARCHAR(50)"`
	Status       int       `xorm:"not null default 0 comment('状态  1:有效  0：失效') INT"`
}

type TzSysConfig struct {
	Id         int64  `xorm:"pk autoincr BIGINT"`
	ParamKey   string `xorm:"comment('key') unique VARCHAR(50)"`
	ParamValue string `xorm:"comment('value') VARCHAR(2000)"`
	Remark     string `xorm:"comment('备注') VARCHAR(500)"`
}

type TzSysLog struct {
	Id         int64     `xorm:"pk autoincr BIGINT"`
	Username   string    `xorm:"comment('用户名') VARCHAR(50)"`
	Operation  string    `xorm:"comment('用户操作') VARCHAR(50)"`
	Method     string    `xorm:"comment('请求方法') VARCHAR(200)"`
	Params     string    `xorm:"comment('请求参数') VARCHAR(5000)"`
	Time       int64     `xorm:"not null comment('执行时长(毫秒)') BIGINT"`
	Ip         string    `xorm:"comment('IP地址') VARCHAR(64)"`
	CreateDate time.Time `xorm:"comment('创建时间') DATETIME"`
}

type TzSysMenu struct {
	MenuId   int64  `xorm:"not null pk autoincr BIGINT"`
	ParentId int64  `xorm:"comment('父菜单ID，一级菜单为0') BIGINT"`
	Name     string `xorm:"comment('菜单名称') VARCHAR(50)"`
	Url      string `xorm:"comment('菜单URL') VARCHAR(200)"`
	Perms    string `xorm:"comment('授权(多个用逗号分隔，如：user:list,user:create)') VARCHAR(500)"`
	Type     int    `xorm:"comment('类型   0：目录   1：菜单   2：按钮') INT"`
	Icon     string `xorm:"comment('菜单图标') VARCHAR(50)"`
	OrderNum int    `xorm:"comment('排序') INT"`
}

type TzSysRole struct {
	RoleId       int64     `xorm:"not null pk autoincr BIGINT"`
	RoleName     string    `xorm:"comment('角色名称') VARCHAR(100)"`
	Remark       string    `xorm:"comment('备注') VARCHAR(100)"`
	CreateUserId int64     `xorm:"comment('创建者ID') BIGINT"`
	CreateTime   time.Time `xorm:"comment('创建时间') DATETIME"`
}

type TzSysRoleMenu struct {
	Id     int64 `xorm:"pk autoincr BIGINT"`
	RoleId int64 `xorm:"comment('角色ID') BIGINT"`
	MenuId int64 `xorm:"comment('菜单ID') BIGINT"`
}

type TzSysUser struct {
	UserId       int64     `xorm:"not null pk autoincr BIGINT"`
	Username     string    `xorm:"not null comment('用户名') unique VARCHAR(50)"`
	Password     string    `xorm:"comment('密码') VARCHAR(100)"`
	Email        string    `xorm:"comment('邮箱') VARCHAR(100)"`
	Mobile       string    `xorm:"comment('手机号') VARCHAR(100)"`
	Status       int       `xorm:"comment('状态  0：禁用   1：正常') TINYINT"`
	CreateUserId int64     `xorm:"comment('创建者ID') BIGINT"`
	CreateTime   time.Time `xorm:"comment('创建时间') DATETIME"`
	ShopId       int64     `xorm:"comment('用户所在的商城Id') BIGINT"`
}

type TzSysUserRole struct {
	Id     int64 `xorm:"pk autoincr BIGINT"`
	UserId int64 `xorm:"comment('用户ID') BIGINT"`
	RoleId int64 `xorm:"comment('角色ID') BIGINT"`
}

type TzTranscity struct {
	TranscityId int64 `xorm:"not null pk autoincr BIGINT"`
	TransfeeId  int64 `xorm:"comment('运费项id') index BIGINT"`
	CityId      int64 `xorm:"comment('城市id') index BIGINT"`
}

type TzTranscityFree struct {
	TranscityFreeId int64 `xorm:"not null pk autoincr comment('指定条件包邮城市项id') BIGINT"`
	TransfeeFreeId  int64 `xorm:"comment('指定条件包邮项id') index BIGINT"`
	FreeCityId      int64 `xorm:"comment('城市id') index BIGINT"`
}

type TzTransfee struct {
	TransfeeId      int64  `xorm:"not null pk autoincr comment('运费项id') BIGINT"`
	TransportId     int64  `xorm:"comment('运费模板id') index BIGINT"`
	ContinuousPiece string `xorm:"comment('续件数量') DECIMAL(15,2)"`
	FirstPiece      string `xorm:"comment('首件数量') DECIMAL(15,2)"`
	ContinuousFee   string `xorm:"comment('续件费用') DECIMAL(15,2)"`
	FirstFee        string `xorm:"comment('首件费用') DECIMAL(15,2)"`
}

type TzTransfeeFree struct {
	TransfeeFreeId int64  `xorm:"not null pk autoincr comment('指定条件包邮项id') BIGINT"`
	TransportId    int64  `xorm:"comment('运费模板id') index BIGINT"`
	FreeType       int    `xorm:"comment('包邮方式 （0 满x件/重量/体积包邮 1满金额包邮 2满x件/重量/体积且满金额包邮）') TINYINT"`
	Amount         string `xorm:"comment('需满金额') DECIMAL(15,2)"`
	Piece          string `xorm:"comment('包邮x件/重量/体积') DECIMAL(15,2)"`
}

type TzTransport struct {
	TransportId      int64     `xorm:"not null pk autoincr comment('运费模板id') BIGINT"`
	TransName        string    `xorm:"comment('运费模板名称') VARCHAR(36)"`
	CreateTime       time.Time `xorm:"comment('创建时间') DATETIME"`
	ShopId           int64     `xorm:"comment('店铺id') index BIGINT"`
	ChargeType       int       `xorm:"comment('收费方式（0 按件数,1 按重量 2 按体积）') TINYINT"`
	IsFreeFee        int       `xorm:"comment('是否包邮 0:不包邮 1:包邮') TINYINT"`
	HasFreeCondition int       `xorm:"comment('是否含有包邮条件 0 否 1是') TINYINT"`
}

type TzUser struct {
	UserId        string    `xorm:"not null pk default '' comment('ID') VARCHAR(36)"`
	NickName      string    `xorm:"comment('用户昵称') VARCHAR(50)"`
	RealName      string    `xorm:"comment('真实姓名') VARCHAR(50)"`
	UserMail      string    `xorm:"comment('用户邮箱') unique VARCHAR(100)"`
	LoginPassword string    `xorm:"comment('登录密码') VARCHAR(255)"`
	PayPassword   string    `xorm:"comment('支付密码') VARCHAR(50)"`
	UserMobile    string    `xorm:"comment('手机号码') unique VARCHAR(50)"`
	ModifyTime    time.Time `xorm:"not null comment('修改时间') DATETIME"`
	UserRegtime   time.Time `xorm:"not null comment('注册时间') DATETIME"`
	UserRegip     string    `xorm:"comment('注册IP') VARCHAR(50)"`
	UserLasttime  time.Time `xorm:"comment('最后登录时间') DATETIME"`
	UserLastip    string    `xorm:"comment('最后登录IP') VARCHAR(50)"`
	UserMemo      string    `xorm:"comment('备注') VARCHAR(500)"`
	Sex           string    `xorm:"default 'M' comment('M(男) or F(女)') CHAR(1)"`
	BirthDate     string    `xorm:"comment('例如：2009-11-27') CHAR(10)"`
	Pic           string    `xorm:"comment('头像图片路径') VARCHAR(255)"`
	Status        int       `xorm:"not null default 1 comment('状态 1 正常 0 无效') INT"`
	Score         int       `xorm:"comment('用户积分') INT"`
}

type TzUserAddr struct {
	AddrId     uint64    `xorm:"not null pk autoincr comment('ID') UNSIGNED BIGINT"`
	UserId     string    `xorm:"not null default '0' comment('用户ID') VARCHAR(36)"`
	Receiver   string    `xorm:"comment('收货人') VARCHAR(50)"`
	ProvinceId int64     `xorm:"comment('省ID') BIGINT"`
	Province   string    `xorm:"comment('省') VARCHAR(100)"`
	City       string    `xorm:"comment('城市') VARCHAR(20)"`
	CityId     int64     `xorm:"comment('城市ID') BIGINT"`
	Area       string    `xorm:"comment('区') VARCHAR(20)"`
	AreaId     int64     `xorm:"comment('区ID') BIGINT"`
	PostCode   string    `xorm:"comment('邮编') VARCHAR(15)"`
	Addr       string    `xorm:"comment('地址') VARCHAR(1000)"`
	Mobile     string    `xorm:"comment('手机') VARCHAR(20)"`
	Status     int       `xorm:"not null comment('状态,1正常，0无效') INT"`
	CommonAddr int       `xorm:"not null default 0 comment('是否默认地址 1是') INT"`
	CreateTime time.Time `xorm:"not null comment('建立时间') DATETIME"`
	Version    int       `xorm:"not null default 0 comment('版本号') INT"`
	UpdateTime time.Time `xorm:"not null comment('更新时间') DATETIME"`
}

type TzUserAddrOrder struct {
	AddrOrderId uint64    `xorm:"not null pk autoincr comment('ID') UNSIGNED BIGINT"`
	AddrId      uint64    `xorm:"not null comment('地址ID') UNSIGNED BIGINT"`
	UserId      string    `xorm:"not null default '0' comment('用户ID') VARCHAR(36)"`
	Receiver    string    `xorm:"comment('收货人') VARCHAR(50)"`
	ProvinceId  int64     `xorm:"comment('省ID') BIGINT"`
	Province    string    `xorm:"comment('省') VARCHAR(100)"`
	AreaId      int64     `xorm:"comment('区域ID') BIGINT"`
	Area        string    `xorm:"comment('区') VARCHAR(20)"`
	CityId      int64     `xorm:"comment('城市ID') BIGINT"`
	City        string    `xorm:"comment('城市') VARCHAR(20)"`
	Addr        string    `xorm:"comment('地址') VARCHAR(1000)"`
	PostCode    string    `xorm:"comment('邮编') VARCHAR(15)"`
	Mobile      string    `xorm:"comment('手机') VARCHAR(20)"`
	CreateTime  time.Time `xorm:"not null comment('建立时间') DATETIME"`
	Version     int       `xorm:"not null default 0 comment('版本号') INT"`
}

type TzUserCollection struct {
	Id         int64     `xorm:"pk autoincr comment('收藏表') BIGINT"`
	ProdId     int64     `xorm:"comment('商品id') BIGINT"`
	UserId     string    `xorm:"not null comment('用户id') VARCHAR(36)"`
	CreateTime time.Time `xorm:"comment('收藏时间') DATETIME"`
}
