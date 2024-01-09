package sys_controller

import (
	"backend/internal/dto"
	"backend/pkg/response"
	"backend/pkg/utils"
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
	var res = dto.NavDTO{}
	// 组装一级菜单和二级菜单
	//id, err := u.menuService.ListMenuByUserId(c.GetInt64(constant.JWT_INFO_UID))
	sysMenus, err := u.menuService.ListMenuByUserId(1)
	if err != nil {
		return
	}
	root := make([]dto.SysMenuDTO, 0)
	authorities := make([]string, 0)
	for _, item := range sysMenus {
		if !utils.IsBlank(item.Perms) {
			authorities = append(authorities, item.Perms) // 切分成字符串
		}
		if item.ParentId == 0 {
			rootSon := make([]dto.SysMenuDTO, 0)
			for _, subItem := range sysMenus {
				if subItem.ParentId == item.ID {
					sub := dto.SysMenuDTO{
						ID:        subItem.ID,
						CreatedAt: subItem.CreatedAt,
						UpdatedAt: subItem.UpdatedAt,
						DeletedAt: subItem.DeletedAt,
						ParentId:  subItem.ParentId,
						Name:      subItem.Name,
						Url:       subItem.Url,
						Perms:     subItem.Perms,
						Type:      subItem.Type,
						Icon:      subItem.Icon,
						OrderNum:  subItem.OrderNum,
					}
					rootSon = append(rootSon, sub)
				}
			}
			menu := dto.SysMenuDTO{
				ID:        item.ID,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
				DeletedAt: item.DeletedAt,
				ParentId:  item.ParentId,
				Name:      item.Name,
				Url:       item.Url,
				Perms:     item.Perms,
				Type:      item.Type,
				Icon:      item.Icon,
				OrderNum:  item.OrderNum,
				List:      &rootSon,
			}
			root = append(root, menu)
			fmt.Println(root)
			fmt.Println(rootSon)
			fmt.Println("")
		}
	}
	res.MenuList = root
	res.Authorities = authorities
	response.JsonSuccessData(c, "哈哈", res)
}
