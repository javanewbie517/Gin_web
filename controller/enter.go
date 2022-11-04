package controller

import (
	s "bubble/service"
)

type ApiGroup struct {
	UserApi
}

var (
	userService = s.ServiceGroupApp.UserService
)

var ApiGroupApp = new(ApiGroup)
