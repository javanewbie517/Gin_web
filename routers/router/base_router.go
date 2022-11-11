package router

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

// InitBaseRouter 基础路由不做鉴权
func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("base")
	userApi := controller.ApiGroupApp.UserApi
	{
		baseRouter.POST("userRegister", userApi.Register)           // 用户注册账号
		baseRouter.POST("userLogin", userApi.Login)                 // 用户登录账号
		baseRouter.POST("testTransaction", userApi.TestTransaction) // 测试事务回滚
	}
}
