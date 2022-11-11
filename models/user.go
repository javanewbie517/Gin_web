package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	UUID      uuid.UUID `json:"id" gorm:"primary_key"`
	Username  string    `json:"userName" gorm:"size:255" `  // 用户登录名
	Password  string    `json:"passWord" gorm:"size:255"`   // 用户登录密码
	Phone     string    `json:"phone" gorm:"size:11"`       // 用户手机号
	Email     string    `json:"email" gorm:"size:30"`       // 用户邮箱
	CreatedAt time.Time `json:"createdAt"`                  // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`                  // 修改时间
	Note      string    `json:"note" gorm:"default:'默认备注'"` // 默认值
	Enable    int       `json:"enable"`                     //用户是否被冻结 0异常 1正常 2冻结
}

func (User) TableName() string {
	return "user"
}
