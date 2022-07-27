package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(100);" json:"name"`        // 用户名
	Password string `gorm:"column:password;type:varchar(32);" json:"password"` //密码
	Phone    string `gorm:"column:phone;type:varchar(20);" json:"phone"`       //手机号
	Mail     string `gorm:"column:mail;type:varchar(100);" json:"mail"`        //邮件

}

func (table *User) TableName() string {
	return "user"
}
