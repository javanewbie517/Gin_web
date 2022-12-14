package controller

import (
	"bubble/models"
	"bubble/server/response"
	"bubble/server/util/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserApi struct{}

func (u *UserApi) Register(c *gin.Context) {
	var user models.User
	err := c.ShouldBindQuery(&user)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	userMage, err := userService.Register(&user)
	if err != nil {
		response.FailWithDetailed(c, userMage, "注册失败")
		return
	}
	response.OkWithDetailed(c, userMage, "注册成功")
}

func (u *UserApi) Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindQuery(&user)
	if err != nil {
		response.FailWithMessage(c, "参数解析错误")
		return
	}
	userMsg, err := userService.Login(&user)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	token, err := jwt.GenerateToken(userMsg.Username, userMsg.Password)
	response.OkWithDetailed(c,
		map[string]interface{}{
			"userMsg": userMsg,
			"token":   token,
		}, "注册成功")
}

func (u *UserApi) GetUserById(c *gin.Context) {

	userId, _ := c.GetQuery("userId")
	claims, _ := c.Get("claims")
	fmt.Println(claims)
	user, err := userService.GetUserById(userId)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithDetailed(c, user, "获取成功")
}

func (u *UserApi) GetAllUser(c *gin.Context) {

	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	userList, total, err := userService.GetAllUser(pageNum, pageSize)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithDetailed(c,
		map[string]interface{}{
			"userList": userList,
			"total":    total,
		}, "获取成功")
}
func (u *UserApi) TestTransaction(c *gin.Context) {
	err := userService.TestTransaction()
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	response.OkWithMessage(c, "操作成功")
}
