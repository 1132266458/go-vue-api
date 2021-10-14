package model

import (
	"go-vue-api/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);DEFAULT null " json:"username" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Phone    string `gorm:"type:varchar(11);not null" json:"phone" validate:"required,len=11" label:"手机号"`
}

//新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
