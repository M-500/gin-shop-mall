package prod_controller

import (
	"backend/internal/dto"
	"backend/internal/forms/cms_prod_form"
	"backend/pkg/response"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/9 21:13
//

/*
http://192.168.1.51:8085/prod/prodTag/page?t=1704805911303&current=1&size=10
获取分组列表
{
    "code": "00000",
    "msg": null,
    "data": {
        "records": [
            {
                "id": 1,
                "title": "每日上新",
                "shopId": 1,
                "status": 1,
                "isDefault": 0,
                "prodCount": 0,
                "seq": 3,
                "style": 2,
                "createTime": "2019-04-18 14:27:02",
                "updateTime": "2019-04-18 14:27:06",
                "deleteTime": null
            },
            {
                "id": 2,
                "title": "商城热卖",
                "shopId": 1,
                "status": 1,
                "isDefault": 0,
                "prodCount": 0,
                "seq": 2,
                "style": 1,
                "createTime": "2019-04-18 14:27:27",
                "updateTime": "2019-04-18 14:27:30",
                "deleteTime": null
            },
            {
                "id": 3,
                "title": "更多宝贝",
                "shopId": 1,
                "status": 1,
                "isDefault": 1,
                "prodCount": 0,
                "seq": 1,
                "style": 0,
                "createTime": "2019-04-18 10:07:17",
                "updateTime": "2019-04-18 10:07:17",
                "deleteTime": null
            }
        ],
        "total": 3,
        "size": 10,
        "current": 1,
        "orders": [],
        "pages": 1
    },
    "version": "mall4j.v230424",
    "timestamp": null,
    "sign": null,
    "success": true,
    "fail": false
}
*/

// 获取产品分类列表
func (p *ProductControllerGroup) AdminProdGroupTagList(c *gin.Context) {
	var queryForm = cms_prod_form.QueryProdTagForm{}
	err := c.ShouldBindJSON(&queryForm)
	if err != nil {

	}
	var resp dto.ProdTagDto
	list, i, err := p.prodTagService.QueryProdTagList(queryForm.Title, int(queryForm.Status), int(queryForm.Size), int(queryForm.Current))
	if err != nil {

	}
	resp.Total = i
	for _, item := range list {
		temp := dto.ProdTagRecord{
			ID:        item.ID,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			DeletedAt: item.DeletedAt,
			Title:     item.Title,
			ShopId:    item.ShopId,
			Status:    item.Status,
			IsDefault: item.IsDefault,
			ProdCount: item.ProdCount,
			Style:     item.Style,
			Seq:       item.Seq,
		}
		resp.Records = append(resp.Records, temp)
	}
	response.JsonSuccessData(c, "", resp)
}
