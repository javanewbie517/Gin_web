package routers

import "bubble/routers/router"

type RouterGroup struct {
	router.UserRouter
	router.BaseRouter
}

var RouterGroupApp = new(RouterGroup)
