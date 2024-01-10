package sys_controller

import (
	"backend/internal/forms/cms_sys_form"
	"backend/pkg/constant"
	"backend/pkg/response"
	"github.com/gin-gonic/gin"
)

/* 创建后台用户 */
func (u *SysController) CreateUserHandler(c *gin.Context) {
	resp := cms_sys_form.CreateAdminUserForm{}
	err := c.ShouldBindJSON(&resp)
	if err != nil {
		return
	}
	user, err := u.userService.CreateUser(&resp)
	if err != nil {
		response.JsonFailMsg(c, err.Error())
		return
	}
	response.JsonSuccessData(c, "成功！", user)
	return
}

// 获取用户信息
/*
{
    "code": "00000",
    "msg": null,
    "data": {
        "userId": 1,
        "username": "admin",
        "email": "root@123.com",
        "mobile": "11111111111",
        "status": 1,
        "shopId": 1,
        "roleIdList": null,
        "createTime": "2016-11-11 11:11:11"
    },
    "version": "mall4j.v230424",
    "timestamp": null,
    "sign": null,
    "success": true,
    "fail": false
}
*/
func (u *SysController) AdminUserInfo(c *gin.Context) {
	//var res dto.UserInfoDTO
	user, err := u.userService.Get(c.GetInt64(constant.JWT_INFO_UID))
	if err != nil {
		response.JsonFailMsg(c, err.Error())
		return
	}
	response.JsonSuccessData(c, "成功！", user)
}
