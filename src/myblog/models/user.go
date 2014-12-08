package models

import (
	"github.com/astaxie/beego/orm"
)

// 用户
type User struct {
	Id       int
	Username string
	Password string
	Email    string
}

func (this *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

// 查找用户，username
func GetUserByName(username string) (User, error) {
	o := orm.NewOrm()
	user := User{Username: username}
	err := o.Read(&user, "username")

	return user, err
}
