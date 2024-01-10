package cms_sys_form

type AdminLoginParam struct {
	UserName    string `json:"userName"`
	PasswordMd5 string `json:"password"`
}

type CreateAdminUserForm struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=30"`
	Mobile   string `form:"mobile" json:"mobile"`
}
