package system

import (
	"bubble/models"
	"bubble/server/global"
	"bubble/server/util"
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type UserService struct{}

func (userService *UserService) Register(u *models.User) (user models.User, err error) {
	if !errors.Is(global.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return user, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = util.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return *u, err
}

func (userService *UserService) Login(u *models.User) (user models.User, err error) {
	err = global.GVA_DB.Where("username = ?", u.Username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return user, errors.New("当前用户不存在")
	} else if err == nil {
		if ok := util.BcryptCheck(u.Password, user.Password); !ok {
			return models.User{}, errors.New("密码错误")
		}
	}
	return user, err
}
func (userService *UserService) GetUserById(userId string) (user models.User, err error) {
	err = global.GVA_DB.Where("uuid = ?", userId).First(&user).Error
	return user, err
}

func (userService *UserService) GetAllUser() (userList []models.User, err error) {
	err = global.GVA_DB.Find(&userList).Error
	return userList, err
}
