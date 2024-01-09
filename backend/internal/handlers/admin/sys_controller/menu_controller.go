package sys_controller

import (
	"backend/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/9 18:54
//

// 获取用户所拥有的菜单和权限 通过登陆用户的userId获取用户所拥有的菜单和权限
func (u *SysController) MenuNavHandler(c *gin.Context) {
	// 如果是系统管理员 拥有最高权限

	// 组装一级菜单和二级菜单
	//id, err := u.menuService.ListMenuByUserId(c.GetInt64(constant.JWT_INFO_UID))
	id, err := u.menuService.ListMenuByUserId(1)
	if err != nil {
		return
	}
	fmt.Println(id)
	response.JsonFailMsg(c, "哈哈")
}
