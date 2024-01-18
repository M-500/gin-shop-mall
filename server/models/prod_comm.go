package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:55
//

type TzProdComm struct {
	ProdCommId   uint64     `xorm:"not null pk autoincr comment('ID') UNSIGNED BIGINT" json:"prodCommId"`
	ProdId       uint64     `xorm:"not null comment('商品ID') index UNSIGNED BIGINT" json:"prodId"`
	OrderItemId  uint64     `xorm:"comment('订单项ID') UNSIGNED BIGINT" json:"orderItemId"`
	UserId       string     `xorm:"comment('评论用户ID') VARCHAR(36)" json:"userId"`
	Content      string     `xorm:"default '' comment('评论内容') VARCHAR(500)" json:"content"`
	ReplyContent string     `xorm:"default '' comment('掌柜回复') VARCHAR(500)" json:"replyContent"`
	RecTime      *time.Time `xorm:"comment('记录时间') DATETIME" json:"recTime"`
	ReplyTime    *time.Time `xorm:"comment('回复时间') DATETIME" json:"replyTime"`
	ReplySts     int        `xorm:"default 0 comment('是否回复 0:未回复  1:已回复') INT" json:"replySts"`
	Postip       string     `xorm:"comment('IP来源') VARCHAR(16)" json:"postip"`
	Score        int        `xorm:"default 0 comment('得分，0-5分') TINYINT" json:"score"`
	UsefulCounts int        `xorm:"default 0 comment('有用的计数') INT" json:"usefulCounts"`
	Pics         string     `xorm:"comment('晒图的json字符串') VARCHAR(1000)" json:"pics"`
	IsAnonymous  int        `xorm:"default 0 comment('是否匿名(1:是  0:否)') INT" json:"isAnonymous"`
	Status       int        `xorm:"comment('是否显示，1:为显示，0:待审核， -1：不通过审核，不显示。 如果需要审核评论，则是0,，否则1') INT" json:"status"`
	Evaluate     int        `xorm:"comment('评价(0好评 1中评 2差评)') TINYINT" json:"evaluate"`
}
