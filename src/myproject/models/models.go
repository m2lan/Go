package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int64
	Username string
	Password string
}

type Topic struct {
	Id           int64
	Uid          int64
	Title        string
	Content      string `orm:"size(5000)"`
	Attachment   string
	Created      time.Time `orm:index`
	Updated      time.Time `orm:index`
	Views        int64     `orm:index`
	Athor        string
	ReplyTime    time.Time `orm:index`
	ReplyCount   int64
	ReplyLastUid int64
}

const (
	_MYSQL_DRIVER = "mysql"
	_DATASOURCE   = "root:feng@/blog?charset=utf8"
)

func RegisterDB() {
	orm.RegisterModel(new(User), new(Topic))
	orm.RegisterDriver(_MYSQL_DRIVER, orm.DR_MySQL)
	orm.RegisterDataBase("default", _MYSQL_DRIVER, _DATASOURCE, 30)
}
