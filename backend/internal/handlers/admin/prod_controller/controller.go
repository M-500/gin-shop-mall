package prod_controller

import "backend/internal/service"

/*
包含了产品管理的全部路由
*/
type ProductController struct {
	prodTagService service.ProdService
}

func NewProductController() *ProductController {
	return &ProductController{}
}
