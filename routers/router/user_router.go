package router

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	userApi := controller.ApiGroupApp.UserApi
	{
		userRouter.GET("getUserById", userApi.GetUserById) // 获取指定用户
		userRouter.GET("getAllUser", userApi.GetAllUser)   // 获取全部用户列表
	}
}
