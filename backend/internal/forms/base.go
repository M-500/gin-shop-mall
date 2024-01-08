package forms

type PaginateForm struct {
	PageSize int64 `json:"pageSize" form:"pageSize" binding:"-"`
	PageNum  int64 `json:"pageNum" form:"pageNum" binding:"-"`
}
