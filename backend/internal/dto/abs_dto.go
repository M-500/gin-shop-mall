package dto

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/11 14:57
//

type SysBaseDTO struct {
	Current int64
	Orders  []interface{}
	Pages   int64
	Records []interface{}
	Size    int64
	Total   int64
}
