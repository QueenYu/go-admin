package router

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Dlimingliang/go-admin/api/v1"
)

func InitUserRouter(group *gin.RouterGroup) {
	userGroup := group.Group("user")
	{
		userGroup.POST("admin_register", v1.RegisterAdmin) //注册管理员
		userGroup.PUT("setUserInfo", v1.SetUserInfo)       //设置用户信息
		userGroup.PUT("resetPassword", v1.ResetPassword)   //重置密码为123456
		userGroup.DELETE("deleteUser", v1.DeleteUser)      //删除用户
	}

	userGroupWithoutRecord := group.Group("user")
	{
		userGroupWithoutRecord.POST("getUserList", v1.GetUserList)
	}
}
