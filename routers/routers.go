package routers

import (
	"bubble/server/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	Router := gin.Default()
	//// 告诉gin框架模板文件引用的静态文件去哪里找
	//r.Static("/static", "static")
	//// 告诉gin框架去哪里找模板文件
	//r.LoadHTMLGlob("templates/*")

	Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		RouterGroupApp.BaseRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	PrivateGroup := Router.Group("")
	//PrivateGroup.Use(middleware.JWTAuth()) //使用中间件
	{
		RouterGroupApp.UserRouter.InitUserRouter(PrivateGroup)
	}

	return Router
}
