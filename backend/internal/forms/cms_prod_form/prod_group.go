package cms_prod_form

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/10 9:13
//

type QueryProdTagForm struct {
	Title   string `json:"title"`
	Current int64  `json:"current"`
	Size    int64  `json:"size"`
	T       int64  `json:"t"`
	Status  int64  `json:"status"`
}
